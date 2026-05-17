/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net"
	"strings"
	"testing"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/sec"
)

func decodePEMCert(t *testing.T, b64 string) *x509.Certificate {
	t.Helper()
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		t.Fatalf("base64 decode failed: %v", err)
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		t.Fatalf("pem decode returned nil block")
	}
	if block.Type != "CERTIFICATE" {
		t.Fatalf("expected CERTIFICATE block, got %q", block.Type)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("x509 parse failed: %v", err)
	}
	return cert
}

func TestCreateCertBundle(t *testing.T) {
	caB64, keyB64, crtB64 := sec.CreateCertBundle()

	if caB64 == "" || keyB64 == "" || crtB64 == "" {
		t.Fatalf("CreateCertBundle returned empty string(s): ca=%d key=%d crt=%d", len(caB64), len(keyB64), len(crtB64))
	}

	caCert := decodePEMCert(t, caB64)
	if !caCert.IsCA {
		t.Error("expected CA cert to have IsCA=true")
	}
	if caCert.Subject.CommonName != "saichler" {
		t.Errorf("expected CA CommonName 'saichler', got %q", caCert.Subject.CommonName)
	}

	leafCert := decodePEMCert(t, crtB64)
	if leafCert.IsCA {
		t.Error("expected leaf cert to have IsCA=false")
	}
	if leafCert.Subject.CommonName != "www.layer8vibe.dev" {
		t.Errorf("expected leaf CommonName 'www.layer8vibe.dev', got %q", leafCert.Subject.CommonName)
	}

	// Verify the leaf cert was signed by the CA
	roots := x509.NewCertPool()
	roots.AddCert(caCert)
	if _, err := leafCert.Verify(x509.VerifyOptions{Roots: roots}); err != nil {
		t.Errorf("leaf cert failed to verify against CA: %v", err)
	}

	// Verify the private key is a valid RSA key
	keyPEM, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		t.Fatalf("private key base64 decode failed: %v", err)
	}
	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		t.Fatal("private key pem decode returned nil block")
	}
	if keyBlock.Type != "RSA PRIVATE KEY" {
		t.Errorf("expected 'RSA PRIVATE KEY' block, got %q", keyBlock.Type)
	}
	if _, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes); err != nil {
		t.Errorf("failed to parse private key: %v", err)
	}
}

func TestCreateCertBundle_UniqueAcrossCalls(t *testing.T) {
	ca1, key1, crt1 := sec.CreateCertBundle()
	ca2, key2, crt2 := sec.CreateCertBundle()

	if ca1 == ca2 {
		t.Error("expected distinct CA certs across calls")
	}
	if key1 == key2 {
		t.Error("expected distinct private keys across calls")
	}
	if crt1 == crt2 {
		t.Error("expected distinct leaf certs across calls")
	}
}

func TestNewShallowSecurityProvider(t *testing.T) {
	p := sec.NewShallowSecurityProvider()
	if p == nil {
		t.Fatal("expected non-nil provider")
	}
}

func TestShallowSecurityProvider_EncryptDecryptRoundTrip(t *testing.T) {
	p := sec.NewShallowSecurityProvider()
	cases := [][]byte{
		[]byte("hello world"),
		[]byte(""),
		[]byte("\x00\x01\x02\xff"),
		[]byte(strings.Repeat("a", 1024)),
	}
	for _, plain := range cases {
		ciphertext, err := p.Encrypt(plain)
		if err != nil {
			t.Fatalf("Encrypt(%q) failed: %v", plain, err)
		}
		got, err := p.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Decrypt failed: %v", err)
		}
		if string(got) != string(plain) {
			t.Errorf("round trip mismatch: want %q, got %q", plain, got)
		}
	}
}

func TestShallowSecurityProvider_EncryptProducesUniqueOutput(t *testing.T) {
	p := sec.NewShallowSecurityProvider()
	a, err := p.Encrypt([]byte("same plaintext"))
	if err != nil {
		t.Fatalf("Encrypt #1 failed: %v", err)
	}
	b, err := p.Encrypt([]byte("same plaintext"))
	if err != nil {
		t.Fatalf("Encrypt #2 failed: %v", err)
	}
	// AES-GCM (used by aes.Encrypt) generates a fresh nonce each call,
	// so identical plaintexts must yield distinct ciphertexts.
	if a == b {
		t.Error("expected distinct ciphertexts for identical plaintexts (nonce reuse)")
	}
}

func TestShallowSecurityProvider_PermissiveAuthorization(t *testing.T) {
	p := sec.NewShallowSecurityProvider()

	if err := p.CanAccept(nil); err != nil {
		t.Errorf("CanAccept: expected nil, got %v", err)
	}
	if err := p.CanDoAction(nil, ifs.POST, nil, "", ""); err != nil {
		t.Errorf("CanDoAction: expected nil, got %v", err)
	}
	if got := p.ScopeView(nil, nil, "", ""); got != nil {
		t.Errorf("ScopeView: expected pass-through nil, got %v", got)
	}
	if got := p.AllowedTypes(nil, ""); got != nil {
		t.Errorf("AllowedTypes: expected nil, got %v", got)
	}
	if got := p.AllowedActions(nil, ""); got != nil {
		t.Errorf("AllowedActions: expected nil, got %v", got)
	}
}

func TestShallowSecurityProvider_AuthenticateAndValidateToken(t *testing.T) {
	p := sec.NewShallowSecurityProvider()

	token, _, _, _, _, err := p.Authenticate("user", "pass", nil)
	if err != nil {
		t.Fatalf("Authenticate returned error: %v", err)
	}
	if token != "bearer token" {
		t.Errorf("Authenticate: expected token 'bearer token', got %q", token)
	}

	uuid, ok := p.ValidateToken("any-token", nil)
	if !ok {
		t.Error("ValidateToken: expected ok=true")
	}
	if len(uuid) != 36 {
		t.Errorf("ValidateToken: expected 36-char uuid, got %q (len=%d)", uuid, len(uuid))
	}
}

func TestShallowSecurityProvider_StubMethodsDoNotPanic(t *testing.T) {
	p := sec.NewShallowSecurityProvider()

	msg, err := p.Message("aaaid", nil)
	if err != nil {
		t.Errorf("Message: unexpected error %v", err)
	}
	if msg == nil {
		t.Error("Message: expected non-nil message")
	}

	if _, _, err := p.TFASetup("user", nil); err != nil {
		t.Errorf("TFASetup: unexpected error %v", err)
	}
	if err := p.TFAVerify("user", "code", "bearer", nil); err != nil {
		t.Errorf("TFAVerify: unexpected error %v", err)
	}
	if got := p.Captcha(); got != nil {
		t.Errorf("Captcha: expected nil, got %v", got)
	}
	if err := p.Register("user", "pass", "captcha", nil); err != nil {
		t.Errorf("Register: unexpected error %v", err)
	}

	// Should not panic
	p.AddAdjacent(nil)
}

func TestShallowSecurityProvider_Credential(t *testing.T) {
	p := sec.NewShallowSecurityProvider()

	cases := []struct {
		name        string
		crId, cId   string
		wantA, wantB string
	}{
		{"sim-ssh", "sim", "ssh", "simadmin", "simadmin"},
		{"sim-snmp", "sim", "snmp", "public", "private"},
		{"unknown defaults to admin/admin", "other", "other", "admin", "admin"},
		{"empty defaults to admin/admin", "", "", "admin", "admin"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, b, _, _, err := p.Credential(c.crId, c.cId, nil)
			if err != nil {
				t.Fatalf("Credential returned error: %v", err)
			}
			if a != c.wantA || b != c.wantB {
				t.Errorf("Credential(%q,%q) = (%q,%q), want (%q,%q)", c.crId, c.cId, a, b, c.wantA, c.wantB)
			}
		})
	}
}

func TestShallowSecurityProvider_NewSystemConfig(t *testing.T) {
	p := sec.NewShallowSecurityProvider()
	cfg := p.NewSystemConfig()

	if cfg == nil {
		t.Fatal("expected non-nil sys config")
	}
	if cfg.MaxDataSize != 1024*1024*50 {
		t.Errorf("MaxDataSize: got %d, want %d", cfg.MaxDataSize, 1024*1024*50)
	}
	if cfg.VnetPort != 10005 {
		t.Errorf("VnetPort: got %d, want 10005", cfg.VnetPort)
	}
	if cfg.WebConfig == nil {
		t.Fatal("expected non-nil WebConfig")
	}
	if cfg.WebConfig.WebPort != 4443 {
		t.Errorf("WebPort: got %d, want 4443", cfg.WebConfig.WebPort)
	}
	if cfg.WebConfig.DomainCertPem == "" || cfg.WebConfig.PrivateKeyPem == "" || cfg.WebConfig.PublicKeyPem == "" {
		t.Error("expected WebConfig cert fields to be populated by CreateCertBundle")
	}
	// Validate that the populated cert fields are real, parseable cert material.
	_ = decodePEMCert(t, cfg.WebConfig.DomainCertPem)
	_ = decodePEMCert(t, cfg.WebConfig.PublicKeyPem)
	if cfg.DataDirectory == "" {
		t.Error("expected DataDirectory to be set")
	}
}

func TestShallowSecurityProvider_CanDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start listener: %v", err)
	}
	defer listener.Close()

	port := uint32(listener.Addr().(*net.TCPAddr).Port)

	p := sec.NewShallowSecurityProvider()
	conn, err := p.CanDial("127.0.0.1", port)
	if err != nil {
		t.Fatalf("CanDial failed: %v", err)
	}
	if conn == nil {
		t.Fatal("expected non-nil connection")
	}
	conn.Close()
}

func TestShallowSecurityProvider_CanDialIPv6Bracketing(t *testing.T) {
	// Listen on the IPv6 loopback. Skip if the platform doesn't support IPv6.
	listener, err := net.Listen("tcp", "[::1]:0")
	if err != nil {
		t.Skipf("IPv6 loopback not available: %v", err)
	}
	defer listener.Close()

	port := uint32(listener.Addr().(*net.TCPAddr).Port)

	p := sec.NewShallowSecurityProvider()
	// Pass the raw IPv6 address (no brackets); CanDial must add brackets.
	conn, err := p.CanDial("::1", port)
	if err != nil {
		t.Fatalf("CanDial(::1) failed: %v", err)
	}
	conn.Close()
}

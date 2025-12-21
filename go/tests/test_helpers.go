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
	"errors"
	"net"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types/l8api"
	"github.com/saichler/l8types/go/types/l8sysconfig"
)

// MockSecurityProvider implements ISecurityProvider for testing
type MockSecurityProvider struct {
	encryptError bool
	decryptError bool
}

func (m *MockSecurityProvider) Authenticate(string, string) (string, bool, bool, error) {
	return "", false, false, nil
}
func (m *MockSecurityProvider) ValidateToken(string) (string, bool)      { return "", true }
func (m *MockSecurityProvider) Message(string) (*ifs.Message, error)     { return nil, nil }
func (m *MockSecurityProvider) CanDial(string, uint32) (net.Conn, error) { return nil, nil }
func (m *MockSecurityProvider) CanAccept(net.Conn) error                 { return nil }
func (m *MockSecurityProvider) ValidateConnection(net.Conn, *l8sysconfig.L8SysConfig) error {
	return nil
}
func (m *MockSecurityProvider) CanDoAction(ifs.Action, ifs.IElements, string, string, ...string) error {
	return nil
}
func (m *MockSecurityProvider) ScopeView(ifs.IElements, string, string, ...string) ifs.IElements {
	return nil
}
func (m *MockSecurityProvider) TFASetup(string, ifs.IVNic) (string, []byte, error) {
	return "", nil, nil
}
func (m *MockSecurityProvider) TFAVerify(string, string, string, ifs.IVNic) error { return nil }
func (m *MockSecurityProvider) Captcha() []byte                                   { return nil }
func (m *MockSecurityProvider) Register(string, string, string, ifs.IVNic) error  { return nil }
func (m *MockSecurityProvider) Credential(string, string, ifs.IResources) (string, string, string, string, error) {
	return "", "", "", "", nil
}

func (m *MockSecurityProvider) Encrypt(data []byte) (string, error) {
	if m.encryptError {
		return "", errors.New("encryption failed")
	}
	return string(data), nil
}

func (m *MockSecurityProvider) Decrypt(data string) ([]byte, error) {
	if m.decryptError {
		return nil, errors.New("decryption failed")
	}
	return []byte(data), nil
}

// MockResources implements IResources for testing
type MockResources struct {
	security *MockSecurityProvider
}

func (m *MockResources) Registry() ifs.IRegistry                       { return nil }
func (m *MockResources) Services() ifs.IServices                       { return nil }
func (m *MockResources) Security() ifs.ISecurityProvider               { return m.security }
func (m *MockResources) DataListener() ifs.IDatatListener              { return nil }
func (m *MockResources) Serializer(ifs.SerializerMode) ifs.ISerializer { return nil }
func (m *MockResources) Logger() ifs.ILogger                           { return nil }
func (m *MockResources) SysConfig() *l8sysconfig.L8SysConfig           { return nil }
func (m *MockResources) Introspector() ifs.IIntrospector               { return nil }
func (m *MockResources) AddService(string, int32)                      {}
func (m *MockResources) Set(interface{})                               {}
func (m *MockResources) Copy(ifs.IResources)                           {}
func (m *MockResources) DefaultUser() *l8api.AuthUser                  { return nil }

func newMockResources() *MockResources {
	return &MockResources{
		security: &MockSecurityProvider{},
	}
}

func newMockResourcesWithError(encryptError, decryptError bool) *MockResources {
	return &MockResources{
		security: &MockSecurityProvider{
			encryptError: encryptError,
			decryptError: decryptError,
		},
	}
}

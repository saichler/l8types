package tests

import (
	"strings"
	"testing"

	aeslib "github.com/saichler/l8types/go/aes"
)

func TestGenerateAES256Key(t *testing.T) {
	t.Run("KeyLength", func(t *testing.T) {
		key := aeslib.GenerateAES256Key()
		if len(key) != 32 {
			t.Errorf("Expected key length 32, got %d", len(key))
		}
	})

	t.Run("KeyRandomness", func(t *testing.T) {
		key1 := aeslib.GenerateAES256Key()
		key2 := aeslib.GenerateAES256Key()
		if key1 == key2 {
			t.Error("Generated keys should be different")
		}
	})

	t.Run("KeyCharacterSet", func(t *testing.T) {
		key := aeslib.GenerateAES256Key()
		validChars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for _, char := range key {
			if !strings.ContainsRune(validChars, char) {
				t.Errorf("Key contains invalid character: %c", char)
			}
		}
	})

	t.Run("MultipleKeys", func(t *testing.T) {
		keys := make(map[string]bool)
		for i := 0; i < 100; i++ {
			key := aeslib.GenerateAES256Key()
			if keys[key] {
				t.Error("Duplicate key generated")
			}
			keys[key] = true
		}
	})
}

func TestEncryptDecrypt(t *testing.T) {
	key := aeslib.GenerateAES256Key()
	testData := []byte("Hello, World! This is a test message.")

	t.Run("BasicEncryptDecrypt", func(t *testing.T) {
		encrypted, err := aeslib.Encrypt(testData, key)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		if encrypted == "" {
			t.Error("Encrypted string should not be empty")
		}

		decrypted, err := aeslib.Decrypt(encrypted, key)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != string(testData) {
			t.Errorf("Decrypted data doesn't match original. Expected: %s, Got: %s", testData, decrypted)
		}
	})

	t.Run("EmptyData", func(t *testing.T) {
		emptyData := []byte("")
		encrypted, err := aeslib.Encrypt(emptyData, key)
		if err != nil {
			t.Fatalf("Encryption of empty data failed: %v", err)
		}

		decrypted, err := aeslib.Decrypt(encrypted, key)
		if err != nil {
			t.Fatalf("Decryption of empty data failed: %v", err)
		}

		if len(decrypted) != 0 {
			t.Error("Decrypted empty data should be empty")
		}
	})

	t.Run("LargeData", func(t *testing.T) {
		largeData := make([]byte, 1024*1024) // 1MB
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		encrypted, err := aeslib.Encrypt(largeData, key)
		if err != nil {
			t.Fatalf("Encryption of large data failed: %v", err)
		}

		decrypted, err := aeslib.Decrypt(encrypted, key)
		if err != nil {
			t.Fatalf("Decryption of large data failed: %v", err)
		}

		if len(decrypted) != len(largeData) {
			t.Errorf("Decrypted data length mismatch. Expected: %d, Got: %d", len(largeData), len(decrypted))
		}

		for i, b := range decrypted {
			if b != largeData[i] {
				t.Errorf("Data mismatch at index %d. Expected: %d, Got: %d", i, largeData[i], b)
				break
			}
		}
	})

	t.Run("DifferentKeys", func(t *testing.T) {
		key1 := aeslib.GenerateAES256Key()
		key2 := aeslib.GenerateAES256Key()

		encrypted, err := aeslib.Encrypt(testData, key1)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		decrypted, err := aeslib.Decrypt(encrypted, key2)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) == string(testData) {
			t.Error("Decrypted data should not match original when using wrong key")
		}
	})

	t.Run("UniqueEncryption", func(t *testing.T) {
		encrypted1, err := aeslib.Encrypt(testData, key)
		if err != nil {
			t.Fatalf("First encryption failed: %v", err)
		}

		encrypted2, err := aeslib.Encrypt(testData, key)
		if err != nil {
			t.Fatalf("Second encryption failed: %v", err)
		}

		if encrypted1 == encrypted2 {
			t.Error("Same plaintext should produce different ciphertext due to random IV")
		}
	})
}

func TestDecryptErrors(t *testing.T) {
	key := aeslib.GenerateAES256Key()

	t.Run("InvalidBase64", func(t *testing.T) {
		_, err := aeslib.Decrypt("invalid-base64!", key)
		if err == nil {
			t.Error("Should fail with invalid base64")
		}
	})

	t.Run("ShortData", func(t *testing.T) {
		shortData := "dGVzdA==" // "test" in base64, shorter than AES block size
		_, err := aeslib.Decrypt(shortData, key)
		if err == nil {
			t.Error("Should fail with data shorter than AES block size")
		}
		if !strings.Contains(err.Error(), "iv spec") {
			t.Errorf("Expected error about IV spec, got: %v", err)
		}
	})

	t.Run("InvalidKey", func(t *testing.T) {
		testData := []byte("test")
		encrypted, _ := aeslib.Encrypt(testData, key)
		
		shortKey := "short"
		_, err := aeslib.Decrypt(encrypted, shortKey)
		if err == nil {
			t.Error("Should fail with invalid key length")
		}
	})
}

func TestEncryptErrors(t *testing.T) {
	t.Run("InvalidKeyLength", func(t *testing.T) {
		testData := []byte("test")
		shortKey := "short"
		
		_, err := aeslib.Encrypt(testData, shortKey)
		if err == nil {
			t.Error("Should fail with invalid key length")
		}
	})

	t.Run("ValidKeyLengths", func(t *testing.T) {
		testData := []byte("test")
		
		// Test 16-byte key (AES-128)
		key16 := strings.Repeat("a", 16)
		_, err := aeslib.Encrypt(testData, key16)
		if err != nil {
			t.Errorf("16-byte key should work: %v", err)
		}

		// Test 24-byte key (AES-192)
		key24 := strings.Repeat("a", 24)
		_, err = aeslib.Encrypt(testData, key24)
		if err != nil {
			t.Errorf("24-byte key should work: %v", err)
		}

		// Test 32-byte key (AES-256)
		key32 := strings.Repeat("a", 32)
		_, err = aeslib.Encrypt(testData, key32)
		if err != nil {
			t.Errorf("32-byte key should work: %v", err)
		}
	})
}

func BenchmarkGenerateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aeslib.GenerateAES256Key()
	}
}

func BenchmarkEncrypt(b *testing.B) {
	key := aeslib.GenerateAES256Key()
	data := []byte("This is a test message for benchmarking encryption performance.")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aeslib.Encrypt(data, key)
		if err != nil {
			b.Fatalf("Encryption failed: %v", err)
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	key := aeslib.GenerateAES256Key()
	data := []byte("This is a test message for benchmarking decryption performance.")
	encrypted, err := aeslib.Encrypt(data, key)
	if err != nil {
		b.Fatalf("Setup encryption failed: %v", err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aeslib.Decrypt(encrypted, key)
		if err != nil {
			b.Fatalf("Decryption failed: %v", err)
		}
	}
}
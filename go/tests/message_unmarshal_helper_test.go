package tests

import (
	"strings"
	"testing"
	"github.com/saichler/l8types/go/ifs"
)

func TestHeaderOf(t *testing.T) {
	// Create test data that matches the expected header format
	data := make([]byte, 120) // Total header size (119 + 1 for priority)
	
	// Set source (exactly 36 bytes)
	source36 := "source-uuid-123456789012345678901234"
	copy(data[0:36], source36)
	
	// Set vnet (exactly 36 bytes with padding)
	vnet36 := "vnet-uuid-123456789012345678901234"
	copy(data[36:72], vnet36)
	
	// Set destination (exactly 36 bytes with padding)
	dest36 := "destination-uuid-12345678901234567"
	copy(data[72:108], dest36)
	
	// Set service name (10 bytes) - null terminated
	copy(data[108:118], "test-serv\x00")
	
	// Set service area (1 byte)
	data[118] = 5
	
	// Set priority (1 byte)
	data[119] = byte(ifs.P3)
	
	// Call HeaderOf
	source, vnet, destination, serviceName, serviceArea, priority := ifs.HeaderOf(data)
	
	// Verify results - account for null padding that becomes spaces
	if !strings.HasPrefix(source, source36) || len(source) != 36 {
		t.Errorf("Source mismatch: expected to start with '%s' and be 36 chars, got '%s' (len=%d)", source36, source, len(source))
	}
	
	if !strings.HasPrefix(vnet, vnet36) || len(vnet) != 36 {
		t.Errorf("Vnet mismatch: expected to start with '%s' and be 36 chars, got '%s' (len=%d)", vnet36, vnet, len(vnet))
	}
	
	if !strings.HasPrefix(destination, dest36) || len(destination) != 36 {
		t.Errorf("Destination mismatch: expected to start with '%s' and be 36 chars, got '%s' (len=%d)", dest36, destination, len(destination))
	}
	
	expectedServiceName := "test-serv"
	if serviceName != expectedServiceName {
		t.Errorf("ServiceName mismatch: expected '%s', got '%s'", expectedServiceName, serviceName)
	}
	
	if serviceArea != 5 {
		t.Errorf("ServiceArea mismatch: expected 5, got %d", serviceArea)
	}
	
	if priority != ifs.P3 {
		t.Errorf("Priority mismatch: expected P3, got %v", priority)
	}
}

func TestToDestinationWithValidData(t *testing.T) {
	data := make([]byte, 119)
	
	// Set destination with non-zero bytes
	copy(data[72:108], "valid-destination-123456789012345678")
	
	result := ifs.ToDestination(data)
	expected := "valid-destination-123456789012345678"
	
	if result != expected {
		t.Errorf("ToDestination mismatch: expected '%s', got '%s'", expected, result)
	}
}

func TestToDestinationWithZeroBytes(t *testing.T) {
	data := make([]byte, 119)
	
	// Set first byte of destination to zero
	data[72] = 0
	data[73] = 0
	
	result := ifs.ToDestination(data)
	
	if result != "" {
		t.Errorf("ToDestination should return empty string for zero bytes, got '%s'", result)
	}
}

func TestToDestinationWithFirstByteZero(t *testing.T) {
	data := make([]byte, 119)
	
	// Set first byte to zero, second byte non-zero
	data[72] = 0
	data[73] = 65 // 'A'
	
	result := ifs.ToDestination(data)
	
	if result != "" {
		t.Errorf("ToDestination should return empty string when first byte is zero, got '%s'", result)
	}
}

func TestToDestinationWithSecondByteZero(t *testing.T) {
	data := make([]byte, 119)
	
	// Set first byte non-zero, second byte zero
	data[72] = 65 // 'A'
	data[73] = 0
	
	result := ifs.ToDestination(data)
	
	if result != "" {
		t.Errorf("ToDestination should return empty string when second byte is zero, got '%s'", result)
	}
}

func TestToServiceNameWithValidData(t *testing.T) {
	data := make([]byte, 119)
	
	// Set service name with null terminator
	copy(data[108:118], "service\x00\x00\x00")
	
	result := ifs.ToServiceName(data)
	expected := "service"
	
	if result != expected {
		t.Errorf("ToServiceName mismatch: expected '%s', got '%s'", expected, result)
	}
}

func TestToServiceNameWithoutNullTerminator(t *testing.T) {
	data := make([]byte, 119)
	
	// Set service name without null terminator (fill all 10 bytes)
	copy(data[108:118], "1234567890")
	
	result := ifs.ToServiceName(data)
	expected := "1234567890"
	
	if result != expected {
		t.Errorf("ToServiceName mismatch: expected '%s', got '%s'", expected, result)
	}
}

func TestToServiceNameWithEarlyNullTerminator(t *testing.T) {
	data := make([]byte, 119)
	
	// Set service name with early null terminator
	copy(data[108:118], "ab\x00cdefghi")
	
	result := ifs.ToServiceName(data)
	expected := "ab"
	
	if result != expected {
		t.Errorf("ToServiceName mismatch: expected '%s', got '%s'", expected, result)
	}
}

func TestToServiceNameEmpty(t *testing.T) {
	data := make([]byte, 119)
	
	// Set service name area to all zeros
	for i := 108; i < 118; i++ {
		data[i] = 0
	}
	
	result := ifs.ToServiceName(data)
	expected := ""
	
	if result != expected {
		t.Errorf("ToServiceName mismatch: expected empty string, got '%s'", result)
	}
}

func TestToServiceNameWithSpecialCharacters(t *testing.T) {
	data := make([]byte, 119)
	
	// Set service name with special characters
	copy(data[108:118], "test-svc_1\x00")
	
	result := ifs.ToServiceName(data)
	expected := "test-svc_1"
	
	if result != expected {
		t.Errorf("ToServiceName mismatch: expected '%s', got '%s'", expected, result)
	}
}

// Test with buffer boundary conditions
func TestToServiceNameBoundaryConditions(t *testing.T) {
	// Test with exact minimum header size
	data := make([]byte, 120) // Full header size
	
	// Fill service name area with zeros
	for i := 108; i < 118; i++ {
		data[i] = 0
	}
	
	result := ifs.ToServiceName(data)
	expected := ""
	
	if result != expected {
		t.Errorf("ToServiceName with minimal data should return empty string, got '%s'", result)
	}
	
	// Test with exact service name size
	data = make([]byte, 120) // Full header size
	copy(data[108:118], "exactsize\x00")
	
	result = ifs.ToServiceName(data)
	expected = "exactsize"
	
	if result != expected {
		t.Errorf("ToServiceName with exact size mismatch: expected '%s', got '%s'", expected, result)
	}
}
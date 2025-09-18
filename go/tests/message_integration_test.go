package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/saichler/l8types/go/ifs"
)

func TestCompleteMessageWorkflow(t *testing.T) {
	resources := newMockResources()

	// Create a realistic message scenario
	originalMsg := &ifs.Message{}
	originalMsg.Init(
		"microservice-auth-uuid-123456789012",
		"auth-svc", // Keep service name <= 10 chars
		5,
		ifs.P2,
		ifs.M_All,
		ifs.POST,
		"client-uuid-987654321098765432109876",
		"vnet-production-cluster-east-12345",
		[]byte(`{"username":"testuser","password":"hashedpwd","timestamp":"2024-01-01T00:00:00Z"}`),
		true,
		false,
		98765432,
		ifs.Start,
		"transaction-auth-session-uuid-567890",
		"",
		time.Now().Unix(), 30,
	)
	originalMsg.SetAAAId("aaa-session-uuid-456789012345678901")
	originalMsg.SetTimeout(5000)

	// Marshal
	marshaledData, err := originalMsg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Verify marshal produces data
	if len(marshaledData) == 0 {
		t.Fatal("Marshal produced empty data")
	}

	// Unmarshal
	reconstructedMsg := &ifs.Message{}
	_, err = reconstructedMsg.Unmarshal(marshaledData, resources)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Comprehensive verification
	verifyMessageEquality(t, originalMsg, reconstructedMsg)

	// Test header extraction without full unmarshal
	source, vnet, destination, serviceName, serviceArea, priority, multicastMode := ifs.HeaderOf(marshaledData)

	if !strings.HasPrefix(source, originalMsg.Source()) {
		t.Errorf("HeaderOf source mismatch: expected to start with '%s', got '%s'", originalMsg.Source(), source)
	}
	if !strings.HasPrefix(vnet, originalMsg.Vnet()) {
		t.Errorf("HeaderOf vnet mismatch: expected to start with '%s', got '%s'", originalMsg.Vnet(), vnet)
	}
	if !strings.HasPrefix(destination, originalMsg.Destination()) {
		t.Errorf("HeaderOf destination mismatch: expected to start with '%s', got '%s'", originalMsg.Destination(), destination)
	}
	if serviceName != originalMsg.ServiceName() {
		t.Errorf("HeaderOf serviceName mismatch: expected '%s', got '%s'", originalMsg.ServiceName(), serviceName)
	}
	if serviceArea != originalMsg.ServiceArea() {
		t.Errorf("HeaderOf serviceArea mismatch: expected %d, got %d", originalMsg.ServiceArea(), serviceArea)
	}
	if priority != originalMsg.Priority() {
		t.Errorf("HeaderOf priority mismatch: expected %v, got %v", originalMsg.Priority(), priority)
	}
	if multicastMode != originalMsg.MulticastMode() {
		t.Errorf("HeaderOf multicastMode mismatch: expected %v, got %v", originalMsg.MulticastMode(), multicastMode)
	}
}

func verifyMessageEquality(t *testing.T, original, reconstructed *ifs.Message) {
	// For fixed-size fields, check that reconstructed values start with original values
	if !strings.HasPrefix(reconstructed.Source(), original.Source()) {
		t.Errorf("Source mismatch: expected to start with '%s', got '%s'", original.Source(), reconstructed.Source())
	}
	if !strings.HasPrefix(reconstructed.Vnet(), original.Vnet()) {
		t.Errorf("Vnet mismatch: expected to start with '%s', got '%s'", original.Vnet(), reconstructed.Vnet())
	}
	if !strings.HasPrefix(reconstructed.Destination(), original.Destination()) {
		t.Errorf("Destination mismatch: expected to start with '%s', got '%s'", original.Destination(), reconstructed.Destination())
	}
	if reconstructed.ServiceName() != original.ServiceName() {
		t.Errorf("ServiceName mismatch: expected '%s', got '%s'", original.ServiceName(), reconstructed.ServiceName())
	}
	if reconstructed.ServiceArea() != original.ServiceArea() {
		t.Errorf("ServiceArea mismatch: expected %d, got %d", original.ServiceArea(), reconstructed.ServiceArea())
	}
	if reconstructed.Priority() != original.Priority() {
		t.Errorf("Priority mismatch: expected %v, got %v", original.Priority(), reconstructed.Priority())
	}
	if reconstructed.Action() != original.Action() {
		t.Errorf("Action mismatch: expected %v, got %v", original.Action(), reconstructed.Action())
	}
	if !strings.HasPrefix(reconstructed.AAAId(), original.AAAId()) {
		t.Errorf("AAAId mismatch: expected to start with '%s', got '%s'", original.AAAId(), reconstructed.AAAId())
	}
	if reconstructed.Sequence() != original.Sequence() {
		t.Errorf("Sequence mismatch: expected %d, got %d", original.Sequence(), reconstructed.Sequence())
	}
	if reconstructed.Timeout() != original.Timeout() {
		t.Errorf("Timeout mismatch: expected %d, got %d", original.Timeout(), reconstructed.Timeout())
	}
	if reconstructed.Request() != original.Request() {
		t.Errorf("Request mismatch: expected %v, got %v", original.Request(), reconstructed.Request())
	}
	if reconstructed.Reply() != original.Reply() {
		t.Errorf("Reply mismatch: expected %v, got %v", original.Reply(), reconstructed.Reply())
	}
	if reconstructed.FailMessage() != original.FailMessage() {
		t.Errorf("FailMessage mismatch: expected '%s', got '%s'", original.FailMessage(), reconstructed.FailMessage())
	}
	if string(reconstructed.Data()) != string(original.Data()) {
		t.Errorf("Data mismatch: expected '%s', got '%s'", string(original.Data()), string(reconstructed.Data()))
	}
	if reconstructed.Tr_State() != original.Tr_State() {
		t.Errorf("Transaction state mismatch: expected %v, got %v", original.Tr_State(), reconstructed.Tr_State())
	}
	if !strings.HasPrefix(reconstructed.Tr_Id(), original.Tr_Id()) {
		t.Errorf("Transaction ID mismatch: expected to start with '%s', got '%s'", original.Tr_Id(), reconstructed.Tr_Id())
	}
	if reconstructed.Tr_ErrMsg() != original.Tr_ErrMsg() {
		t.Errorf("Transaction error message mismatch: expected '%s', got '%s'", original.Tr_ErrMsg(), reconstructed.Tr_ErrMsg())
	}
	if reconstructed.Tr_StartTime() != original.Tr_StartTime() {
		t.Errorf("Transaction start time mismatch: expected %d, got %d", original.Tr_StartTime(), reconstructed.Tr_StartTime())
	}
}

func TestMultipleRoundTrips(t *testing.T) {
	resources := newMockResources()

	// Create initial message
	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("initial data"), true, false, 123, ifs.Empty, "", "", 0, 0)

	// Perform multiple marshal/unmarshal cycles
	for i := 0; i < 10; i++ {
		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed on iteration %d: %v", i, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed on iteration %d: %v", i, err)
		}

		// Verify data integrity after each cycle
		if string(newMsg.Data()) != "initial data" {
			t.Errorf("Data corruption detected on iteration %d", i)
		}
		if newMsg.Sequence() != 123 {
			t.Errorf("Sequence corruption detected on iteration %d", i)
		}

		// Use the reconstructed message for next iteration
		msg = newMsg
	}
}

func TestConcurrentMarshalUnmarshal(t *testing.T) {
	resources := newMockResources()

	// Test concurrent access to marshal/unmarshal operations
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			defer func() { done <- true }()

			msg := &ifs.Message{}
			msg.Init(
				fmt.Sprintf("dest-%d", index),
				fmt.Sprintf("service-%d", index),
				byte(index%256),
				ifs.Priority(index%8),
				ifs.M_All,
				ifs.Action((index%11)+1),
				fmt.Sprintf("source-%d", index),
				fmt.Sprintf("vnet-%d", index),
				[]byte(fmt.Sprintf("data-%d", index)),
				index%2 == 0,
				index%3 == 0,
				uint32(index*1000),
				ifs.TransactionState(index%14),
				fmt.Sprintf("tr-id-%d", index),
				fmt.Sprintf("tr-err-%d", index),
				int64(index*10000), 30,
			)

			data, err := msg.Marshal(nil, resources)
			if err != nil {
				t.Errorf("Concurrent marshal failed for goroutine %d: %v", index, err)
				return
			}

			newMsg := &ifs.Message{}
			_, err = newMsg.Unmarshal(data, resources)
			if err != nil {
				t.Errorf("Concurrent unmarshal failed for goroutine %d: %v", index, err)
				return
			}

			// Verify specific fields - check prefix due to fixed-size field padding
			expected := fmt.Sprintf("dest-%d", index)
			if !strings.HasPrefix(newMsg.Destination(), expected) {
				t.Errorf("Concurrent data corruption in goroutine %d: expected prefix '%s', got '%s'", index, expected, newMsg.Destination())
			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestMessageSizeCalculation(t *testing.T) {
	resources := newMockResources()

	testCases := []struct {
		name     string
		dataSize int
		trState  ifs.TransactionState
	}{
		{"empty_no_transaction", 0, ifs.Empty},
		{"small_no_transaction", 100, ifs.Empty},
		{"large_no_transaction", 10000, ifs.Empty},
		{"empty_with_transaction", 0, ifs.Locked},
		{"small_with_transaction", 100, ifs.Locked},
		{"large_with_transaction", 10000, ifs.Locked},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg := &ifs.Message{}
			data := ""
			if tc.dataSize > 0 {
				data = fmt.Sprintf("%*s", tc.dataSize, "x") // Create string of specified size
			}

			var trId, trErr string
			var trStart int64
			var trTimeout int64
			if tc.trState != ifs.Empty {
				trId = "transaction-id-12345678901234567890"
				trErr = "transaction error message"
				trStart = 1234567890
				trTimeout = 30
			}

			msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte(data), true, false, 123, tc.trState, trId, trErr, trStart, trTimeout)

			marshaledData, err := msg.Marshal(nil, resources)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			// Verify that marshaled size is reasonable
			expectedMinSize := 119 + tc.dataSize // Header + minimum body + data
			if len(marshaledData) < expectedMinSize {
				t.Errorf("Marshaled data too small: expected at least %d bytes, got %d", expectedMinSize, len(marshaledData))
			}

			// Verify successful unmarshal
			newMsg := &ifs.Message{}
			_, err = newMsg.Unmarshal(marshaledData, resources)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if len(newMsg.Data()) != tc.dataSize {
				t.Errorf("Data size mismatch: expected %d, got %d", tc.dataSize, len(newMsg.Data()))
			}
		})
	}
}

func BenchmarkMessageMarshal(b *testing.B) {
	resources := newMockResources()
	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("benchmark data"), true, false, 123, ifs.Empty, "", "", 0, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := msg.Marshal(nil, resources)
		if err != nil {
			b.Fatalf("Marshal failed: %v", err)
		}
	}
}

func BenchmarkMessageUnmarshal(b *testing.B) {
	resources := newMockResources()
	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("benchmark data"), true, false, 123, ifs.Empty, "", "", 0, 0)

	data, err := msg.Marshal(nil, resources)
	if err != nil {
		b.Fatalf("Setup marshal failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMsg := &ifs.Message{}
		_, err := newMsg.Unmarshal(data, resources)
		if err != nil {
			b.Fatalf("Unmarshal failed: %v", err)
		}
	}
}

func BenchmarkMessageRoundTrip(b *testing.B) {
	resources := newMockResources()
	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("benchmark data"), true, false, 123, ifs.Empty, "", "", 0, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := msg.Marshal(nil, resources)
		if err != nil {
			b.Fatalf("Marshal failed: %v", err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			b.Fatalf("Unmarshal failed: %v", err)
		}
	}
}

package tests

import (
	"strings"
	"testing"
	"time"

	"github.com/saichler/l8types/go/ifs"
)

func TestMessageMarshalUnmarshalBasic(t *testing.T) {
	resources := newMockResources()

	// Create a basic message
	msg := &ifs.Message{}
	msg.Init(
		"test-destination",
		"test-svc", // Keep service name <= 10 chars
		1,
		ifs.P1,
		ifs.M_All,
		ifs.POST,
		"test-source",
		"test-vnet",
		[]byte("test-data"),
		true,
		false,
		12345,
		ifs.NotATransaction,
		"",
		"",
		0, 0, 0, 0, 0,
	)

	// Marshal the message
	data, err := msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal the message
	newMsg := &ifs.Message{}
	_, err = newMsg.Unmarshal(data, resources)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify all fields - source, vnet, destination are padded to exactly 36 bytes with null bytes
	if len(newMsg.Source()) != 36 {
		t.Errorf("Source length mismatch: expected 36, got %d", len(newMsg.Source()))
	}
	if !strings.HasPrefix(newMsg.Source(), "test-source") {
		t.Errorf("Source prefix mismatch: expected to start with 'test-source', got '%s'", newMsg.Source())
	}
	if len(newMsg.Vnet()) != 36 {
		t.Errorf("Vnet length mismatch: expected 36, got %d", len(newMsg.Vnet()))
	}
	if !strings.HasPrefix(newMsg.Vnet(), "test-vnet") {
		t.Errorf("Vnet prefix mismatch: expected to start with 'test-vnet', got '%s'", newMsg.Vnet())
	}
	if len(newMsg.Destination()) != 36 {
		t.Errorf("Destination length mismatch: expected 36, got %d", len(newMsg.Destination()))
	}
	if !strings.HasPrefix(newMsg.Destination(), "test-destination") {
		t.Errorf("Destination prefix mismatch: expected to start with 'test-destination', got '%s'", newMsg.Destination())
	}
	// Service name preserves original length due to ToServiceName() behavior
	if newMsg.ServiceName() != "test-svc" {
		t.Errorf("ServiceName mismatch: expected 'test-svc', got '%s'", newMsg.ServiceName())
	}
	if newMsg.ServiceArea() != 1 {
		t.Errorf("ServiceArea mismatch: expected 1, got %d", newMsg.ServiceArea())
	}
	if newMsg.Priority() != ifs.P1 {
		t.Errorf("Priority mismatch: expected P1, got %v", newMsg.Priority())
	}
	if newMsg.Action() != ifs.POST {
		t.Errorf("Action mismatch: expected POST, got %v", newMsg.Action())
	}
	if string(newMsg.Data()) != "test-data" {
		t.Errorf("Data mismatch: expected 'test-data', got '%s'", string(newMsg.Data()))
	}
	if newMsg.Request() != true {
		t.Errorf("Request mismatch: expected true, got %v", newMsg.Request())
	}
	if newMsg.Reply() != false {
		t.Errorf("Reply mismatch: expected false, got %v", newMsg.Reply())
	}
	if newMsg.Sequence() != 12345 {
		t.Errorf("Sequence mismatch: expected 12345, got %d", newMsg.Sequence())
	}
	if newMsg.Tr_State() != ifs.NotATransaction {
		t.Errorf("Transaction state mismatch: expected Empty, got %v", newMsg.Tr_State())
	}
}

func TestMessageMarshalUnmarshalWithTransaction(t *testing.T) {
	resources := newMockResources()

	msg := &ifs.Message{}
	msg.Init(
		"dest",
		"service",
		2,
		ifs.P5,
		ifs.M_All,
		ifs.GET,
		"source",
		"vnet",
		[]byte("data"),
		false,
		true,
		67890,
		ifs.Running,
		"transaction-id-12345678901234567890123456",
		"transaction error message",
		time.Now().Unix(), time.Now().Unix()+10, time.Now().Unix()+20, time.Now().Unix()+30, 30,
	)

	data, err := msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	newMsg := &ifs.Message{}
	_, err = newMsg.Unmarshal(data, resources)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if newMsg.Tr_State() != ifs.Running {
		t.Errorf("Transaction state mismatch: expected Locked, got %v", newMsg.Tr_State())
	}
	// Transaction ID is also exactly 36 bytes
	if len(newMsg.Tr_Id()) != 36 {
		t.Errorf("Transaction ID length mismatch: expected 36, got %d", len(newMsg.Tr_Id()))
	}
	if !strings.HasPrefix(newMsg.Tr_Id(), "transaction-id-12345678901234567890123456"[:36]) {
		t.Errorf("Transaction ID prefix mismatch: got '%s'", newMsg.Tr_Id())
	}
	if newMsg.Tr_ErrMsg() != "transaction error message" {
		t.Errorf("Transaction error message mismatch: expected 'transaction error message', got '%s'", newMsg.Tr_ErrMsg())
	}
	if newMsg.Tr_Created() != msg.Tr_Created() {
		t.Errorf("Transaction created mismatch: expected %d, got %d", msg.Tr_Created(), newMsg.Tr_Created())
	}
	if newMsg.Tr_Queued() != msg.Tr_Queued() {
		t.Errorf("Transaction queued mismatch: expected %d, got %d", msg.Tr_Queued(), newMsg.Tr_Queued())
	}
	if newMsg.Tr_Running() != msg.Tr_Running() {
		t.Errorf("Transaction running mismatch: expected %d, got %d", msg.Tr_Running(), newMsg.Tr_Running())
	}
	if newMsg.Tr_End() != msg.Tr_End() {
		t.Errorf("Transaction end mismatch: expected %d, got %d", msg.Tr_End(), newMsg.Tr_End())
	}

	// Test tr_replica field
	msg.SetTr_Replica(byte(3))
	data, err = msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal with tr_replica failed: %v", err)
	}

	newMsg2 := &ifs.Message{}
	_, err = newMsg2.Unmarshal(data, resources)
	if err != nil {
		t.Fatalf("Unmarshal with tr_replica failed: %v", err)
	}

	if newMsg2.Tr_Replica() != byte(3) {
		t.Errorf("Transaction replica mismatch: expected 3, got %d", newMsg2.Tr_Replica())
	}
}

func TestMessageMarshalUnmarshalEmptyFields(t *testing.T) {
	resources := newMockResources()

	msg := &ifs.Message{}
	msg.Init("", "", 0, ifs.P8, ifs.M_All, ifs.Reply, "", "", []byte(""), false, false, 0, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

	data, err := msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	newMsg := &ifs.Message{}
	_, err = newMsg.Unmarshal(data, resources)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if newMsg.Destination() != "" {
		t.Errorf("Expected empty destination, got '%s'", newMsg.Destination())
	}
	if newMsg.ServiceName() != "" {
		t.Errorf("Expected empty service name, got '%s'", newMsg.ServiceName())
	}
	if string(newMsg.Data()) != "" {
		t.Errorf("Expected empty data, got '%s'", string(newMsg.Data()))
	}
}

func TestMessageMarshalUnmarshalLargeData(t *testing.T) {
	resources := newMockResources()

	// Create large data strings
	largeData := strings.Repeat("A", 10000)
	largeFailMessage := strings.Repeat("F", 255)

	msg := &ifs.Message{}
	msg.Init(
		"destination",
		"service",
		255,
		ifs.P1,
		ifs.M_All,
		ifs.DELETE,
		"source",
		"vnet",
		[]byte(largeData),
		true,
		true,
		4294967295, // max uint32
		ifs.Failed,
		"transaction-id-large-test-case-here",
		"large transaction error: "+strings.Repeat("E", 200),
		9223372036854775807, // max int64
		9223372036854775806,
		9223372036854775805,
		9223372036854775804,
		30,
	)
	msg.SetFailMessage(largeFailMessage)

	data, err := msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	newMsg := &ifs.Message{}
	_, err = newMsg.Unmarshal(data, resources)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if string(newMsg.Data()) != largeData {
		t.Errorf("Large data mismatch")
	}
	if newMsg.FailMessage() != largeFailMessage {
		t.Errorf("Large fail message mismatch")
	}
	if newMsg.Sequence() != 4294967295 {
		t.Errorf("Max sequence mismatch: expected 4294967295, got %d", newMsg.Sequence())
	}
	if newMsg.Tr_Created() != 9223372036854775807 {
		t.Errorf("Max created time mismatch: expected 9223372036854775807, got %d", newMsg.Tr_Created())
	}
	if newMsg.Tr_Queued() != 9223372036854775806 {
		t.Errorf("Max queued time mismatch: expected 9223372036854775806, got %d", newMsg.Tr_Queued())
	}
	if newMsg.Tr_Running() != 9223372036854775805 {
		t.Errorf("Max running time mismatch: expected 9223372036854775805, got %d", newMsg.Tr_Running())
	}
	if newMsg.Tr_End() != 9223372036854775804 {
		t.Errorf("Max end time mismatch: expected 9223372036854775804, got %d", newMsg.Tr_End())
	}
}

func TestMessageMarshalEncryptionError(t *testing.T) {
	resources := newMockResourcesWithError(true, false)

	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("data"), true, false, 123, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

	_, err := msg.Marshal(nil, resources)
	if err == nil {
		t.Fatalf("Expected encryption error, but marshal succeeded")
	}
	if !strings.Contains(err.Error(), "encryption failed") {
		t.Errorf("Expected encryption error message, got: %v", err)
	}
}

func TestMessageUnmarshalDecryptionError(t *testing.T) {
	resources := newMockResources()

	msg := &ifs.Message{}
	msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("data"), true, false, 123, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

	// Marshal with working encryption
	data, err := msg.Marshal(nil, resources)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Try to unmarshal with failing decryption
	resourcesWithError := newMockResourcesWithError(false, true)
	newMsg := &ifs.Message{}
	_, err = newMsg.Unmarshal(data, resourcesWithError)

	if err == nil {
		t.Fatalf("Expected decryption error, but unmarshal succeeded")
	}
	if !strings.Contains(err.Error(), "decryption failed") {
		t.Errorf("Expected decryption error message, got: %v", err)
	}
}

func TestAllActions(t *testing.T) {
	resources := newMockResources()

	actions := []ifs.Action{ifs.POST, ifs.PUT, ifs.PATCH, ifs.DELETE, ifs.GET, ifs.Reply, ifs.Notify, ifs.Sync, ifs.EndPoints}

	for _, action := range actions {
		msg := &ifs.Message{}
		msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, action, "source", "vnet", []byte("data"), true, false, 123, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed for action %v: %v", action, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed for action %v: %v", action, err)
		}

		if newMsg.Action() != action {
			t.Errorf("Action mismatch for %v: expected %v, got %v", action, action, newMsg.Action())
		}
	}
}

func TestAllPriorities(t *testing.T) {
	resources := newMockResources()

	priorities := []ifs.Priority{ifs.P8, ifs.P7, ifs.P6, ifs.P5, ifs.P4, ifs.P3, ifs.P2, ifs.P1}

	for _, priority := range priorities {
		msg := &ifs.Message{}
		msg.Init("dest", "service", 1, priority, ifs.M_All, ifs.POST, "source", "vnet", []byte("data"), true, false, 123, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed for priority %v: %v", priority, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed for priority %v: %v", priority, err)
		}

		if newMsg.Priority() != priority {
			t.Errorf("Priority mismatch for %v: expected %v, got %v", priority, priority, newMsg.Priority())
		}
	}
}

func TestAllTransactionStates(t *testing.T) {
	resources := newMockResources()

	states := []ifs.TransactionState{
		ifs.NotATransaction, ifs.Created, ifs.Queued, ifs.Running,
		ifs.Committed, ifs.Rollback, ifs.Failed, ifs.Cleanup,
	}

	for _, state := range states {
		msg := &ifs.Message{}
		var trId, trErr string
		var trCreated, trQueued, trRunning, trEnd, trTimeout int64
		// Only set transaction fields for non-Empty states
		if state != ifs.NotATransaction {
			trId = "tr-id-123456789012345678901234567890"
			trErr = "transaction error message"
			trCreated = 1234567890
			trQueued = 1234567900
			trRunning = 1234567910
			trEnd = 1234567920
			trTimeout = 30
		}

		msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("data"), true, false, 123, state, trId, trErr, trCreated, trQueued, trRunning, trEnd, trTimeout)

		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed for transaction state %v: %v", state, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed for transaction state %v: %v", state, err)
		}

		if newMsg.Tr_State() != state {
			t.Errorf("Transaction state mismatch for %v: expected %v, got %v", state, state, newMsg.Tr_State())
		}

		if state != ifs.NotATransaction {
			if newMsg.Tr_Id() != trId {
				t.Errorf("Transaction ID mismatch for state %v: expected '%s', got '%s'", state, trId, newMsg.Tr_Id())
			}
			if newMsg.Tr_ErrMsg() != trErr {
				t.Errorf("Transaction error message mismatch for state %v: expected '%s', got '%s'", state, trErr, newMsg.Tr_ErrMsg())
			}
			if newMsg.Tr_Created() != trCreated {
				t.Errorf("Transaction created mismatch for state %v: expected %d, got %d", state, trCreated, newMsg.Tr_Created())
			}
			if newMsg.Tr_Queued() != trQueued {
				t.Errorf("Transaction queued mismatch for state %v: expected %d, got %d", state, trQueued, newMsg.Tr_Queued())
			}
			if newMsg.Tr_Running() != trRunning {
				t.Errorf("Transaction running mismatch for state %v: expected %d, got %d", state, trRunning, newMsg.Tr_Running())
			}
			if newMsg.Tr_End() != trEnd {
				t.Errorf("Transaction end mismatch for state %v: expected %d, got %d", state, trEnd, newMsg.Tr_End())
			}

			// Test tr_replica for transaction states
			msg.SetTr_Replica(byte(5))
			data, err = msg.Marshal(nil, resources)
			if err != nil {
				t.Fatalf("Marshal with tr_replica failed for state %v: %v", state, err)
			}

			testMsg := &ifs.Message{}
			_, err = testMsg.Unmarshal(data, resources)
			if err != nil {
				t.Fatalf("Unmarshal with tr_replica failed for state %v: %v", state, err)
			}

			if testMsg.Tr_Replica() != byte(5) {
				t.Errorf("Transaction replica mismatch for state %v: expected 5, got %d", state, testMsg.Tr_Replica())
			}
		}
	}
}

func TestBoolCombinations(t *testing.T) {
	resources := newMockResources()

	combinations := []struct {
		request bool
		reply   bool
	}{
		{false, false},
		{true, false},
		{false, true},
		{true, true},
	}

	for _, combo := range combinations {
		msg := &ifs.Message{}
		msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, ifs.POST, "source", "vnet", []byte("data"), combo.request, combo.reply, 123, ifs.NotATransaction, "", "", 0, 0, 0, 0, 0)

		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed for request=%v, reply=%v: %v", combo.request, combo.reply, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed for request=%v, reply=%v: %v", combo.request, combo.reply, err)
		}

		if newMsg.Request() != combo.request {
			t.Errorf("Request mismatch: expected %v, got %v", combo.request, newMsg.Request())
		}
		if newMsg.Reply() != combo.reply {
			t.Errorf("Reply mismatch: expected %v, got %v", combo.reply, newMsg.Reply())
		}
	}
}

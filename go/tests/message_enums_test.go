package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestTransactionStateString(t *testing.T) {
	tests := []struct {
		state    ifs.TransactionState
		expected string
	}{
		{ifs.Empty, "Unknown"},
		{ifs.Create, "Create"},
		{ifs.Created, "Created"},
		{ifs.Start, "Start"},
		{ifs.Lock, "Lock"},
		{ifs.Locked, "Locked"},
		{ifs.LockFailed, "LockFailed"},
		{ifs.Commit, "Commit"},
		{ifs.Commited, "Commited"},
		{ifs.Rollback, "Rollback"},
		{ifs.Rollbacked, "Rollbacked"},
		{ifs.Finish, "Finish"},
		{ifs.Finished, "Finished"},
		{ifs.Errored, "Errored"},
		{ifs.TransactionState(99), "Unknown"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("For state %d, expected '%s', got '%s'", test.state, test.expected, result)
		}
	}
}

func TestPriorityConstants(t *testing.T) {
	// Test that priority constants are defined correctly
	if ifs.P8 != 0 {
		t.Errorf("Expected P8 to be 0, got %d", ifs.P8)
	}
	if ifs.P1 != 7 {
		t.Errorf("Expected P1 to be 7, got %d", ifs.P1)
	}
}

func TestMulticastModeConstants(t *testing.T) {
	// Test that multicast mode constants are defined correctly
	if ifs.M_All != 0 {
		t.Errorf("Expected M_All to be 0, got %d", ifs.M_All)
	}
	if ifs.M_Unicast != 128 {
		t.Errorf("Expected M_Unicast to be 128, got %d", ifs.M_Unicast)
	}
}

func TestActionConstants(t *testing.T) {
	// Test that action constants are defined correctly
	actions := map[ifs.Action]string{
		ifs.POST:      "POST",
		ifs.PUT:       "PUT",
		ifs.PATCH:     "PATCH",
		ifs.DELETE:    "DELETE",
		ifs.GET:       "GET",
		ifs.Reply:     "Reply",
		ifs.Notify:    "Notify",
		ifs.Sync:      "Sync",
		ifs.EndPoints: "EndPoints",
	}

	for action := range actions {
		if action < 1 || action > 9 {
			t.Errorf("Action %v is out of expected range", action)
		}
	}
}

func TestDestinationConstants(t *testing.T) {
	// Test that destination constants are defined
	if ifs.DESTINATION_Single == "" {
		t.Error("DESTINATION_Single should not be empty")
	}
	if ifs.DESTINATION_Leader == "" {
		t.Error("DESTINATION_Leader should not be empty")
	}
	if len(ifs.DESTINATION_Single) != 36 {
		t.Errorf("Expected DESTINATION_Single length 36, got %d", len(ifs.DESTINATION_Single))
	}
	if len(ifs.DESTINATION_Leader) != 36 {
		t.Errorf("Expected DESTINATION_Leader length 36, got %d", len(ifs.DESTINATION_Leader))
	}
}

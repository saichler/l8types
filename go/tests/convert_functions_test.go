package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestBytes2LongAndLong2Bytes(t *testing.T) {
	testCases := []int64{
		0,
		1,
		-1,
		255,
		256,
		65535,
		65536,
		16777215,
		16777216,
		4294967295,
		4294967296,
		9223372036854775807,  // max int64
		-9223372036854775808, // min int64
	}

	for _, original := range testCases {
		bytes := ifs.Long2Bytes(original)
		if len(bytes) != 8 {
			t.Errorf("Long2Bytes should return 8 bytes, got %d for value %d", len(bytes), original)
		}

		result := ifs.Bytes2Long(bytes)
		if result != original {
			t.Errorf("Long roundtrip failed: expected %d, got %d", original, result)
		}
	}
}

func TestBytes2IntAndInt2Bytes(t *testing.T) {
	testCases := []int32{
		0,
		1,
		-1,
		255,
		256,
		65535,
		65536,
		16777215,
		16777216,
		2147483647,  // max int32
		-2147483648, // min int32
	}

	for _, original := range testCases {
		bytes := ifs.Int2Bytes(original)
		if len(bytes) != 4 {
			t.Errorf("Int2Bytes should return 4 bytes, got %d for value %d", len(bytes), original)
		}

		result := ifs.Bytes2Int(bytes)
		if result != original {
			t.Errorf("Int roundtrip failed: expected %d, got %d", original, result)
		}
	}
}

func TestBytes2UInt16AndUInt162Bytes(t *testing.T) {
	testCases := []uint16{
		0,
		1,
		255,
		256,
		65535, // max uint16
	}

	for _, original := range testCases {
		bytes := ifs.UInt162Bytes(original)
		if len(bytes) != 2 {
			t.Errorf("UInt162Bytes should return 2 bytes, got %d for value %d", len(bytes), original)
		}

		result := ifs.Bytes2UInt16(bytes)
		if result != original {
			t.Errorf("UInt16 roundtrip failed: expected %d, got %d", original, result)
		}
	}
}

func TestBytes2UInt32AndUInt322Bytes(t *testing.T) {
	testCases := []uint32{
		0,
		1,
		255,
		256,
		65535,
		65536,
		16777215,
		16777216,
		4294967295, // max uint32
	}

	for _, original := range testCases {
		bytes := ifs.UInt322Bytes(original)
		if len(bytes) != 4 {
			t.Errorf("UInt322Bytes should return 4 bytes, got %d for value %d", len(bytes), original)
		}

		result := ifs.Bytes2UInt32(bytes)
		if result != original {
			t.Errorf("UInt32 roundtrip failed: expected %d, got %d", original, result)
		}
	}
}

func TestBoolsAndBoolOf(t *testing.T) {
	testCases := []struct {
		request  bool
		reply    bool
		expected byte
	}{
		{false, false, 0},
		{true, false, 1},
		{false, true, 2},
		{true, true, 3},
	}

	for _, tc := range testCases {
		result := ifs.Bools(tc.request, tc.reply)
		if result != tc.expected {
			t.Errorf("Bools(%v, %v) expected %d, got %d", tc.request, tc.reply, tc.expected, result)
		}

		// Test reverse conversion
		request, reply := ifs.BoolOf(tc.expected)
		if request != tc.request || reply != tc.reply {
			t.Errorf("BoolOf(%d) expected (%v, %v), got (%v, %v)", tc.expected, tc.request, tc.reply, request, reply)
		}
	}
}

func TestBoolOfPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("BoolOf should panic for invalid input")
		}
	}()

	ifs.BoolOf(4) // Should panic
}

func TestByteToActionState(t *testing.T) {
	// Test ByteToActionState function directly with known byte values
	testCases := []struct {
		input  byte
		action ifs.Action
		state  ifs.TransactionState
	}{
		{0x00, ifs.Action(0), ifs.TransactionState(0)},   // 0000 0000
		{0x11, ifs.Action(1), ifs.TransactionState(1)},   // 0001 0001
		{0x23, ifs.Action(2), ifs.TransactionState(3)},   // 0010 0011
		{0x45, ifs.Action(4), ifs.TransactionState(5)},   // 0100 0101
		{0x67, ifs.Action(6), ifs.TransactionState(7)},   // 0110 0111
		{0x89, ifs.Action(8), ifs.TransactionState(9)},   // 1000 1001
		{0xAB, ifs.Action(10), ifs.TransactionState(11)}, // 1010 1011
		{0xCD, ifs.Action(12), ifs.TransactionState(13)}, // 1100 1101
		{0xFF, ifs.Action(15), ifs.TransactionState(15)}, // 1111 1111
	}

	for _, tc := range testCases {
		action, state := ifs.ByteToActionState(tc.input)

		if action != tc.action {
			t.Errorf("Action decode failed for byte 0x%02X: expected %v, got %v", tc.input, tc.action, action)
		}
		if state != tc.state {
			t.Errorf("State decode failed for byte 0x%02X: expected %v, got %v", tc.input, tc.state, state)
		}
	}
}

func TestActionStateBitManipulation(t *testing.T) {
	// Test that we can properly extract action and state from known combinations
	// Since actionStateToByte is not exported, we test through marshal/unmarshal
	resources := newMockResources()

	testCases := []struct {
		action ifs.Action
		state  ifs.TransactionState
	}{
		{ifs.POST, ifs.Empty},
		{ifs.GET, ifs.Locked},
		{ifs.PUT, ifs.Commited},
		{ifs.DELETE, ifs.Errored},
		{ifs.PATCH, ifs.Finished},
		{ifs.Reply, ifs.Create},
		{ifs.Notify, ifs.Start},
		{ifs.Sync, ifs.Lock},
		{ifs.EndPoints, ifs.Commit},
		{ifs.Sync, ifs.Rollback},
		{ifs.Sync, ifs.Finish},
	}

	for _, tc := range testCases {
		msg := &ifs.Message{}
		var trId, trErr string
		var trStart int64
		var trTimeout int64

		if tc.state != ifs.Empty {
			trId = "tr-id"
			trErr = "tr-err"
			trStart = 123
			trTimeout = 30
		}

		msg.Init("dest", "service", 1, ifs.P1, ifs.M_All, tc.action, "source", "vnet", []byte("data"), true, false, 123, tc.state, trId, trErr, trStart, trTimeout)

		data, err := msg.Marshal(nil, resources)
		if err != nil {
			t.Fatalf("Marshal failed for action %v, state %v: %v", tc.action, tc.state, err)
		}

		newMsg := &ifs.Message{}
		_, err = newMsg.Unmarshal(data, resources)
		if err != nil {
			t.Fatalf("Unmarshal failed for action %v, state %v: %v", tc.action, tc.state, err)
		}

		if newMsg.Action() != tc.action {
			t.Errorf("Action roundtrip failed: expected %v, got %v", tc.action, newMsg.Action())
		}
		if newMsg.Tr_State() != tc.state {
			t.Errorf("State roundtrip failed: expected %v, got %v", tc.state, newMsg.Tr_State())
		}
	}
}

// Test specific byte patterns
func TestConversionBytePatterns(t *testing.T) {
	// Test Long2Bytes with specific pattern
	value := int64(0x0102030405060708)
	bytes := ifs.Long2Bytes(value)
	expected := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	for i, b := range bytes {
		if b != expected[i] {
			t.Errorf("Long2Bytes byte pattern mismatch at index %d: expected 0x%02x, got 0x%02x", i, expected[i], b)
		}
	}

	// Test UInt322Bytes with specific pattern
	uint32Value := uint32(0x01020304)
	uint32Bytes := ifs.UInt322Bytes(uint32Value)
	expectedUint32 := []byte{0x01, 0x02, 0x03, 0x04}

	for i, b := range uint32Bytes {
		if b != expectedUint32[i] {
			t.Errorf("UInt322Bytes byte pattern mismatch at index %d: expected 0x%02x, got 0x%02x", i, expectedUint32[i], b)
		}
	}

	// Test UInt162Bytes with specific pattern
	uint16Value := uint16(0x0102)
	uint16Bytes := ifs.UInt162Bytes(uint16Value)
	expectedUint16 := []byte{0x01, 0x02}

	for i, b := range uint16Bytes {
		if b != expectedUint16[i] {
			t.Errorf("UInt162Bytes byte pattern mismatch at index %d: expected 0x%02x, got 0x%02x", i, expectedUint16[i], b)
		}
	}
}

// Test endianness consistency
func TestEndiannessConsistency(t *testing.T) {
	// Test that all conversion functions use the same endianness (big-endian)

	// For a multi-byte value, the most significant byte should come first
	value := uint32(0x12345678)
	bytes := ifs.UInt322Bytes(value)

	if bytes[0] != 0x12 {
		t.Errorf("Expected big-endian format: first byte should be 0x12, got 0x%02x", bytes[0])
	}
	if bytes[3] != 0x78 {
		t.Errorf("Expected big-endian format: last byte should be 0x78, got 0x%02x", bytes[3])
	}

	// Test with int64
	longValue := int64(0x123456789ABCDEF0)
	longBytes := ifs.Long2Bytes(longValue)

	if longBytes[0] != 0x12 {
		t.Errorf("Expected big-endian format for long: first byte should be 0x12, got 0x%02x", longBytes[0])
	}
	if longBytes[7] != 0xF0 {
		t.Errorf("Expected big-endian format for long: last byte should be 0xF0, got 0x%02x", longBytes[7])
	}
}

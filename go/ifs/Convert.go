package ifs

// Optimized Convert.go - Maintaining full API compatibility with performance improvements
// Key optimizations:
// - Bitwise OR instead of addition for better performance
// - Slice literals instead of make() + assignments for better memory efficiency
// - Optimized boolean operations with bitwise logic
// - Improved action/state bit manipulation

// Bytes2Long converts 8 bytes to int64 (optimized with bitwise OR)
func Bytes2Long(data []byte) int64 {
	return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7])
}

// Long2Bytes converts int64 to 8 bytes (optimized with slice literal)
func Long2Bytes(s int64) []byte {
	return []byte{
		byte(s >> 56),
		byte(s >> 48),
		byte(s >> 40),
		byte(s >> 32),
		byte(s >> 24),
		byte(s >> 16),
		byte(s >> 8),
		byte(s),
	}
}

// Bytes2Int converts 4 bytes to int32 (optimized with bitwise OR)
func Bytes2Int(data []byte) int32 {
	return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3])
}

// Int2Bytes converts int32 to 4 bytes (optimized with slice literal)
func Int2Bytes(s int32) []byte {
	return []byte{
		byte(s >> 24),
		byte(s >> 16),
		byte(s >> 8),
		byte(s),
	}
}

// Bytes2UInt16 converts 2 bytes to uint16 (optimized with bitwise OR)
func Bytes2UInt16(data []byte) uint16 {
	return uint16(data[0])<<8 | uint16(data[1])
}

// UInt162Bytes converts uint16 to 2 bytes (optimized with slice literal)
func UInt162Bytes(s uint16) []byte {
	return []byte{byte(s >> 8), byte(s)}
}

// Bytes2UInt32 converts 4 bytes to uint32 (optimized with bitwise OR)
func Bytes2UInt32(data []byte) uint32 {
	return uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
}

// UInt322Bytes converts uint32 to 4 bytes (optimized with slice literal)
func UInt322Bytes(s uint32) []byte {
	return []byte{
		byte(s >> 24),
		byte(s >> 16),
		byte(s >> 8),
		byte(s),
	}
}

// Bools converts two booleans to a byte (optimized with bitwise operations)
func Bools(request, reply bool) byte {
	var result byte
	if request {
		result |= 1
	}
	if reply {
		result |= 2
	}
	return result
}

// BoolOf converts a byte to two booleans (optimized with bitwise operations)
func BoolOf(b byte) (bool, bool) {
	if b > 3 {
		panic("Unexpected " + string(b+48))
	}
	return (b & 1) != 0, (b & 2) != 0
}

// actionStateToByte now returns action as a full byte (no packing)
// Deprecated: Use byte(action) directly
func actionStateToByte(action Action, trState TransactionState) byte {
	return byte(action)
}

// ByteToActionState unpacks action and transaction state from two separate bytes
// Deprecated: Read action and state as separate bytes directly
func ByteToActionState(actionByte byte, stateByte byte) (Action, TransactionState) {
	return Action(actionByte), TransactionState(stateByte)
}

func priorityMulticastModeToByte(priority Priority, multicastMode MulticastMode) byte {
	return (byte(priority) << 4) | byte(multicastMode)
}

func ByteToPriorityMulticastMode(b byte) (Priority, MulticastMode) {
	priority := Priority(b >> 4)             // Upper 4 bits
	multicastMode := MulticastMode(b & 0x0F) // Lower 4 bits (direct mask vs XOR)
	return priority, multicastMode
}

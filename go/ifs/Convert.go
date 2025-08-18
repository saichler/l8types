package ifs

func Bytes2Long(data []byte) int64 {
	v1 := int64(data[0]) << 56
	v2 := int64(data[1]) << 48
	v3 := int64(data[2]) << 40
	v4 := int64(data[3]) << 32
	v5 := int64(data[4]) << 24
	v6 := int64(data[5]) << 16
	v7 := int64(data[6]) << 8
	v8 := int64(data[7])
	return v1 + v2 + v3 + v4 + v5 + v6 + v7 + v8
}

func Long2Bytes(s int64) []byte {
	size := make([]byte, 8)
	size[7] = byte(s)
	size[6] = byte(s >> 8)
	size[5] = byte(s >> 16)
	size[4] = byte(s >> 24)
	size[3] = byte(s >> 32)
	size[2] = byte(s >> 40)
	size[1] = byte(s >> 48)
	size[0] = byte(s >> 56)
	return size
}

func Bytes2Int(data []byte) int32 {
	v1 := int32(data[0]) << 24
	v2 := int32(data[1]) << 16
	v3 := int32(data[2]) << 8
	v4 := int32(data[3])
	return v1 + v2 + v3 + v4
}

func Int2Bytes(s int32) []byte {
	size := make([]byte, 4)
	size[3] = byte(s)
	size[2] = byte(s >> 8)
	size[1] = byte(s >> 16)
	size[0] = byte(s >> 24)
	return size
}

func Bytes2UInt16(data []byte) uint16 {
	v1 := uint16(data[0]) << 8
	v2 := uint16(data[1])
	return v1 + v2
}

func UInt162Bytes(s uint16) []byte {
	size := make([]byte, 2)
	size[1] = byte(s)
	size[0] = byte(s >> 8)
	return size
}

func Bytes2UInt32(data []byte) uint32 {
	v1 := uint32(data[0]) << 24
	v2 := uint32(data[1]) << 16
	v3 := uint32(data[2]) << 8
	v4 := uint32(data[3])
	return v1 + v2 + v3 + v4
}

func UInt322Bytes(s uint32) []byte {
	size := make([]byte, 4)
	size[3] = byte(s)
	size[2] = byte(s >> 8)
	size[1] = byte(s >> 16)
	size[0] = byte(s >> 24)
	return size
}

func Bools(request, reply bool) byte {
	if !request && !reply {
		return 0
	} else if request && !reply {
		return 1
	} else if !request {
		return 2
	} else {
		return 3
	}
}

func BoolOf(b byte) (bool, bool) {
	switch b {
	case 0:
		return false, false
	case 1:
		return true, false
	case 2:
		return false, true
	case 3:
		return true, true
	}
	panic("Unexpected " + string(b+48))
}

func actionStateToByte(action Action, trState TransactionState) byte {
	b := byte(action) << 4
	b = b | byte(trState)
	return b
}

func ByteToActionState(b byte) (Action, TransactionState) {
	action := Action(b >> 4)
	ab := byte(action) << 4
	state := TransactionState(b ^ ab)
	return action, state
}

func putUInt32(buf []byte, offset int, val uint32) {
	buf[offset] = byte(val >> 24)
	buf[offset+1] = byte(val >> 16)
	buf[offset+2] = byte(val >> 8)
	buf[offset+3] = byte(val)
}

func putUInt16(buf []byte, offset int, val uint16) {
	buf[offset] = byte(val >> 8)
	buf[offset+1] = byte(val)
}

func putInt64(buf []byte, offset int, val int64) {
	buf[offset] = byte(val >> 56)
	buf[offset+1] = byte(val >> 48)
	buf[offset+2] = byte(val >> 40)
	buf[offset+3] = byte(val >> 32)
	buf[offset+4] = byte(val >> 24)
	buf[offset+5] = byte(val >> 16)
	buf[offset+6] = byte(val >> 8)
	buf[offset+7] = byte(val)
}

func encodeBools(request, reply bool) byte {
	if request && reply {
		return 3
	} else if reply {
		return 2
	} else if request {
		return 1
	}
	return 0
}

func getUInt32(buf []byte, offset int) uint32 {
	return uint32(buf[offset])<<24 | uint32(buf[offset+1])<<16 | uint32(buf[offset+2])<<8 | uint32(buf[offset+3])
}

func getUInt16(buf []byte, offset int) uint16 {
	return uint16(buf[offset])<<8 | uint16(buf[offset+1])
}

func getInt64(buf []byte, offset int) int64 {
	return int64(buf[offset])<<56 | int64(buf[offset+1])<<48 | int64(buf[offset+2])<<40 | int64(buf[offset+3])<<32 |
		int64(buf[offset+4])<<24 | int64(buf[offset+5])<<16 | int64(buf[offset+6])<<8 | int64(buf[offset+7])
}

func stringFromBytes(data []byte, start, end int) string {
	if start >= end || start >= len(data) {
		return ""
	}
	if end > len(data) {
		end = len(data)
	}
	
	// Find the actual end by trimming null bytes and spaces
	actualEnd := end
	for actualEnd > start && (data[actualEnd-1] == 0 || data[actualEnd-1] == ' ') {
		actualEnd--
	}
	
	if actualEnd == start {
		return ""
	}
	
	return string(data[start:actualEnd])
}

func nullTerminatedString(data []byte, start, maxLen int) string {
	end := start + maxLen
	if end > len(data) {
		end = len(data)
	}
	
	for i := start; i < end; i++ {
		if data[i] == 0 {
			if i == start {
				return ""
			}
			return string(data[start:i])
		}
	}
	return string(data[start:end])
}

package stream

// 读取bool值
func (ps *PacketStream) ReadBool() bool {
	if ps.ReadUInt8() > 0 {
		return true
	}
	return false
}

// 读取 Uint8 as byte
func (ps *PacketStream) ReadUInt8() (data uint8) {
	if ps.checkLen(1) {
		data = uint8(ps.Data[ps.Index])
		ps.Index++
		return
	}
	return
}

// 读取 Uint16
func (ps *PacketStream) ReadUInt16() (data uint16) {
	if ps.checkLen(2) {
		data = uint16(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 2
		return
	}
	return
}

// 读取 Uint32
func (ps *PacketStream) ReadUInt32() (data uint32) {
	if ps.checkLen(4) {
		data = uint32(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 4
		return
	}
	return
}

// 读取 Uint64
func (ps *PacketStream) ReadUInt64() (data uint64) {
	if ps.checkLen(8) {
		data = uint64(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 8
		return
	}
	return
}

// 读取 int8 as byte
func (ps *PacketStream) ReadInt8() (data int8) {
	if ps.checkLen(1) {
		data = int8(ps.Data[ps.Index])
		ps.Index++
		return
	}
	return
}

// 读取 int16
func (ps *PacketStream) ReadInt16() (data int16) {
	if ps.checkLen(2) {
		data = int16(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 2
		return
	}
	return
}

// 读取 int32
func (ps *PacketStream) ReadInt32() (data int32) {
	if ps.checkLen(4) {
		data = int32(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 4
		return
	}
	return
}

// 读取 int64
func (ps *PacketStream) ReadInt64() (data int64) {
	if ps.checkLen(8) {
		data = int64(BytesToUint64(ps.Data[ps.Index : ps.Index+2]))
		ps.Index += 8
		return
	}
	return
}

// int float

// 读取 Float32
func (ps *PacketStream) ReadFloat32() (data float32) {
	if ps.checkLen(4) {
		data = HexToFloat32(ps.Data[ps.Index : ps.Index+4])
		ps.Index += 4
		return
	}
	return
}

// 读取 Float64
func (ps *PacketStream) ReadFloat64() (data float64) {
	if ps.checkLen(8) {
		data = HexToFloat64(ps.Data[ps.Index : ps.Index+8])
		ps.Index += 8
		return
	}
	return
}

// 读取可变长度字符串
func (ps *PacketStream) ReadString() (data string) {
	// 首先读取uint16 长度的字符串长度
	if ps.checkLen(2) {
		length := ps.Data[ps.Index : ps.Index+2]
		ps.Index += 2
		data = ps.ReadStringL(uint16(BytesToUint64(length)))
		return
	}
	return
}

// 读取固定长度字符串
func (ps *PacketStream) ReadStringL(length uint16) (data string) {
	if ps.checkLen(length) {
		data = string(ps.Data[ps.Index : ps.Index+length])
		ps.Index += length
		return
	}
	return
}

// 检查数据长度是否足够
func (ps *PacketStream) checkLen(length uint16) bool {
	if uint16(len(ps.Data)) >= ps.Index+length {
		return true
	}
	return false
}

// Bool
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Uintptr
//	Float32
//	Float64
//	Complex64
//	Complex128
//	Array
//	Chan
//	Func
//	Interface
//	Map
//	Ptr
//	Slice
//	String
//	Struct

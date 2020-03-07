package stream

import (
	"encoding/binary"
	"github.com/yanlong-li/HelloWorld-GO/io/network/stream"
)

// 读取bool值
func (ps *SocketPacketStream) ReadBool() bool {
	if ps.ReadUInt8() > 0 {
		return true
	}
	return false
}

// 读取 Uint8 as byte
func (ps *SocketPacketStream) ReadUInt8() (data uint8) {
	if ps.checkLen(1) {
		_data := ps.GetData()
		data = _data[ps.Index]
		ps.Index++
		return
	}
	return
}

// 读取 Uint16
func (ps *SocketPacketStream) ReadUInt16() (data uint16) {
	if ps.checkLen(2) {
		_data := ps.GetData()
		data = binary.LittleEndian.Uint16(_data[ps.Index : ps.Index+2])
		ps.Index += 2
		return
	}
	return
}

// 读取 Uint32
func (ps *SocketPacketStream) ReadUInt32() (data uint32) {
	if ps.checkLen(4) {
		_data := ps.GetData()
		data = binary.LittleEndian.Uint32(_data[ps.Index : ps.Index+4])
		ps.Index += 4
		return
	}
	return
}

// 读取 Uint64
func (ps *SocketPacketStream) ReadUInt64() (data uint64) {
	if ps.checkLen(8) {
		_data := ps.GetData()
		data = binary.LittleEndian.Uint64(_data[ps.Index : ps.Index+8])
		ps.Index += 8
		return
	}
	return
}

// 读取 int8 as byte
func (ps *SocketPacketStream) ReadInt8() (data int8) {
	if ps.checkLen(1) {
		_data := ps.GetData()
		data = int8(_data[ps.Index])
		ps.Index++
		return
	}
	return
}

// 读取 int16
func (ps *SocketPacketStream) ReadInt16() (data int16) {
	if ps.checkLen(2) {
		_data := ps.GetData()
		data = int16(binary.LittleEndian.Uint16(_data[ps.Index : ps.Index+2]))
		ps.Index += 2
		return
	}
	return
}

// 读取 int32
func (ps *SocketPacketStream) ReadInt32() (data int32) {
	if ps.checkLen(4) {
		_data := ps.GetData()
		data = int32(binary.LittleEndian.Uint32(_data[ps.Index : ps.Index+4]))
		ps.Index += 4
		return
	}
	return
}

// 读取 int64
func (ps *SocketPacketStream) ReadInt64() (data int64) {
	if ps.checkLen(8) {
		_data := ps.GetData()
		data = int64(binary.LittleEndian.Uint64(_data[ps.Index : ps.Index+8]))
		ps.Index += 8
		return
	}
	return
}

// int float

// 读取 Float32
func (ps *SocketPacketStream) ReadFloat32() (data float32) {
	if ps.checkLen(4) {
		_data := ps.GetData()
		data = stream.BytesToFloat32(_data[ps.Index : ps.Index+4])
		ps.Index += 4
		return
	}
	return
}

// 读取 Float64
func (ps *SocketPacketStream) ReadFloat64() (data float64) {
	if ps.checkLen(8) {
		_data := ps.GetData()
		data = stream.BytesToFloat64(_data[ps.Index : ps.Index+8])
		ps.Index += 8
		return
	}
	return
}

// 读取可变长度字符串
func (ps *SocketPacketStream) ReadString() (data string) {
	// 首先读取uint16 长度的字符串长度
	if ps.checkLen(2) {
		_data := ps.GetData()
		length := _data[ps.Index : ps.Index+2]
		ps.Index += 2
		data = ps.ReadStringL(binary.LittleEndian.Uint16(length))
		return
	}
	return
}

// 读取固定长度字符串
func (ps *SocketPacketStream) ReadStringL(length uint16) (data string) {
	if ps.checkLen(length) {
		_data := ps.GetData()
		data = string(_data[ps.Index : ps.Index+length])
		ps.Index += length
		return
	}
	return
}

// 检查数据长度是否足够
func (ps *SocketPacketStream) checkLen(length uint16) bool {
	if uint16(len(ps.GetData())) >= ps.Index+length {
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

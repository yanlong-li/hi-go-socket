package stream

import "github.com/yanlong-li/hi-go-socket/stream"

func (ps *SocketPacketStream) WriteBool(data bool) {
	if data {
		ps.SetData(append(ps.GetData(), byte(1)))
	} else {
		ps.SetData(append(ps.GetData(), byte(0)))
	}
}

func (ps *SocketPacketStream) WriteUint8(n uint8) {
	ps.SetData(append(ps.GetData(), n))
}

func (ps *SocketPacketStream) WriteUint16(n uint16) {
	ps.SetData(append(ps.GetData(), stream.Uint16ToBytes(n)...))
}

func (ps *SocketPacketStream) WriteUint32(n uint32) {
	ps.SetData(append(ps.GetData(), stream.Uint32ToBytes(n)...))
}

func (ps *SocketPacketStream) WriteUint64(n uint64) {
	ps.SetData(append(ps.GetData(), stream.Uint64ToBytes(n)...))
}

func (ps *SocketPacketStream) WriteInt8(n int8) {
	ps.SetData(append(ps.GetData(), byte(n)))
}

func (ps *SocketPacketStream) WriteInt16(n int16) {
	ps.SetData(append(ps.GetData(), stream.Uint16ToBytes(uint16(n))...))
}

func (ps *SocketPacketStream) WriteInt32(n int32) {
	ps.SetData(append(ps.GetData(), stream.Uint32ToBytes(uint32(n))...))
}

func (ps *SocketPacketStream) WriteInt64(n int64) {
	ps.SetData(append(ps.GetData(), stream.Uint64ToBytes(uint64(n))...))
}

func (ps *SocketPacketStream) WriteFloat32(n float32) {
	ps.SetData(append(ps.GetData(), stream.Float32ToBytes(n)...))
}

func (ps *SocketPacketStream) WriteFloat64(n float64) {
	ps.SetData(append(ps.GetData(), stream.Float64ToBytes(n)...))
}

func (ps *SocketPacketStream) WriteString(n string) {
	length := len(n)
	ps.WriteUint16(uint16(length))
	ps.SetData(append(ps.GetData(), n...))
}

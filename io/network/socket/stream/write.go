package stream

func (ps *SocketPacketStream) WriteBool(data bool) {
	if data {
		ps.Data = append(ps.Data, byte(1))
	} else {
		ps.Data = append(ps.Data, byte(0))
	}
}

func (ps *SocketPacketStream) WriteUint8(n uint8) {
	ps.Data = append(ps.Data, n)
}

func (ps *SocketPacketStream) WriteUint16(n uint16) {
	ps.Data = append(ps.Data, []byte{byte(n), byte(n >> 8)}...)
}

func (ps *SocketPacketStream) WriteUint32(n uint32) {
	ps.Data = append(ps.Data, []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}...)
}

func (ps *SocketPacketStream) WriteUint64(n uint64) {
	ps.Data = append(ps.Data, []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}...)
}

func (ps *SocketPacketStream) WriteInt8(n int8) {
	ps.Data = append(ps.Data, byte(n))
}

func (ps *SocketPacketStream) WriteInt16(n int16) {
	ps.Data = append(ps.Data, []byte{byte(n), byte(n >> 8)}...)
}

func (ps *SocketPacketStream) WriteInt32(n int32) {
	ps.Data = append(ps.Data, []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}...)
}

func (ps *SocketPacketStream) WriteInt64(n int64) {
	ps.Data = append(ps.Data, []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}...)
}

func (ps *SocketPacketStream) WriteFloat32(n float32) {
	ps.Data = append(ps.Data, Float32ToByte(n)...)
}

func (ps *SocketPacketStream) WriteFloat64(n float64) {
	ps.Data = append(ps.Data, Float64ToByte(n)...)
}

func (ps *SocketPacketStream) WriteString(n string) {
	length := len(n)
	ps.WriteUint16(uint16(length))
	ps.Data = append(ps.Data, n...)
}

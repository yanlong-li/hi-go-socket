package stream

func (ps *PacketStream) WriteBool(data bool) {
	if data {
		ps.Data = append(ps.Data, byte(1))
	} else {
		ps.Data = append(ps.Data, byte(0))
	}
}

func (ps *PacketStream) WriteUint16(n uint16) {
	ps.Data = append(ps.Data, []byte{byte(n), byte(n >> 8)}...)
}

func (ps *PacketStream) WriteUint32(n uint32) {
	ps.Data = append(ps.Data, []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}...)
}

func (ps *PacketStream) WriteUint64(n uint64) {
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

func (ps *PacketStream) WriteString(n string) {
	length := len(n)
	ps.WriteUint16(uint16(length))
	ps.Data = append(ps.Data, n...)
}

package stream

import (
	"encoding/binary"
	"math"
)

func Uint16ToBytes(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}
func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

package packet

import (
	"math"
)

// []byte转uint64
func BytesToUint64(data []byte) uint64 {
	result := uint64(0)
	for k, v := range data {
		if k == 0 {
			result += uint64(v)
		} else {
			result += uint64(math.Pow(256, float64(k))) * uint64(v)
		}
	}
	return result
}

// []byte 转float32 ieee754
func HexToFloat32(data []byte) float32 {
	length := 4
	if len(data) != length {
		return 0
	}
	zf := 0
	for k, v := range data {
		ks := uint32(((length - 1) * 8) - k*8)
		zf += int(v) << ks
	}
	x := (zf & ((1 << (uint32((length-1)*8) - 1)) - 1)) + 1<<(uint32((length-1)*8)-1)
	exp := (zf >> (uint32((length-1)*8) - 1) & 255) - 127
	return float32(float64(x) * math.Pow(2, float64(exp-(((length-1)*8)-1))))
}

// []byte 转float64 ieee754
func HexToFloat64(data []byte) float64 {
	length := 8
	if len(data) != length {
		return 0
	}
	zf := 0
	for k, v := range data {
		ks := uint32(((length - 1) * 8) - k*8)
		zf += int(v) << ks
	}
	x := (zf & ((1 << (uint32((length-1)*8) - 1)) - 1)) + 1<<(uint32((length-1)*8)-1)
	exp := (zf >> (uint32((length-1)*8) - 1) & 255) - 127
	return float64(x) * math.Pow(2, float64(exp-(((length-1)*8)-1)))
}

package packet

import (
	"reflect"
)

// 注册包处理动作编码
var Packets = make(map[uint32]interface{})

// 反注册包处理动作编码
var RPackets = make(map[interface{}]uint32)

// 包服务注册应由引用源程序处理并注册到当前包的packets之下 支持多包注册，op自 ++
func Register(op uint32, packet interface{}, packets ...interface{}) {
	Packets[op] = packet
	RPackets[reflect.TypeOf(packet)] = op

	// 多包注册，op自 ++
	for _, v := range packets {
		op++
		Packets[op] = v
		RPackets[reflect.TypeOf(packet)] = op
	}
}

func OpCode(PacketModel interface{}) uint32 {

	elem := reflect.TypeOf(PacketModel)

	if v, ok := RPackets[elem]; ok {
		return v
	}
	return 0
}

func Packet(op uint32) interface{} {
	if v, ok := Packets[op]; ok {
		return v
	}
	return nil
}

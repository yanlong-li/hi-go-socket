package packet

import (
	"reflect"
)

// 注册包处理动作编码
var registerPackets = make(map[uint8]map[uint32]interface{})

// 反注册包处理动作编码
var registerReversalPackets = make(map[uint8]map[interface{}]uint32)

// 包服务注册应由引用源程序处理并注册到当前包的packets之下 支持多包注册，op自 ++
func Register(group uint8, op uint32, packet interface{}, packets ...interface{}) {

	groupPackets := registerPackets[group]
	if groupPackets == nil {
		groupPackets = make(map[uint32]interface{})
	}
	groupPackets[op] = packet
	registerPackets[group] = groupPackets

	groupRPackets := registerReversalPackets[group]
	if groupRPackets == nil {
		groupRPackets = make(map[interface{}]uint32)
	}
	groupRPackets[reflect.TypeOf(packet)] = op
	registerReversalPackets[group] = groupRPackets

	// 多包注册，op自 ++
	for _, v := range packets {
		op++
		registerPackets[group][op] = v
		registerReversalPackets[group][reflect.TypeOf(v)] = op
	}
}

func OpCode(group uint8, PacketModel interface{}) uint32 {

	elem := reflect.TypeOf(PacketModel)

	if v, ok := registerReversalPackets[group][elem]; ok {
		return v
	}
	return 0
}

func Packet(group uint8, op uint32) interface{} {
	if v, ok := registerPackets[group][op]; ok {
		return v
	}
	return nil
}

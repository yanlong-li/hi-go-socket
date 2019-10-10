package packet

import (
	"reflect"
)

// 注册包处理动作编码
var Packets = make(map[uint16]interface{})

// 反注册包处理动作编码
var RPackets = make(map[interface{}]uint16)

// 包服务注册应由引用源程序处理并注册到当前包的packets之下
func Register(op uint16, packet interface{}) {
	Packets[op] = packet
	RPackets[reflect.TypeOf(packet)] = op
}

func OpCode(packet interface{}) uint16 {

	elem := reflect.TypeOf(packet)

	if v, ok := RPackets[elem]; ok {
		return v
	}
	return 0
}

func Packet(op uint16) interface{} {
	if v, ok := Packets[op]; ok {
		return v
	}
	return nil
}

package route

import (
	"github.com/yanlong-li/hi-go-socket/packet"
	"reflect"
)

var routes = make(map[uint8]map[interface{}]interface{})

func Register(group uint8, packet, fun interface{}) {

	groupRoutes := routes[group]
	if groupRoutes == nil {
		groupRoutes = make(map[interface{}]interface{})
	}
	groupRoutes[reflect.TypeOf(packet)] = fun
	routes[group] = groupRoutes
}

func Handle(group uint8, op uint32) interface{} {

	p := packet.Packet(group, op)

	if p != nil {

		if v, ok := routes[group][reflect.TypeOf(p)]; ok {
			return v
		}
	}

	return nil

}

package route

import (
	"HelloWorld/io/network/packet"
	"reflect"
)

var routes = make(map[interface{}]interface{})

func Register(packet, fun interface{}) {
	routes[reflect.TypeOf(packet)] = fun
}

func Handle(op uint16) interface{} {

	p := packet.Packet(op)

	if p != nil {

		if v, ok := routes[reflect.TypeOf(p)]; ok {
			return v
		}
	}

	return nil

}

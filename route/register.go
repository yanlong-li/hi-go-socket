package route

import (
	"github.com/yanlong-li/hi-go-socket/packet"
	"reflect"
)

var routes = make(map[interface{}]interface{})

func Register(packet, fun interface{}) {
	routes[reflect.TypeOf(packet)] = fun
}

func Handle(op uint32) interface{} {

	p := packet.Packet(op)

	if p != nil {

		if v, ok := routes[reflect.TypeOf(p)]; ok {
			return v
		}
	}

	return nil

}

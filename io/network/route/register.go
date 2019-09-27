package route

import "HelloWorld/io/network/packet"

var routes = make(map[interface{}]interface{})

func Register(packet, fun interface{}) {
	routes[packet] = fun
}

func Handle(op uint16) interface{} {

	p := packet.Packet(op)

	if p != nil {

		if v, ok := routes[p]; ok {
			return v
		}
	}

	return nil

}

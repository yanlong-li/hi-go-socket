package packet

var packets map[uint16]interface{} = make(map[uint16]interface{})

func init() {
	register(0x0001, handleHelloWorld)
}

func register(op uint16, packet interface{}) {
	packets[op] = packet
}

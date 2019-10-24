package connect

import "net"

type Connector struct {
	Conn net.Conn
	ID   uint32
	Type uint8
}

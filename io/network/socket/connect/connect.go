package connect

import "net"

type Connector struct {
	Conn net.Conn
	ID   uint64
	Type uint8
}

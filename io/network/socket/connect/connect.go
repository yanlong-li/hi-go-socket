package connect

import "net"

type SocketConnector struct {
	Conn net.Conn
	ID   uint64
	Type uint8
}

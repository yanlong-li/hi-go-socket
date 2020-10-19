package connect

import (
	"github.com/yanlong-li/hi-go-socket/connect"
	"net"
)

// Socket 连接器
type SocketConnector struct {
	Conn net.Conn
	connect.BaseConnector
}

package connect

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"net"
)

// Socket 连接器
type SocketConnector struct {
	Conn net.Conn
	connect.BaseConnector
}

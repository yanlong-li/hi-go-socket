package socket

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/connect"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Client(address string) {
	conn, err := net.Dial(Tcp, address)
	defer CloseClient(conn)
	if err != nil {
		logger.Fatal("连接服务器失败", 0, err)
	}
	logger.Debug("已连接到服务器", 0)
	// 写入本地连接列表
	connector := connect.SocketConnector{Conn: conn}
	connector.Connected()
}

func CloseClient(conn net.Conn) {
	_ = conn.Close()
}

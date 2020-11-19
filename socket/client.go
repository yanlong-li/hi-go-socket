package socket

import (
	"github.com/yanlong-li/hi-go-logger"
	baseConnect "github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket/connect"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Client(group uint8, address string) {
	conn, err := net.Dial(Tcp, address)
	if err != nil {
		logger.Warning("连接服务器失败", 0, err)
		return
	}
	defer CloseClient(conn)
	logger.Debug("已连接到服务器", 0)
	// 写入本地连接列表
	connector := connect.SocketConnector{
		Conn: conn,
		BaseConnector: baseConnect.BaseConnector{
			ID:    baseConnect.GetAutoSequenceID(),
			Type:  baseConnect.TcpSocketClient,
			Group: group,
		},
	}
	connector.Connected()
}

func CloseClient(conn net.Conn) {
	_ = conn.Close()
}

package socket

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	baseConnect "github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/connect"
	"log"
	"net"
)

//开始服务
// 需要参数 监听地址:监听端口
func Server(address string) {

	service, err := net.Listen(Tcp, address)
	if err != nil {
		logger.Fatal("SOCKET服务开启失败", 0, err)
	}
	logger.Debug("SOCKET服务开启成功", 0, address)
	defer service.Close()

	for {
		//time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			// 写入本地连接列表
			socketConnect := &connect.SocketConnector{Conn: conn, BaseConnector: baseConnect.BaseConnector{ID: baseConnect.GetAutoSequenceID()}}
			go baseConnect.Add(socketConnect)
			go socketConnect.Connected()
		}

	}

}

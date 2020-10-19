package socket

import (
	"github.com/yanlong-li/hi-go-logger"
	baseConnect "github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket/connect"
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
	defer CloseService(service)

	for {
		//time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			// 写入本地连接列表
			socketConnect := &connect.SocketConnector{Conn: conn, BaseConnector: baseConnect.BaseConnector{ID: baseConnect.GetAutoSequenceID()}}
			go socketConnect.Connected()
		}

	}

}

func CloseService(service net.Listener) {
	_ = service.Close()
}

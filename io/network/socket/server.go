package socket

import (
	"fmt"
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
		log.Fatal(err)
	}
	fmt.Println("SOCKET服务开启成功", address)
	defer service.Close()
	var i uint64
	for {
		//time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			// 写入本地连接列表
			socketConnect := &connect.SocketConnector{Conn: conn, ID: i}
			baseConnect.Add(socketConnect)
			go socketConnect.Connected()
		}
		i++

	}

}

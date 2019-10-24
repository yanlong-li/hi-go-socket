package socket

import (
	baseConnect "HelloWorld/io/network/connect"
	"HelloWorld/io/network/socket/connect"
	"fmt"
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
	fmt.Println("服务开启成功", address)
	defer service.Close()
	var i uint32
	for {
		//time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			// 写入本地连接列表
			socketConnect := &connect.Connector{Conn: conn, ID: i}
			baseConnect.Add(socketConnect)
			go socketConnect.Connected()
		}
		i++

	}

}

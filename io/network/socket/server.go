package socket

import (
	"HelloWorld/io/network/socket/connect"
	"fmt"
	"log"
	"net"
)

//开始服务
// 需要参数 监听地址:监听端口
func Server() {

	service, err := net.Listen(Tcp, ":3001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务开启成功")
	defer service.Close()
	var i uint32
	for {
		//time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			// 写入本地连接列表
			connector := connect.Connector{Conn: conn, ID: i}
			connect.Save(connector)
			//connect.List[i] = connector
			go connector.Connected()
		}
		i++
		log.Printf("%d: accept a new connection\n", i)

	}

}

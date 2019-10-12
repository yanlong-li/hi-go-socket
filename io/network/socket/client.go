package socket

import (
	"HelloWorld/io/network/socket/connect"
	"fmt"
	"log"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Client() {
	conn, err := net.Dial(Tcp, "118.187.7.147:3001")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("已连接到服务器")
	// 写入本地连接列表
	connector := connect.Connector{Conn: conn}
	connect.List[99] = connector
	connector.Connected()
}

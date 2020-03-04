package socket

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/connect"
	"log"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Client(address string) {
	conn, err := net.Dial(Tcp, address)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("已连接到服务器")
	// 写入本地连接列表
	connector := connect.SocketConnector{Conn: conn}
	connector.Connected()
}

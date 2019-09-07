package client

import (
	"HelloWorld/io/network/base"
	"HelloWorld/io/network/packetStream"
	"fmt"
	"log"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Start() {
	conn, err := net.Dial(base.Tcp, ":8080")
	fmt.Println(conn, err)
	if err != nil {
		log.Fatal(err)
	}

	ps := packetStream.PacketStream{Stream: conn}
	for {
		var buf = make([]byte, 8192)
		_, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		ps.Len = uint16(packetStream.BytesToUint64(buf[0:2]))
		ps.Data = buf[2 : ps.Len+2]
		ps.Handle()
	}
}

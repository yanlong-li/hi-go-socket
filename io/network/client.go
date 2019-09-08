package network

import (
	"HelloWorld/io/network/stream"
	"fmt"
	"log"
	"net"
)

//连接服务
// 需要参数 监听地址:监听端口
func Client() {
	conn, err := net.Dial(Tcp, ":8080")
	fmt.Println(conn, err)
	if err != nil {
		log.Fatal(err)
	}

	ps := stream.PacketStream{Stream: conn}
	for {
		var buf = make([]byte, 8192)
		_, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		ps.Len = uint16(stream.BytesToUint64(buf[0:2]))
		ps.Data = buf[2 : ps.Len+2]
		ps.Handle()
	}
}

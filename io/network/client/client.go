package client

import (
	"HelloWorld/io/network/base"
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
	for {
		var buf = make([]byte, 8192)
		mType, err := conn.Read(buf)
		fmt.Println(mType, err, string(buf))
		_, _ = conn.Write([]byte(`A`))
		fmt.Println(0xF)
	}
}

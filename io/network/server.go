package network

import (
	"fmt"
	"log"
	"net"
	"time"
)

//开始服务
// 需要参数 监听地址:监听端口
func Server() {

	service, err := net.Listen(Tcp, ":8080")
	fmt.Println(service, err)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()
	var i int
	for {
		time.Sleep(time.Second * 10)
		if conn, err := service.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			go handleConnected(conn)
		}
		i++
		log.Printf("%d: accept a new connection\n", i)

	}

}

// 处理每个连接
func handleConnected(conn net.Conn) {

	qs := make([]byte, 0)
	// 字符长度
	qs = append(qs, byte(0x03))
	qs = append(qs, byte(0x00))
	// 操作码
	qs = append(qs, byte(0x01))
	qs = append(qs, byte(0x00))
	// 内容
	qs = append(qs, byte(0xFF))
	conn.Write(qs)
	//fmt.Println(conn.Write([]byte(`A`)))

}

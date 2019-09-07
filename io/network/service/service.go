package service

import (
	"HelloWorld/io/network/base"
	"fmt"
	"log"
	"net"
	"time"
)

//开始服务
// 需要参数 监听地址:监听端口
func Start() {

	service, err := net.Listen(base.Tcp, ":8080")
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
			handleMessage(conn)
		}
		i++
		log.Printf("%d: accept a new connection\n", i)

	}

}

func handleMessage(conn net.Conn) {

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

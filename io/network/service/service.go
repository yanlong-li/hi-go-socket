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

	fmt.Println(conn.Write([]byte(`A`)))

}

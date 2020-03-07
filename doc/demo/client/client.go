package main

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket"
)

// 定义数据包模型结构体，暂时不支持map类型

type user struct {
	Id   uint32
	Name string
	Age  bool
}
type message struct {
	Msg string
}

type connected struct {
}

func main() {

	//注册包 op uint32 0-5999为保留(或许实际只使用10个左右)，请从6000开始定义， 6000~4,294,967,295
	packet.Register(6000, user{})
	packet.Register(6001, message{})

	packet.Register(packet.CONNECTION, connected{})

	//注册动作 匿名注册
	route.Register(connected{}, func(connector connect.Connector) {
		fmt.Println("已连接服务器")
		fmt.Println("发送用户资料到服务器")
		connector.Send(user{Id: 1, Name: "客户端1号", Age: true})
	})
	//或 实名注册
	route.Register(message{}, recvMessage)

	socket.Client("127.0.0.1:3000")
}

// 收到消息
func recvMessage(m message, connector connect.Connector) {
	fmt.Println("收到服务器反馈的信息")
	fmt.Println(m.Msg)
}

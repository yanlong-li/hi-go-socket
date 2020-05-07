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

// 定义一个模型
type connected struct {
}

func main() {

	//注册数据包模型
	packet.Register(packet.Connection, connected{})
	//注册系统路由
	route.Register(connected{}, func(connector connect.Connector) {
		fmt.Println("一个新的客户端已连接")
	})
	//info 关于系统级动作，独立于通信之外的动作，系统动作传递的参数请参考系统动作表
	//info 普通动作固定两个参数 1：普通包 2 connect.Connector 连接器接口

	//注册普通包 op uint32 0-5999为保留(或许实际只使用10个左右)，请从6000开始定义， 6000~4,294,967,295
	packet.Register(6000, user{})
	packet.Register(6001, message{})

	//注册普通动作
	route.Register(user{}, func(u user, connector connect.Connector) {
		fmt.Println("收到客户端上传的用户资料")
		fmt.Println(u)
		_ = connector.Send(message{
			Msg: "已收到您的用户信息",
		})
	})

	// 该操作会阻塞执行，如果用 goroutine 请用其它操作进行阻塞保持程序运行
	socket.Server("127.0.0.1:3000")

}

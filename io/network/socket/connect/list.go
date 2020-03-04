package connect

import (
	"encoding/binary"
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/stream"
	"reflect"
)

// 处理每个连接
func (conn *SocketConnector) Connected() {

	//处理首次连接动作
	conn.beforeAction()
	// 处理连接断开后的动作
	defer conn.afterAction()

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("一个连接发生异常")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		_ = conn.Conn.Close()
		fmt.Println("断开连接")
	}()

	for {
		var buf = make([]byte, 8192)
		_, err := conn.Conn.Read(buf)
		if err != nil {
			//log.Fatal(err)
			fmt.Println("连接断开")
			break
		}
		// 每次动作不一致都注册一个单独的动作来处理
		ps := stream.PacketStream{}
		ps.Len = binary.LittleEndian.Uint16(buf[0:2])
		ps.OpCode = binary.LittleEndian.Uint32(buf[2:6])
		if !conn.RecvAction(ps.OpCode) {
			continue
		}
		if uint16(len(buf)) < ps.Len+2 {
			fmt.Println("数据不正确")
			continue
		}
		ps.Data = buf[6 : ps.Len+2]
		f := route.Handle(ps.OpCode)
		if f != nil {
			in := ps.Unmarshal(f)
			in[len(in)-1] = reflect.ValueOf(conn)
			go reflect.ValueOf(f).Call(in)
		} else {
			fmt.Println("未注册的包:", ps.OpCode)
		}

	}
}

// 建立连接时
func (conn *SocketConnector) beforeAction() {

	f := route.Handle(0)
	if f != nil {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(conn)
		reflect.ValueOf(f).Call(in)
	} else {
		fmt.Println("没有设置连接包:", 0)
	}
}

// 断开连接时
func (conn *SocketConnector) afterAction() {

	_ = conn.Conn.Close()

	connect.Del(conn.ID)

	f := route.Handle(1)
	if f != nil {
		//构造一个存放函数实参 Value 值的数纽
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(conn.ID)
		reflect.ValueOf(f).Call(in)
	} else {
		fmt.Println("没有设置连接包:", 1)
	}
}

// 收到数据包时
func (conn *SocketConnector) RecvAction(opCode uint32) bool {
	f := route.Handle(2)
	if f != nil {
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(opCode)
		in[1] = reflect.ValueOf(conn)
		result := reflect.ValueOf(f).Call(in)
		if len(result) >= 1 {
			return result[0].Bool()
		}
		return false
	} else {
		return true
	}
}

// 发送数据包
func (conn *SocketConnector) Send(model interface{}) {
	ps := stream.PacketStream{}
	ps.Marshal(model)
	//创建固定长度的数组节省内存
	data := make([]byte, 0, ps.GetLen()+4)
	data = append(data, connect.WriteUint16(ps.GetLen()+4)...)
	data = append(data, connect.Uint32ToHex(ps.OpCode)...)
	data = append(data, ps.Data...)

	_, err := conn.Conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}

//获取连接id
func (conn *SocketConnector) GetId() uint64 {
	return conn.ID
}

//广播数据包
// yourself 是否广播给自己
func (conn *SocketConnector) Broadcast(model interface{}, yourself bool) {
	connect.BroadcastChan <- connect.BroadcastModel{Model: model, Conn: conn, Self: yourself}
}

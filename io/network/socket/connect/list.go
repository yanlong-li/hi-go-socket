package connect

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/stream"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
)

// 处理每个连接
func (conn *Connector) Connected() {

	//处理首次连接动作
	conn.beforeAction()
	// 处理连接断开后的动作
	defer conn.afterAction()
	for {
		var buf = make([]byte, 8192)
		bufLen, err := conn.Conn.Read(buf)
		if err != nil {
			//log.Fatal(err)
			fmt.Println("连接断开")
			break
		}
		log.Print(buf[0:bufLen])
		// 每次动作不一致都注册一个单独的动作来处理
		ps := stream.PacketStream{}
		ps.Len = binary.LittleEndian.Uint16(buf[0:2])
		ps.OpCode = binary.LittleEndian.Uint32(buf[2:6])
		if uint16(len(buf)) < ps.Len+2 {
			fmt.Println("数据不正确")
			break
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
func (conn *Connector) beforeAction() {

	f := route.Handle(0)
	if f != nil {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(conn)
		reflect.ValueOf(f).Call(in)
	} else {
		fmt.Println("没有设置连接包:", 0)
	}
}

// 准备断开连接
func (conn *Connector) afterAction() {

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

// 发送数据包
func (conn *Connector) Send(model interface{}) {
	ps := stream.PacketStream{}
	ps.Marshal(model)
	//创建固定长度的数组节省内存
	data := make([]byte, 0, ps.GetLen()+2)
	data = append(data, connect.WriteUint16(ps.GetLen()+2)...)
	data = append(data, connect.Uint32ToHex(ps.OpCode)...)
	data = append(data, ps.Data...)

	_, err := conn.Conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}

//获取连接id
func (conn *Connector) GetId() uint64 {
	return conn.ID
}

//广播数据包
// yourself 是否广播给自己
func (conn *Connector) Broadcast(model interface{}, yourself bool) {
	connect.BroadcastChan <- connect.BroadcastModel{Model: model, Conn: conn, Self: yourself}
}

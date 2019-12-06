package connect

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/websocket/stream"
	"encoding/binary"
	"encoding/hex"
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
		// 读取消息
		_, message, err := conn.Conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			// 停止继续循环
			break
		}
		log.Printf("recv: %s", message)
		// uint16 = 4 uint32 = 8 uint64 = 16
		var OpCodeType uint8 = 8
		//监听动作
		if len(message) >= int(OpCodeType) {
			OpCode, err := hex.DecodeString(string(message[0:OpCodeType]))
			if err != nil {
				fmt.Println("获取动作错误")
			} else {
				actionOp := binary.LittleEndian.Uint32(OpCode)
				//p := packet.Packet(actionOp)
				data := message[OpCodeType:]

				f := route.Handle(actionOp)
				if f != nil {
					in := stream.Unmarshal(f, data)
					in[len(in)-1] = reflect.ValueOf(conn)
					reflect.ValueOf(f).Call(in)
				} else {
					fmt.Println("未注册的包:", actionOp)
				}
			}
		} else {
			fmt.Println("动作长度不足4位")
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

func (conn *Connector) Send(model interface{}) {
	pd, err := stream.Marshal(model)
	data := make([]byte, 0)
	data = append(data, connect.WriteUint16(uint16(len(pd)+2))...)
	op := packet.OpCode(model)
	data = append(data, connect.Uint32ToHex(op)...)
	data = append(data, pd...)

	err = conn.Conn.WriteMessage(2, data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}

func (conn *Connector) GetId() uint64 {
	return conn.ID
}

//广播数据包
// yourself 是否广播给自己
func (conn *Connector) Broadcast(model interface{}, yourself bool) {
	connect.BroadcastChan <- connect.BroadcastModel{Model: model, Conn: conn, Self: yourself}
}

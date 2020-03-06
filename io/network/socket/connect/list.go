package connect

import (
	"encoding/binary"
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/stream"
	"reflect"
)

// 处理每个连接
func (conn *SocketConnector) Connected() {

	//处理首次连接动作
	conn.ConnectedAction()
	// 处理连接断开后的动作
	defer conn.DisconnectAction()

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
		conn.HandleData(buf)

	}
}

// 处理数据包
// 将数据处理流程拆分成独立公开的方法，方便二次调用
func (conn *SocketConnector) HandleData(data []byte) {

	// 每次动作不一致都注册一个单独的动作来处理
	ps := stream.SocketPacketStream{}
	ps.Len = binary.LittleEndian.Uint16(data[0:2])
	ps.OpCode = binary.LittleEndian.Uint32(data[2:6])

	if uint16(len(data)) < ps.Len+2 {
		logger.Debug("数据不正确", 0)
		return
	}
	ps.Data = data[6 : ps.Len+2]

	if !conn.RecvAction(ps.OpCode, ps.Data) {
		return
	}

	f := route.Handle(ps.OpCode)
	if f != nil {
		in := ps.Unmarshal(f)
		in = append(in, reflect.ValueOf(conn))
		go reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("未注册的包", 0, ps.OpCode)
	}
}

// 建立连接时
func (conn *SocketConnector) ConnectedAction() {

	f := route.Handle(packet.CONNECTION)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(conn))
		reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("没有设置连接成功动作:", 1)
	}
}

// 断开连接时
func (conn *SocketConnector) DisconnectAction() {

	_ = conn.Conn.Close()

	go connect.Del(conn.ID)
	go connect.AddIdleSequenceId(conn.ID)

	f := route.Handle(packet.DISCONNECTION)
	if f != nil {
		//构造一个存放函数实参 Value 值的数纽
		var in []reflect.Value
		in = append(in, reflect.ValueOf(conn.ID))
		reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("没有设置断开连接动作:", 1)
	}
}

// 收到数据包时
func (conn *SocketConnector) RecvAction(opCode uint32, data []byte) bool {
	f := route.Handle(packet.BEFORE_RECVING)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(opCode))
		in = append(in, reflect.ValueOf(data))
		in = append(in, reflect.ValueOf(conn))

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
	ps := stream.SocketPacketStream{}
	ps.Marshal(model)
	// 封包
	data := ps.ToData()

	// 发送之前进行数据处理：加密、压缩
	f := route.Handle(packet.BEFORE_SENDING)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(ps.OpCode))
		in = append(in, reflect.ValueOf(data))

		result := reflect.ValueOf(f).Call(in)
		if len(result) >= 1 {
			data = result[0].Bytes()
		}
	}

	_, err := conn.Conn.Write(data)
	if err != nil {
		logger.Debug("发送数据失败", 0)
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

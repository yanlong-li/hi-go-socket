package connect

import (
	"encoding/binary"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/stream"
	baseStream "github.com/yanlong-li/HelloWorld-GO/io/network/stream"
	"reflect"
)

// 处理每个连接
func (conn *SocketConnector) Connected() {

	//处理首次连接动作
	conn.ConnectedAction()
	// 处理连接断开后的动作
	defer conn.DisconnectAction()

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			logger.Debug("一个连接发生异常", 0, err) // 这里的err其实就是panic传入的内容
		}
		conn.Disconnect()

	}()

	for {
		var buf = make([]byte, 8192)
		_, err := conn.Conn.Read(buf)
		if err != nil {
			logger.Debug("连接断开", 0)
			break
		}
		conn.HandleData(buf)

	}
}

func (conn *SocketConnector) Disconnect() {
	_ = conn.Conn.Close()
	logger.Debug("断开连接", 0)
}

// 处理数据包
func (conn *SocketConnector) HandleData(data []byte) {

	//2byte的uint16长度+4byte的uint64OpCode码
	if len(data) < 6 {
		logger.Debug("数据不正确", 0)
		return
	}

	// 每次动作不一致都注册一个单独的动作来处理
	ps := stream.SocketPacketStream{}
	ps.SetLen(binary.LittleEndian.Uint16(data[0:2]))
	ps.SetOpCode(binary.LittleEndian.Uint32(data[2:6]))

	if uint16(len(data)) < ps.GetLen()+2 {
		logger.Debug("数据流不完整", 0)
		return
	}
	ps.SetData(data[6 : ps.GetLen()+2])

	if !conn.RecvAction(&ps) {
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
	go connect.Add(conn)

	f := route.Handle(packet.CONNECTION)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(conn))
		reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("没有设置连接成功动作:", 1)
	}
}

// 准备断开连接
func (conn *SocketConnector) DisconnectAction() {

	f := route.Handle(packet.DISCONNECTION)
	if f != nil {
		//构造一个存放函数实参 Value 值的数纽
		var in []reflect.Value
		in = append(in, reflect.ValueOf(conn))
		reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("没有设置断开连接动作:", 1)
	}

	_ = conn.Conn.Close()

	go connect.Del(conn.ID)
}

// 收到数据包时
func (conn *SocketConnector) RecvAction(bs baseStream.Interface) bool {
	f := route.Handle(packet.BEFORE_RECVING)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(bs))
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
func (conn *SocketConnector) Send(model interface{}) error {

	ps := &stream.SocketPacketStream{}
	ps.Marshal(model)
	// 封包
	data := ps.ToData()

	// 发送之前进行数据处理：加密、压缩
	f := route.Handle(packet.BEFORE_SENDING)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(ps))

		result := reflect.ValueOf(f).Call(in)
		if len(result) >= 1 {
			data = result[0].Bytes()
		}
	}

	_, err := conn.Conn.Write(data)
	if err != nil {
		logger.Debug("发送数据失败", 0)
	}
	return err
}

//广播数据包
func (conn *SocketConnector) Broadcast(model interface{}, yourself bool) {
	go connect.Broadcast(connect.BroadcastModel{Model: model, Conn: conn, Self: yourself})
}

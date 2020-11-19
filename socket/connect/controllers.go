package connect

import (
	"encoding/binary"
	"fmt"
	"github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/packet"
	"github.com/yanlong-li/hi-go-socket/route"
	"github.com/yanlong-li/hi-go-socket/socket/stream"
	baseStream "github.com/yanlong-li/hi-go-socket/stream"
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
		// 读取包体长度
		lenBuf, err := conn.readLenBuf(uint16(packet.BufLenLen))
		if err != nil {
			logger.Debug("数据包长度读取失败", 0)
			break
		}
		bufLen := binary.LittleEndian.Uint16(lenBuf)
		// 容不下 OpCode 要你何用
		if bufLen < uint16(packet.OpCodeLen) {
			logger.Debug("数据长度标识不正确", 0)
			break
		}
		buf, err := conn.readLenBuf(bufLen)
		if err != nil {
			logger.Debug("数据包体读取失败", 0)
			break
		}
		conn.HandleData(buf)
	}
}

// 读取指定长度的数据 解决粘包和拆包问题
// 极端情况会导致所有的后续包无法正常读取
func (conn *SocketConnector) readLenBuf(bufLen uint16) ([]byte, error) {
	//读取真正的数据包
	buf := make([]byte, bufLen)
	// 实际读取长度 redBufLen
	readBufLen, err := conn.Conn.Read(buf)
	if err != nil {
		return buf, err
	}

	if uint16(readBufLen) < bufLen {
		newBuf, err := conn.readLenBuf(bufLen - uint16(readBufLen))
		// 将newBuf和buf进行合并
		if err != nil {
			return buf, nil
		}
		buf = append(buf[:readBufLen], newBuf...)
	}

	return buf, nil

}

func (conn *SocketConnector) Disconnect() {
	_ = conn.Conn.Close()
	logger.Debug("断开连接", 0)
}

// 处理单个数据包 包体不含长度标识
func (conn *SocketConnector) HandleData(data []byte) {

	// 每次动作不一致都注册一个单独的动作来处理
	ps := stream.SocketPacketStream{}
	ps.SetLen(uint16(len(data)))
	if ps.GetLen() < uint16(packet.OpCodeLen) {
		logger.Debug("数据长度标识不正确", 0)
		return
	}
	ps.SetOpCode(binary.LittleEndian.Uint32(data[:packet.OpCodeLen]))
	if ps.OpCode < packet.ReservedCode {
		logger.Debug("OP码范围不正确", 0)
		return
	}

	ps.SetData(data[packet.OpCodeLen:ps.GetLen()])

	if !conn.RecvAction(&ps) {
		return
	}

	f := route.Handle(conn.GetGroup(), ps.OpCode)
	if f != nil {
		in := ps.Unmarshal(f)
		in = append(in, reflect.ValueOf(conn))
		go reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug(fmt.Sprintf("未注册的包:%d", ps.OpCode), 0, ps.OpCode)
	}
}

// 建立连接时
func (conn *SocketConnector) ConnectedAction() {
	go connect.Add(conn)

	f := route.Handle(conn.GetGroup(), packet.Connection)
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

	f := route.Handle(conn.GetGroup(), packet.Disconnection)
	if f != nil {
		//构造一个存放函数实参 Value 值的数纽
		var in []reflect.Value
		in = append(in, reflect.ValueOf(conn))
		reflect.ValueOf(f).Call(in)
	} else {
		logger.Debug("没有设置断开连接动作:", 1)
	}

	_ = conn.Conn.Close()

	go connect.Del(conn.GetId())
}

// 收到数据包时
func (conn *SocketConnector) RecvAction(bs baseStream.Interface) bool {
	f := route.Handle(conn.GetGroup(), packet.BeforeRecv)
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
	ps.Marshal(conn.GetGroup(), model)
	// 封包
	data := ps.ToData()

	// 发送之前进行数据处理：加密、压缩
	f := route.Handle(conn.GetGroup(), packet.BeforeSending)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(ps))
		in = append(in, reflect.ValueOf(conn))

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

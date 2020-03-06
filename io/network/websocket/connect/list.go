package connect

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/websocket/stream"
	"log"
	"reflect"
)

// 处理每个连接
func (conn *WebSocketConnector) Connected() {

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
		// 读取消息
		_, buf, err := conn.Conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			// 停止继续循环
			break
		}
		log.Printf("recv: %s", buf)
		conn.HandleData(buf)
	}
}

// 处理数据包
// 将数据处理流程拆分成独立公开的方法，方便二次调用
func (conn *WebSocketConnector) HandleData(buf []byte) {
	// uint16 = 4 uint32 = 8 uint64 = 16
	var OpCodeType uint8 = 8
	//监听动作
	if len(buf) >= int(OpCodeType) {
		OpCode, err := hex.DecodeString(string(buf[0:OpCodeType]))
		if err != nil {
			logger.Debug("获取动作错误", 0)
		} else {
			opCode := binary.LittleEndian.Uint32(OpCode)
			data := buf[OpCodeType:]
			if !conn.RecvAction(opCode, data) {
				return
			}

			f := route.Handle(opCode)
			if f != nil {
				in := stream.Unmarshal(f, data)
				in[len(in)-1] = reflect.ValueOf(conn)
				reflect.ValueOf(f).Call(in)
			} else {
				logger.Debug("未注册的包", 0, opCode)
			}
		}
	} else {
		logger.Debug("动作长度不足4位", 0)
	}
}

// 建立连接时
func (conn *WebSocketConnector) ConnectedAction() {

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
func (conn *WebSocketConnector) DisconnectAction() {

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
func (conn *WebSocketConnector) RecvAction(opCode uint32, data []byte) bool {
	f := route.Handle(packet.BEFORE_RECVING)
	if f != nil {
		var in []reflect.Value
		in = append(in, reflect.ValueOf(opCode))
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

func (conn *WebSocketConnector) Send(PacketModel interface{}) {
	ps := stream.WebSocketPacketStream{}
	ps.Marshal(PacketModel)
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

	err := conn.Conn.WriteMessage(2, data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}

func (conn *WebSocketConnector) GetId() uint64 {
	return conn.ID
}

//广播数据包
// yourself 是否广播给自己
func (conn *WebSocketConnector) Broadcast(model interface{}, yourself bool) {
	connect.BroadcastChan <- connect.BroadcastModel{Model: model, Conn: conn, Self: yourself}
}

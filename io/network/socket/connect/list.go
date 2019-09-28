package connect

import (
	packet2 "HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/stream"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
)

var List = make(map[net.Conn]Connector, 1)

// 处理每个连接
func (conn *Connector) Connected() {

	//处理首次连接动作
	beforeAction(conn)
	// 处理连接断开后的动作
	defer afterAction(conn.ID)

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
		ps.Data = buf[2 : ps.Len+2]
		ps.OpCode = ps.ReadUInt16()
		f := route.Handle(ps.OpCode)
		if f != nil {
			in := ps.Unmarshal(f)
			in[len(in)-1] = reflect.ValueOf(conn)
			reflect.ValueOf(f).Call(in)
		} else {
			fmt.Println("未注册的包:", ps.OpCode)
		}

	}
}

// 建立连接时
func beforeAction(conn *Connector) {

	ps := stream.PacketStream{}
	ps.Len = uint16(2)
	ps.Data = []byte{0, 0}
	ps.OpCode = ps.ReadUInt16()
	f := route.Handle(ps.OpCode)
	if f != nil {
		in := ps.Unmarshal(f)
		in[len(in)-1] = reflect.ValueOf(conn)
		reflect.ValueOf(f).Call(in)
	} else {
		fmt.Println("没有设置连接包:", ps.OpCode)
	}
}

// 准备断开连接
func afterAction(ID uint32) {
	ps := stream.PacketStream{}
	ps.Len = uint16(2)
	ps.Data = []byte{0, 0}
	ps.OpCode = 0x06
	f := route.Handle(ps.OpCode)
	if f != nil {
		in := ps.Unmarshal(f)
		in[len(in)-1] = reflect.ValueOf(ID)
		reflect.ValueOf(f).Call(in)
	} else {
		fmt.Println("没有设置连接包:", ps.OpCode)
	}
}

func (conn *Connector) Send(model interface{}) {
	packetStream := stream.PacketStream{}
	packetStream.Marshal(model)
	data := make([]byte, 0)
	data = append(data, WriteUint16(uint16(len(packetStream.Data)+2))...)
	op := packet2.OpCode(model)
	data = append(data, WriteUint16(op)...)
	data = append(data, packetStream.Data...)

	_, err := conn.Conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}
func WriteUint16(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}

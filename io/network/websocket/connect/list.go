package connect

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/websocket/stream"
	"encoding/hex"
	"fmt"
	"gorilla/websocket"
	"log"
	"net/http"
	"reflect"
)

var List = make(map[*websocket.Conn]Connector, 1)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var i uint32 = 0

func Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	i++
	// 写入本地连接列表
	connector := Connector{Conn: conn, ID: i}
	List[conn] = connector
	connector.Connected()
	defer conn.Close()
}

// 处理每个连接
func (conn *Connector) Connected() {

	//处理首次连接动作
	beforeAction(conn)
	// 处理连接断开后的动作
	defer afterAction(conn.ID)
	for {
		// 读取消息
		_, message, err := conn.Conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			// 停止继续循环
			break
		}
		log.Printf("recv: %s", message)
		//监听动作
		if len(message) >= 4 {
			OpCode, err := hex.DecodeString(string(message[0:4]))
			if err != nil {
				fmt.Println("获取动作错误")
			} else {
				actionOp := uint16(OpCode[0])*256 + uint16(OpCode[1])
				//p := packet.Packet(actionOp)
				data := message[4:]

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
func beforeAction(conn *Connector) {

	//ps := stream.PacketStream{}
	//ps.Len = uint16(2)
	//ps.Data = []byte{0, 0}
	//ps.OpCode = ps.ReadUInt16()
	//f := route.Handle(ps.OpCode)
	//if f != nil {
	//	in := ps.Unmarshal(f)
	//	in[len(in)-1] = reflect.ValueOf(conn)
	//	reflect.ValueOf(f).Call(in)
	//} else {
	//	fmt.Println("没有设置连接包:", ps.OpCode)
	//}
}

// 准备断开连接
func afterAction(ID uint32) {
	//ps := stream.PacketStream{}
	//ps.Len = uint16(2)
	//ps.Data = []byte{0, 0}
	//ps.OpCode = 0x06
	//f := route.Handle(ps.OpCode)
	//if f != nil {
	//	in := ps.Unmarshal(f)
	//	in[len(in)-1] = reflect.ValueOf(ID)
	//	reflect.ValueOf(f).Call(in)
	//} else {
	//	fmt.Println("没有设置连接包:", ps.OpCode)
	//}
}

func (conn *Connector) Send(model interface{}) {
	pd, err := stream.Marshal(model)
	data := make([]byte, 0)
	data = append(data, WriteUint16(uint16(len(pd)+2))...)
	op := packet.OpCode(model)
	data = append(data, WriteUint16(op)...)
	data = append(data, pd...)

	err = conn.Conn.WriteMessage(1, data)
	if err != nil {
		fmt.Println("发送数据失败", err)
	}
}
func WriteUint16(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}

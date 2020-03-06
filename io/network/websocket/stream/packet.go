package stream

import "github.com/yanlong-li/HelloWorld-GO/io/network/connect"

type WebSocketPacketStream struct {
	Data   []byte //数据储存体
	Index  uint16 //当前指针
	Len    uint16 //数据长度--来自消息告知
	OpCode uint32 //操作码
}

func (wsps *WebSocketPacketStream) GetLen() uint16 {
	wsps.Len = uint16(len(wsps.Data))
	return wsps.Len
}

func (wsps *WebSocketPacketStream) ToData() []byte {
	//创建固定长度的数组节省内存
	var data []byte
	data = append(data, connect.WriteUint16(wsps.GetLen()+4)...)
	data = append(data, connect.Uint32ToHex(wsps.OpCode)...)
	data = append(data, wsps.Data...)
	return data
}

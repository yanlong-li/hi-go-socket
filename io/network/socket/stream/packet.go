package stream

import "github.com/yanlong-li/HelloWorld-GO/io/network/connect"

type SocketPacketStream struct {
	Data   []byte //数据储存体
	Index  uint16 //当前指针
	Len    uint16 //数据长度--来自消息告知
	OpCode uint32 //操作码
}

func (sps *SocketPacketStream) ToData() []byte {
	//创建固定长度的数组节省内存
	data := make([]byte, 0, sps.GetLen()+4)
	data = append(data, connect.WriteUint16(sps.GetLen()+4)...)
	data = append(data, connect.Uint32ToHex(sps.OpCode)...)
	data = append(data, sps.Data...)
	return data
}

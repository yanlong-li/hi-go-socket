package packet

type PacketStream struct {
	Data   []byte //数据储存体
	Index  uint16 //当前指针
	Len    uint16 //数据长度--来自消息告知
	OpCode uint16 //操作码
}

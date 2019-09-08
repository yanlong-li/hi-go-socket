package stream

import "fmt"

func (ps *PacketStream) Handle() {
	// 开始处理收到的数据

	// 读取操作编码
	opcode := ps.ReadUInt16()
	data := ps.ReadUInt8()
	fmt.Println(opcode, data)
}

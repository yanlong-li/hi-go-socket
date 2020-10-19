package stream

import (
	"encoding/json"
	"github.com/yanlong-li/hi-go-socket/packet"
)

//将包结构体反射写入字节流中
func (websocketPacketStream *WebSocketPacketStream) Marshal(PacketModel interface{}) {
	websocketPacketStream.SetData([]byte{})

	data, _ := json.Marshal(PacketModel)
	websocketPacketStream.SetData(data)
	websocketPacketStream.SetOpCode(packet.OpCode(PacketModel))
	websocketPacketStream.SetLen(uint16(len(data)))
}

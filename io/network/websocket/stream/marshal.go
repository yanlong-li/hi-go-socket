package stream

import (
	"encoding/json"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
)

//将包结构体反射写入字节流中
func (wsps *WebSocketPacketStream) Marshal(PacketModel interface{}) {
	wsps.Data, _ = json.Marshal(PacketModel)
	wsps.OpCode = packet.OpCode(PacketModel)
	wsps.GetLen()
}

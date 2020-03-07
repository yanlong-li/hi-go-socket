package connect

//广播通道
var BroadcastChan = make(chan BroadcastModel, 10)

//广播模型
type BroadcastModel struct {
	Conn  Connector
	Model interface{}
	Self  bool
}

func Broadcast(model BroadcastModel) {
	BroadcastChan <- model
}

package connect

//连接器接口
type Connector interface {
	//Connected()
	Send(interface{}) error
	//connectedAction()
	//disconnectAction()
	GetId() uint64
	Broadcast(interface{}, bool)
	HandleData([]byte)
	//recvAction(stream.BaseStream) bool
	Disconnect()
}

const (
	SOCKET = iota
	WEBSOCKET
)

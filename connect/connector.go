package connect

import "github.com/yanlong-li/hi-go-socket/stream"

// 公开连接器要实现的接口
type Connector interface {
	Send(interface{}) error
	GetId() uint64
	GetType() uint8
	GetGroup() uint8
	Broadcast(interface{}, bool)
	HandleData([]byte)
	Disconnect()
}

// 不公开连接器要实现的接口
type PrivateConnector interface {
	RecvAction(stream.Interface) bool
	Connected()
	ConnectedAction()
	DisconnectAction()
}

const (
//SOCKET = iota
//WEBSOCKET
)

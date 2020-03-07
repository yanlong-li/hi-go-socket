package connect

import "github.com/yanlong-li/HelloWorld-GO/io/network/stream"

// 公开连接器要实现的接口
type Connector interface {
	Send(interface{}) error
	GetId() uint64
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
	SOCKET = iota
	WEBSOCKET
)

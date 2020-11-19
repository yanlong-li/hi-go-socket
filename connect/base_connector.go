package connect

const (
	TcpSocketServer uint8 = iota
	TcpSocketClient
	UdpSocketServer
	UdpSocketClient
	WebSocketServer
	WebSocketClient
	WEBRTC
)

// 连接器基础数据
type BaseConnector struct {
	ID    uint64
	Type  uint8
	Group uint8
}

//获取连接器ID
func (conn *BaseConnector) GetId() uint64 {
	return conn.ID
}

//获取连接器类型
func (conn *BaseConnector) GetType() uint8 {
	return conn.Type
}

//获取连接器类型
func (conn *BaseConnector) GetGroup() uint8 {
	return conn.Group
}

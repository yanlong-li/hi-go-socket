package connect

// 连接器基础数据
type BaseConnector struct {
	ID   uint64
	Type uint8
}

//获取连接器ID
func (conn *BaseConnector) GetId() uint64 {
	return conn.ID
}

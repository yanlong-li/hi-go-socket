package connect

//广播模型
type BroadcastModel struct {
	Conn  Connector
	Model interface{}
	Self  bool
}

// 广播
func Broadcast(model BroadcastModel) {
	RWConnectListLock.RLock()
	defer RWConnectListLock.RUnlock()
	for id, v := range connectList {

		//不含自己 则不发送给自己
		if model.Self == false && id == model.Conn.GetId() {
			continue
		}
		// 非同组，不可广播
		if model.Conn.GetGroup() != v.GetGroup() {
			continue
		}

		go broadcastSend(v, model)
	}
}

// 广播单播
func broadcastSend(v Connector, model BroadcastModel) {
	_ = v.Send(model.Model)
}

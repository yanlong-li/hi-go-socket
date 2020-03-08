package connect

var connectList = make(map[uint64]Connector)

// 添加连接通道
var AddChan = make(chan Connector, 10)

//删除连接通道
var DelChan = make(chan uint64, 10)

func init() {

	go func() {
		for {
			select {
			case conn := <-AddChan:
				connectList[conn.GetId()] = conn
			case ID := <-DelChan:
				delete(connectList, ID)
				go AddIdleSequenceId(ID)
			case BM := <-BroadcastChan:

				for id, v := range connectList {

					//不含自己 则不发送给自己
					if BM.Self == false && id == BM.Conn.GetId() {
						continue
					}
					go v.Send(BM.Model)
				}
			}
		}
	}()
}

func Add(conn Connector) {
	AddChan <- conn
}
func Del(ID uint64) {
	DelChan <- ID
}

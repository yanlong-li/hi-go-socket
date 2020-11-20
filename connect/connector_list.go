package connect

import "sync"

var connectList = make(map[uint64]Connector)

var RWConnectListLock sync.RWMutex

func Add(conn Connector) {
	RWConnectListLock.Lock()
	defer RWConnectListLock.Unlock()
	connectList[conn.GetId()] = conn
}
func Del(ID uint64) {
	RWConnectListLock.Lock()
	defer RWConnectListLock.Unlock()
	delete(connectList, ID)
	go AddIdleSequenceId(ID)
}

func Count() uint32 {
	return uint32(len(connectList))
}

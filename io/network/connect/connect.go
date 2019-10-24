package connect

var List = make(map[uint32]Connector, 1)

var SaveChan = make(chan Connector)
var DelChan = make(chan uint32)
var BroadcastChan = make(chan interface{})

type Connector interface {
	Send(interface{})
	Connected()
	//AfterAction()
	//BeforeAction()
	GetId() uint32
}

const (
	SOCKET = iota
	WEBSOCKET
)

func init() {

	go func() {
		for {
			select {
			case conn := <-SaveChan:
				List[conn.GetId()] = conn
			case ID := <-DelChan:
				delete(List, ID)
			case model := <-BroadcastChan:
				for _, v := range List {
					v.Send(model)
				}
			}
		}
	}()
}

func Add(conn Connector) {
	SaveChan <- conn
}
func Del(ID uint32) {
	DelChan <- ID
}

func WriteUint16(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}

func Broadcast(model interface{}) {
	BroadcastChan <- model
}

package connect

var List = make(map[uint64]Connector, 1)

//
var AddChan = make(chan Connector)
var DelChan = make(chan uint64)
var BroadcastChan = make(chan BroadcastModel)

type BroadcastModel struct {
	Conn  Connector
	Model interface{}
	Self  bool
}

type Connector interface {
	Send(interface{})
	Connected()
	//AfterAction()
	//BeforeAction()
	GetId() uint64
	Broadcast(interface{}, bool)
}

const (
	SOCKET = iota
	WEBSOCKET
)

func init() {

	go func() {
		for {
			select {
			case conn := <-AddChan:
				List[conn.GetId()] = conn
			case ID := <-DelChan:
				delete(List, ID)
			case BM := <-BroadcastChan:

				for id, v := range List {

					//不含自己 则不发送给自己
					if BM.Self == false && id == BM.Conn.GetId() {
						continue
					}
					v.Send(BM.Model)
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

func WriteUint16(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}

func Uint32ToHex(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

func Broadcast(model interface{}) {
	BroadcastChan <- BroadcastModel{Model: model, Self: true}
}

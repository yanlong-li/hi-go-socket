package connect

var connectList = make(map[uint64]Connector)

// 添加连接通道
var AddChan = make(chan Connector, 10)

//删除连接通道
var DelChan = make(chan uint64, 10)

//广播通道
var BroadcastChan = make(chan BroadcastModel, 10)

type BroadcastModel struct {
	Conn  Connector
	Model interface{}
	Self  bool
}

type Connector interface {
	Connected()
	Send(interface{})
	ConnectedAction()
	DisconnectAction()
	GetId() uint64
	Broadcast(interface{}, bool)
	HandleData([]byte)
	RecvAction(uint32, []byte) bool
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
				connectList[conn.GetId()] = conn
			case ID := <-DelChan:
				delete(connectList, ID)
			case BM := <-BroadcastChan:

				for id, v := range connectList {

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

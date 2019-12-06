package connect

import (
	"gorilla/websocket"
)

type Connector struct {
	Conn *websocket.Conn
	ID   uint64
	Type uint8
}

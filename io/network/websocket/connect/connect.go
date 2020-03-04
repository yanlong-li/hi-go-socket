package connect

import (
	"github.com/gorilla/websocket"
)

type WebSocketConnector struct {
	Conn *websocket.Conn
	ID   uint64
	Type uint8
}

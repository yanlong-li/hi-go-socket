package connect

import (
	"github.com/gorilla/websocket"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
)

// WebSocket 连接器
type WebSocketConnector struct {
	Conn *websocket.Conn
	connect.BaseConnector
}

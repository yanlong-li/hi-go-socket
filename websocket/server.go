package websocket

import (
	gorillaWebsocket "github.com/gorilla/websocket"
	"github.com/yanlong-li/hi-go-logger"
	baseConnect "github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/websocket/connect"
	"log"
	"net/http"
	"time"
)

var server *http.Server
var upGrader = gorillaWebsocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func Server(address string, group uint8) {
	logger.Debug("WS服务开启成功", 0, address)
	mux := http.NewServeMux()
	server = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}
	mux.HandleFunc("/", Connect)
	logger.Fatal("WS服务遇到错误", 0, server.ListenAndServe())
}

func Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// 写入本地连接列表
	connector := &connect.WebSocketConnector{
		Conn: conn,
		BaseConnector: baseConnect.BaseConnector{
			ID:    baseConnect.GetAutoSequenceID(),
			Type:  baseConnect.WebSocketServer,
			Group: 1,
		},
	}

	go connector.Connected()
}

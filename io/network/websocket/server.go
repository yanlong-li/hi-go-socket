package websocket

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	baseConnect "github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/websocket/connect"
	"log"
	"net/http"
	"time"
)

var server *http.Server
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func Server(address string) {

	var addr = flag.String("addr", address, "http service address")

	flag.Parse()
	log.SetFlags(0)
	fmt.Println("WS服务开启成功", address)
	mux := http.NewServeMux()
	server = &http.Server{
		Addr:         *addr,
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}
	mux.HandleFunc("/", Connect)
	log.Fatal(server.ListenAndServe())
}

func Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// 写入本地连接列表
	connector := &connect.WebSocketConnector{Conn: conn, ID: baseConnect.GetAutoSequenceID()}
	go baseConnect.Add(connector)
	go connector.Connected()
}

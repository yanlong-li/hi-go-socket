package websocket

import (
	"HelloWorld/io/network/websocket/connect"
	"flag"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8082", "http service address")

var server *http.Server

func Server() {
	flag.Parse()
	log.SetFlags(0)

	mux := http.NewServeMux()
	server = &http.Server{
		Addr:         *addr,
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}
	mux.HandleFunc("/ws", connect.Connect)
	log.Fatal(server.ListenAndServe())
}

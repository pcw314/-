package message

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

const (
	AppName = "message"
)

type Session struct {
	ConvID      string
	Connections map[uint]*websocket.Conn
	// 其他元数据
}

var Clients2 = make(map[uint]*Session)

// websocket.go
var Clients = make(map[uint]*websocket.Conn)
var ConvClients = make(map[uint]*websocket.Conn)
var Mutex sync.Mutex
var Upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

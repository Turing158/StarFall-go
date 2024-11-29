package util

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var WebSocketUpGrader websocket.Upgrader
var WsMap sync.Map

func WebSocketInit() {
	WebSocketUpGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func SendMessageFunc(user, message string) {
	if conn, ok := WsMap.Load(user); ok {
		conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(message))
	}
}

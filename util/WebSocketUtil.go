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

func WsSendMessageWithString(user, message string) {
	WsSendMessage(user, []byte(message))
}

func WsSendMessage(user string, message []byte) {
	if conn, ok := WsMap.Load(user); ok {
		conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, message)
	}
}

package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)
	mu      sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	conn.SetReadLimit(64 * 1024)
	_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	})

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	// Read loop to detect disconnects.
	go func() {
		defer func() {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			conn.Close()
		}()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()
}

func Broadcast(message []byte) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			delete(clients, client)
			client.Close()
		}
	}
}

package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WsMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type WsBroadcaster struct {
	clients    map[*websocket.Conn]bool
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast  chan []byte
	mu         sync.Mutex
}

func NewWsBroadcaster() *WsBroadcaster {
	return &WsBroadcaster{
		clients:    make(map[*websocket.Conn]bool),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		broadcast:  make(chan []byte),
	}
}

func (b *WsBroadcaster) Run() {
	log.Println("INFO [WsBroadcaster]: Starting broadcaster loop...")
	for {
		select {
		case conn := <-b.register:
			b.mu.Lock()
			b.clients[conn] = true
			b.mu.Unlock()
			log.Printf("INFO [WsBroadcaster]: Client connected: %s", conn.RemoteAddr())

		case conn := <-b.unregister:
			b.mu.Lock()
			if _, ok := b.clients[conn]; ok {
				delete(b.clients, conn)
				conn.Close()
				log.Printf("INFO [WsBroadcaster]: Client disconnected: %s", conn.RemoteAddr())
			}
			b.mu.Unlock()

		case message := <-b.broadcast:
			b.mu.Lock()
			for conn := range b.clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Printf("WARN [WsBroadcaster]: Failed to write message to client %s: %v. Unregistering.", conn.RemoteAddr(), err)
					go func(c *websocket.Conn) {
						b.unregister <- c
					}(conn)
				}
			}
			b.mu.Unlock()
		}
	}
}

func (b *WsBroadcaster) Broadcast(msgType string, payload interface{}) {
	message := WsMessage{Type: msgType, Payload: payload}
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Printf("ERROR [WsBroadcaster]: Failed to marshal broadcast message: %v", err)
		return
	}
	b.broadcast <- jsonMsg
}

func (b *WsBroadcaster) RegisterClient(conn *websocket.Conn) {
	b.register <- conn
}

func (b *WsBroadcaster) ClientLoop(conn *websocket.Conn) {
	defer func() {
		b.unregister <- conn
		conn.Close()
	}()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WARN [WsBroadcaster]: Client read error: %v", err)
			} else {
				log.Printf("INFO [WsBroadcaster]: Client closed connection gracefully: %s", conn.RemoteAddr())
			}
			break
		}
	}
}

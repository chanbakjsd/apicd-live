package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingInterval = time.Minute
	pongWait     = pingInterval + 5*time.Second
)

// Conn is a connection to the server.
type Conn struct {
	ws     *websocket.Conn
	closed bool
	mu     *sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// handleWebsocket handles a new websocket connection.
func (s *Server) handleWebsocket(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "This endpoint is GET only.", http.StatusMethodNotAllowed)
		return
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %s\n", err)
		return
	}
	conn := &Conn{
		ws: ws,
		mu: new(sync.Mutex),
	}
	s.Register(conn)
	go conn.readPump(s)
	go func() {
		for range time.NewTicker(pingInterval).C {
			conn.mu.Lock()
			if conn.closed {
				return
			}
			conn.ws.WriteMessage(websocket.PingMessage, []byte{})
			conn.mu.Unlock()
		}
	}()
	data, err := json.Marshal(s.config)
	if err != nil {
		log.Fatalln("Error marshalling config:", err)
	}
	msg := "config " + string(data)
	go conn.Write(msg)
}

// Close closes the underlying connections and channels.
func (conn *Conn) Close() {
	conn.closed = true
	conn.ws.Close()
}

// Write writes the message to the connection.
func (conn *Conn) Write(msg string) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.closed {
		// If the connection is closed, just drop the message.
		return
	}
	err := conn.ws.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
}

// readPump copies the commands provided by the client to the game.
func (conn *Conn) readPump(s *Server) {
	defer conn.Close()
	conn.ws.SetReadDeadline(time.Now().Add(pongWait))
	conn.ws.SetPongHandler(func(string) error {
		conn.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := conn.ws.ReadMessage()
		switch {
		case err == nil:
			// Process the message below.
		case websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway):
			// Unexpected close.
			log.Printf("Unexpected close: %s\n", err)
			return
		default:
			return
		}
		s.ProcessCommand(strings.Fields(string(message)))
	}
}

// pushState pushes the state to the specified connection.
func (s *Server) pushState(c *Conn) {
	data, err := json.Marshal(s.state)
	if err != nil {
		log.Fatalln("Error marshalling game state:", err)
	}
	msg := "state " + string(data)
	go c.Write(msg)
}

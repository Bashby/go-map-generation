package websocket

// Note: Adapted Heavily from https://github.com/gorilla/websocket/blob/master/examples/chat

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Constants
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Package Globals
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client Message handler structure around peers and their websocket connection
type Client struct {
	// Reference to parent Hub
	hub *Hub

	// Peer's websocket connection
	conn *websocket.Conn

	// Outbound message channel to peer
	outbound chan []byte
}

// handleWebsocketRequest Negotiates initial websocket request from the peer.
func handleWebsocketRequest(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, outbound: make(chan []byte, 256)}
	client.hub.register <- client

	// Start client message handler routines
	go client.outboundHandler()
	go client.inboundHandler()
}

// outboundHandler pumps messages from the parent hub to the peer websocket connection.
func (c *Client) outboundHandler() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.outbound:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				// TODO: Error here, log it
				return
			}
		}
	}
}

// inboundHandler Pumps messages from the peer websocket connection to the parent hub.
func (c *Client) inboundHandler() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// c.conn.SetReadLimit(maxMessageSize) // TODO: Clean this up using gorilla best practice?
	// c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// for {
	// 	_, message, err := c.conn.ReadMessage()
	// 	if err != nil {
	// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
	// 			log.Printf("error: %v", err)
	// 		}
	// 		break
	// 	}
	// 	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
	// 	c.hub.inbound <- message
	// }
}

package websocket

import "bitbucket.org/ashbyb/go-map-generation/protobuf"

type Hub struct {
	// Client set
	clients map[*Client]bool

	// Inbound messages channel from Clients
	inbound chan []byte

	// Register requests channel from new Clients
	register chan *Client

	// Unregister requests channel from existing Clients
	unregister chan *Client
}

// newHub Construct a websocket Hub to manage clients and messages to and from them
func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		inbound:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, exists := h.clients[client]; exists {
				delete(h.clients, client)
				close(client.outbound)
			}
		case message := <-h.inbound:
			go protobuf.Decode(message)
		}
	}
}

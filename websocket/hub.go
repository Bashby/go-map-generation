package websocket

type Hub struct {
	// Client set
	clients map[*Client]struct{}

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
		clients:    make(map[*Client]struct{}),
		inbound:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {

}

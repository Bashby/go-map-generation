package websocket

import (
	"log"

	"bitbucket.org/ashbyb/go-map-generation/protobuf"

	"github.com/golang/protobuf/proto"
)

type Hub struct {
	// Client set
	clients map[*Client]bool

	// Inbound messages channel from Clients
	inbound chan ClientMessage

	// Outbound messages channel to Clients
	outbound chan ClientMessage

	// Register requests channel from new Clients
	register chan *Client

	// Unregister requests channel from existing Clients
	unregister chan *Client
}

// newHub Construct a websocket Hub to manage clients and messages to and from them
func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		inbound:    make(chan ClientMessage),
		outbound:   make(chan ClientMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("[Client Birth][%s] Registered", client.conn.RemoteAddr().String())
			h.clients[client] = true
		case client := <-h.unregister:
			if _, exists := h.clients[client]; exists {
				log.Printf("[Client Death][%s] Unregistered\n", client.conn.RemoteAddr().String())
				delete(h.clients, client)
				close(client.outbound)
			}
		case message := <-h.inbound:
			go ProcessMessage(message, h)
		case message := <-h.outbound:
			if _, exists := h.clients[message.client]; exists {
				select {
				case message.client.outbound <- message.message:
				default:
					log.Printf("[Client Death][%s] Buffer full", message.client.conn.RemoteAddr().String())
					delete(h.clients, message.client)
					close(message.client.outbound)
				}
			}
		}
	}
}

func ProcessMessage(message ClientMessage, hub *Hub) {
	wrapper := &protobuf.Message{}

	// Decode message
	err := proto.Unmarshal(message.message, wrapper)
	if err != nil {
		log.Fatal("Unmarshaling: ", err)
	}

	// Process payload
	switch msg := wrapper.Payload.(type) {
	case *protobuf.Message_Move:
		handleMove(msg)
	case *protobuf.Message_Attack:
		handleAttack(msg)
	}
}

func handleMove(msg *protobuf.Message_Move) {
	log.Println("Twas a Move message: ", msg.Move.Direction)
}

func handleAttack(msg *protobuf.Message_Attack) {
	log.Println("Twas a Attack message: ", msg.Attack.Target)
}

// test := &websocket.Message{
// 	Type: 1,
// 	Payload: &websocket.Message_Move{
// 		Move: &websocket.Move{
// 			Direction: "Left",
// 		},
// 	},
// }
// data, err := proto.Marshal(test)
// if err != nil {
// 	log.Fatal("marshaling error: ", err)
// }

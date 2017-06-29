package protobuf

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func Decode(data []byte) {
	wrapper := &Message{}

	// Decode message
	err := proto.Unmarshal(data, wrapper)
	if err != nil {
		log.Fatal("Unmarshaling: ", err)
	}

	// Process payload
	switch msg := wrapper.Payload.(type) {
	case *Message_Move:
		handleMove(msg)
	case *Message_Attack:
		handleAttack(msg)
	}
}

func handleMove(msg *Message_Move) {
	fmt.Println("Twas a Move message: ", msg.Move.Direction)
}

func handleAttack(msg *Message_Attack) {
	fmt.Println("Twas a Attack message: ", msg.Attack.Target)
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

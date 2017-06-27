package main

import (
	"log"
	"math/rand"
	"time"

	"fmt"

	"bitbucket.org/ashbyb/go-map-generation/websocket"
	"github.com/golang/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// width, height := 500, 500
	// world.Generate(width, height)

	test := &websocket.Message{
		Type: 1,
		Payload: &websocket.Message_Move{
			Move: &websocket.Move{
				Direction: "Left",
			},
		},
	}
	test2 := &websocket.Message{
		Type: 1,
		Payload: &websocket.Message_Attack{
			Attack: &websocket.Attack{
				Target: "John Doe",
			},
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	data2, err2 := proto.Marshal(test2)
	if err2 != nil {
		log.Fatal("marshaling error: ", err2)
	}
	fmt.Println("Raw: ", data)
	fmt.Println("Raw2: ", data2)
	newTest := &websocket.Message{}
	newTest2 := &websocket.Message{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	err2 = proto.Unmarshal(data2, newTest2)
	if err2 != nil {
		log.Fatal("unmarshaling error: ", err2)
	}
	// Use a type switch to determine which oneof was set.
	switch u := test.Payload.(type) {
	case *websocket.Message_Move:
		fmt.Println("Twas a Move message: ", u.Move.Direction)
	case *websocket.Message_Attack:
		fmt.Println("Twas a Attack message: ", u.Attack.Target)
	}
	switch u := test2.Payload.(type) {
	case *websocket.Message_Move:
		fmt.Println("Twas a Move message: ", u.Move.Direction)
	case *websocket.Message_Attack:
		fmt.Println("Twas a Attack message: ", u.Attack.Target)
	}
}

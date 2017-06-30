package main

import (
	"math/rand"
	"time"

	"log"

	"flag"

	"bitbucket.org/ashbyb/go-map-generation/websocket"
	"bitbucket.org/ashbyb/go-map-generation/world"
)

var width int
var height int
var address string

func init() {
	// Define input parameters
	flag.IntVar(&width, "width", 1024, "width of the game world")
	flag.IntVar(&height, "height", 1024, "height of the game world")
	flag.StringVar(&address, "address", ":8080", "http service address to listen on")

	// Seed pRNG
	rand.Seed(time.Now().UTC().UnixNano())

	// Configure logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	// Parse inputs
	flag.Parse()

	// Generate the world
	world.Generate(width, height)

	// Serve
	websocket.Serve(address)
}

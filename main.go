package main

import (
	"math/rand"
	"time"

	"bitbucket.org/ashbyb/go-map-generation/world"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	width, height := 500, 500
	world.Generate(width, height)
}

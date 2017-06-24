package main

import (
	"bitbucket.org/ashbyb/go-map-generation/world"
)

func main() {
	width, height := 500, 500
	world.Generate(width, height)
}

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

	// minn := 0.0
	// maxx := 0.0

	// noiseGenerator := noise.NewFBMNoiseGenerator2D(16, .007, 2.0, 1.13, 0.5, 0.57, 0.0, 1.0)

	// for i := 0; i < 100000; i++ {
	// 	// dat := float64(i)
	// 	point := noiseGenerator.Get2D(rand.Float64()*10000, rand.Float64()*10000)
	// 	maxx = math.Max(maxx, point)
	// 	minn = math.Min(minn, point)
	// }
	// fmt.Println(maxx, minn)
}

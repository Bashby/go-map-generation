package world

import (
	"math/rand"

	"bitbucket.org/ashbyb/go-map-generation/noise"
)

// BuildImageNoise Builds a map using simplex noise
func BuildImageNoise(w int, h int) [][]float64 {
	r := noise.Generate()

	data := make([][]float64, w)
	for x := 0; x < w; x++ {
		data[x] = make([]float64, h)
		for y := 0; y < h; y++ {
			data[x][y] = r.Eval2(rand.Float64(), rand.Float64())
		}
	}

	return data
}

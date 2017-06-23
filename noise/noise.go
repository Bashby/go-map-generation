package noise

import (
	"math/rand"

	"github.com/ojrac/opensimplex-go"
)

// Generate This does something
func Generate() *opensimplex.Noise {
	// rand.Seed(time.Now().Unix())
	noise := opensimplex.NewWithSeed(rand.Int63())

	return noise
}

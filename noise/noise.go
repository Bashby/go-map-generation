package noise

import (
	"math/rand"
	"time"

	"github.com/ojrac/opensimplex-go"
)

type FBMOctave struct {
	frequency float64
	amplitude float64
}

type FBMNoiseGenerator struct {
	generator      *opensimplex.Noise
	seed           int64
	octaves        []*FBMOctave
	baseFrequency  float64
	baseAmplitude  float64
	lacunarity     float64
	gain           float64
	totalAmplitude float64
}

func NewFBMOctave(frequency float64, amplitude float64) *FBMOctave {
	return &FBMOctave{frequency, amplitude}
}

func NewFBMNoiseGenerator(octaveCount int, baseFrequency float64, baseAmplitude float64, lacunarity float64, gain float64) *FBMNoiseGenerator {
	// Parse inputs
	f := baseFrequency
	a := baseAmplitude
	l := lacunarity
	g := gain
	o := make([]*FBMOctave, octaveCount)

	// Seed the noise generator
	seed := time.Now().Unix()
	rand.Seed(seed)

	// Build octaves
	totalAmplitude := 0.0
	for i := range o {
		o[i] = NewFBMOctave(f, a)
		totalAmplitude += a
		f *= l
		a *= g
	}
	return &FBMNoiseGenerator{opensimplex.NewWithSeed(rand.Int63()), seed, o, baseFrequency, baseAmplitude, l, g, totalAmplitude}
}

// ReSeed Seed the generator with a new random seed
func (n *FBMNoiseGenerator) ReSeed() {
	newSeed := rand.Int63()
	n.generator = opensimplex.NewWithSeed(newSeed)
	n.seed = newSeed
}

// GenerateFBMNoise Generate a noise value using FBM
func (n *FBMNoiseGenerator) GenerateFBMNoise(x int, y int) float64 {
	noise := 0.0
	for _, octave := range n.octaves {
		noise += n.generator.Eval2(float64(x)*octave.frequency, float64(y)*octave.frequency) * octave.amplitude
	}

	// Average noise values across octaves
	noise /= n.totalAmplitude

	return noise
}

// BuildNoiseMatrix Builds a 2d matrix of given size filled with simplex noise and Fractal Brownian Motion
func (n *FBMNoiseGenerator) BuildNoiseMatrix(w int, h int, min int, max int) [][]float64 {
	data := make([][]float64, w)
	for x := 0; x < w; x++ {
		data[x] = make([]float64, h)
		for y := 0; y < h; y++ {
			noise := n.GenerateFBMNoise(x, y)

			// fmt.Println(noise)

			// Normalize noise [-1.0, 1.0] to requested range [min, max]
			noise = float64((max-min))*((noise+0.7)/1.4) + float64(min)

			//noise = noise*float64(max-min)/2 + float64(max+min)/2

			data[x][y] = noise
		}
	}

	return data
}

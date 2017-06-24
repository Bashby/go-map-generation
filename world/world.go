package world

import (
	"image"
	"image/png"
	"log"
	"os"

	"bitbucket.org/ashbyb/go-map-generation/noise"

	"image/color"
)

func Generate(w int, h int) {
	// Generate Noise
	noiseGenerator := noise.NewFBMNoiseGenerator(16, .007, 1.0, 2.0, 0.5)
	elevationNoise := noiseGenerator.BuildNoiseMatrix(w, h, 0, 255)
	noiseGenerator.ReSeed()
	moistureNoise := noiseGenerator.BuildNoiseMatrix(w, h, 0, 255)

	// Create an image from target size
	img := image.NewRGBA(image.Rect(0, 0, w*2, h*2))

	// Load in biome data
	reader, err := os.Open("world/biomes.v2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	biomeData, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	// biomeDataDebug := biomeData

	// Render out the biome debug square
	debugOffset := 50
	for x := 0; x < 256; x++ {
		// Horizontal Lines
		img.Set(w+x+debugOffset, 0+debugOffset, color.RGBA{255, 0, 0, 255})
		img.Set(w+x+debugOffset, 255+debugOffset, color.RGBA{255, 0, 0, 255})
	}
	for y := 0; y < 256; y++ {
		// Vertical Lines
		img.Set(w+debugOffset, y+debugOffset, color.RGBA{255, 0, 0, 255})
		img.Set(w+255+debugOffset, y+debugOffset, color.RGBA{255, 0, 0, 255})
	}
	// biomeDataDebug.Bounds

	// Generate output
	for i := range elevationNoise {
		for j := range elevationNoise[i] {
			elevation := RoundToInt(elevationNoise[i][j])
			moisture := RoundToInt(moistureNoise[i][j])
			biome := DetermineBiome(biomeData, elevation, moisture)
			img.Set(i, j, biome)
			img.Set(i, j+h, color.RGBA{uint8(elevation), uint8(elevation), uint8(elevation), 255})
			img.Set(i+w, j+h, color.RGBA{uint8(moisture), uint8(moisture), uint8(moisture), 255})
			// fmt.Println(elevation, moisture)
			// Debug Biome
			img.Set(w+moisture+debugOffset, 255-elevation+debugOffset, color.White)
		}
	}

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func DetermineBiome(biomes image.Image, elevation int, moisture int) color.Color {
	return biomes.At(moisture, elevation)
}

func min(x, y int) uint8 {
	if x < y {
		return uint8(x)
	}
	return uint8(y)
}

// RoundToInt8 rounds 64-bit floats into integer numbers
func RoundToInt(a float64) int {
	if a < 0 {
		return int(a - 0.5)
	}
	return int(a + 0.5)
}

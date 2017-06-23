package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"bitbucket.org/ashbyb/go-map-generation/world"
)

func main() {
	rand.Seed(time.Now().Unix())
	width, height := 500, 500
	data := world.BuildImageNoise(width, height)
	data2 := world.BuildImageNoise(width, height)
	data3 := world.BuildImageNoise(width, height)

	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := range data {
		for j := range data[i] {
			valr := data[i][j]
			valg := data2[i][j]
			valb := data3[i][j]
			// fmt.Println(i, j, val, min(255, uint8(val*256)))
			r := min(255, int(valr*256))
			g := min(255, int(valg*256))
			b := min(255, int(valb*256))
			// fmt.Println(r, g, b)
			img.Set(i, j, color.RGBA{r, g, b, 255})
		}
	}

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func min(x, y int) uint8 {
	if x < y {
		return uint8(x)
	}
	return uint8(y)
}

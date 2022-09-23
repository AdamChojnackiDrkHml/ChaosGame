package main

import (
	trianglechaos "ChaosGame/internal/triangleChaos"
	"image"
	"image/png"
	"os"
)

func main() {
	iterations := 100000
	size := 1000
	stPt := [2]int{400, 350}

	bitmap := trianglechaos.CreateChaosTriangle(size, stPt, iterations)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewGray(image.Rectangle{upLeft, lowRight})

	img.Pix = bitmap

	f, _ := os.Create("data/chTri.png")
	png.Encode(f, img)

}

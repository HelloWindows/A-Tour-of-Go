package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is struct
type Image struct {
	width, hight int
}

// ColorModel id implementated
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds id implementated
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.hight)
}

// At id implementated
func (img Image) At(x, y int) color.Color {
	imgFunc := func(x, y int) uint8 {
		return uint8((x * y) / 2)
	}
	v := imgFunc(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}

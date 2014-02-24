package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	height int
	width  int
}

func (this Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (this Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, this.height, this.width)
}

func (this Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}

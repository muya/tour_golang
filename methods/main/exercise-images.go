package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(d, e int) color.Color {
	t := uint8((d * d * d) * (e * e * e))
	return color.RGBA{t, t, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}

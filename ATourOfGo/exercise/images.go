package exercise

import (
	"image"
	"image/color"
)

type MyImage struct {
}

func (mi MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (mi MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (mi MyImage) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}
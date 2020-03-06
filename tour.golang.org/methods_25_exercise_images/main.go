#tour.golang.org/methods/25
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}
func (i Image) At(x, y int) color.Color {
	return color.RGBA{255, 0, uint8(x^y+y^x), 255}
}
func main() {
	m := Image{}
	pic.ShowImage(m)
}

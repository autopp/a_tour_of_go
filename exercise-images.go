package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, 10, 20)
}

func (img Image) At(x, y int) color.Color {
  return color.RGBA{20, 20, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

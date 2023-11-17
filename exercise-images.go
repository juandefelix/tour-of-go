package main

import (
  "golang.org/x/tour/pic"
  "image/color"
  "image"
)

type Image struct{
  width int
  height int
}

func (Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
  a := image.Point{0, 0}
  b := image.Point{i.width, i.height}
  return image.Rectangle{a, b}
}

func (i Image) At(width, height int) color.Color {
  r := uint8(100 + width)
  g := uint8(100 + height)
  b := uint8(128)
  a := uint8(255)

  return color.RGBA{r, g, b, a}
}


func main() {
  m := Image{100, 100}
  pic.ShowImage(m)
}

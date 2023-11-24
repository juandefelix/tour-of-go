package main

import (
  "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
  a := make([][]uint8, dy)

  for y := 0; y < dy; y++ {
    a[uint8(y)] = make([]uint8, dx)
    for x := 0; x < dx; x++ {
      a[uint8(y)][uint8(x)] = uint8(x^y)
    } 
  }
  return a
}

func main() {
  pic.Show(Pic)

}

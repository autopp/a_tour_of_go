package main

import "golang.org/x/tour/pic"

// returns matrix
func Pic(dx, dy int) [][]uint8 {
  matrix := make([][]uint8, dy)
  for i := 0; i < dy; i++ {
    matrix[i] = make([]uint8, dx)
  }

  return matrix
}

func main() {
	pic.Show(Pic)
}

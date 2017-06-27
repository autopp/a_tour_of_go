package main

import (
	"fmt"
  "math"
)

// Calc square root by Newton's method
func Sqrt(x float64) float64 {
  delta := 0.01
  z := 1.0
  z_next := z + delta * 2
  for n := 0; math.Abs(z - z_next) > delta; n++ {
    z_next = z - (z * z - x) / 2 * z
    fmt.Printf("%d: %g\n", n, z_next)
    z, z_next = z_next, z
  }

  return z
}

func main() {
	fmt.Println(Sqrt(2))
}

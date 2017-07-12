package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
    return x, ErrNegativeSqrt(x)
  }

  delta := 0.01
  z := 1.0
  z_next := z + delta * 2
  for n := 0; math.Abs(z - z_next) > delta; n++ {
    z_next = z - (z * z - x) / 2 * z
    z, z_next = z_next, z
  }

  return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  v := fmt.Sprint(float64(e))
  return fmt.Sprintf("cannot Sqrt negative number: %v", v)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

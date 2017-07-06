package main

import (
	"fmt"
	"math"
)

// `compute` receive a function which receive two float and return one float
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
  // function is 1st class
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

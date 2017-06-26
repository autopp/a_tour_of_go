package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
  // You can write a statement before a condition of `if`
  // A variable declared at there is only available in `if` block
	if v := math.Pow(x, n); v < lim {
		return v
	}
  // fmt.Println(v) <- `v` is not available
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

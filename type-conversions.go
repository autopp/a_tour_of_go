package main

import (
	"fmt"
	"math"
)

func main() {
  // `T(v)` cast value `v` to type `T`
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)

  // Unlike C, Go requires explicit type conversion
  // This causes compile error
  // var z uint = f
}

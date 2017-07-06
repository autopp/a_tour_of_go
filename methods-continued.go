package main

import (
	"fmt"
	"math"
)

// Method declaration is not limited on `struct`

// Method declaration is allowed for type in same package
// builtin-type (`int`, `float64`...) is also not allowed

// If you need method for builtin or other package type, use `type`
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

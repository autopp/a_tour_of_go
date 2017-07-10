package main

import (
	"fmt"
	"math"
)

// require a method `M`
type I interface {
	M()
}

type T struct {
	S string
}

// `T` implements `I`
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// `F` implements `I`
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
  // A value of interface can be regarded as a tuple of value and implemented type
	fmt.Printf("(%v, %T)\n", i, i)
}

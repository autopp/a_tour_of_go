package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// receiver type should be unified either value or pointer

// Pointer methods melit 1
// poiner can modify field
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


// Pointer methods melit 2:
// pointer avoid copy of value
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// function version of `Abs` at methods-pointers.go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function version of `Scale` at methods-pointers.go
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

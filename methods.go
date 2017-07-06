package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go dosen't support class system
// But, method is supported
// `Abs` receives a `Vertex` as receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

  // In function, value parameter cannot accept pointer
  // But in method, value receiver can accept pointer,
  // It is converted to value automatically
	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // equals `(*p).Abs()`
	fmt.Println(AbsFunc(*p))
  // fmt.Println(AbsFunc(p)) <- compile error
}

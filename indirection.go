package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
  // In function, pointer parameter cannot accept value
  // But in method, pointer receiver can accept value,
  // It is converted to pointer automatically
	v := Vertex{3, 4}
	v.Scale(2) // equals `(&v).Scale(2)`
	ScaleFunc(&v, 10)
  // ScaleFunc(v, 10) <- compile error

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

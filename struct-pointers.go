package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
  // `p` is the pointer of `Vertex`
	p := &v
  // pointer also can access to field by `.`
  // This is syntax sugarof `(*p).X` (In C, `p->x`)
	p.X = 1e9
	fmt.Println(v)
}

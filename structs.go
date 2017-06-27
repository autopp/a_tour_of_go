package main

import "fmt"

// `struct` is set of fields
// (`type` is used for declaration new type)
type Vertex struct {
	X int
	Y int
}

func main() {
  // `struct` is initilized by `StructName{a, b, c...}`
	fmt.Println(Vertex{1, 2})
}

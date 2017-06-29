package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // struct's field is accessed by `.`
	fmt.Println(v.X)
}

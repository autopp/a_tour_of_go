package main

import "fmt"

func main() {
  // A type of variable declared without notion is The type of right value
	v := 42 // int
  // v := 3.142 // float64
  // v := 0.5i // complex128
  // x := v // just type of v
	fmt.Printf("v is of type %T\n", v)
}

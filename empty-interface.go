package main

import "fmt"

func main() {
  // The type of `I` is a interface which requires zero method
  // So, `i` accepts any value of type
	var i interface{}
  // Empty interface is used when handle unknown type
  // E.g. `fmt.Printf` accepts empty interface

	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

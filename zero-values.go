package main

import "fmt"

func main() {
  // With no initializers, variable is initilized with `zero value` of eath type
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}

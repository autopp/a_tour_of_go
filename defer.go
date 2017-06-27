package main

import "fmt"

func main() {
  // A function call with `defer` is called at end of callee (In this case, end of `main`)
  // Note that the arguments are evaluated immediately (In this case, `"world"`)
	defer fmt.Println("world")

	fmt.Println("hello")
}

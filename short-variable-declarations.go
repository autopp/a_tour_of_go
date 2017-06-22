package main

import "fmt"

func main() {
	var i, j int = 1, 2

  // In function, if you give initializer can omit `var`
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// This is not regal
// ./short-variable-declarations.go:17: syntax error: non-declaration statement outside function body
// x := 42

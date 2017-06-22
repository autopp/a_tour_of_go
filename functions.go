package main

import "fmt"

// A type of paramters/return value is writen after variable name
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

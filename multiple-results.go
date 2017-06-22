package main

import "fmt"

// A function can return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
  // A caller can recieve multiple return values
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

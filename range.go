package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  // array, slice and map is iterated by `range` keyword
  // At each iteration, `range` returns index and value (copied)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

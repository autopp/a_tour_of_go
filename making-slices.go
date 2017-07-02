package main

import "fmt"

func main() {
  // `make` create a array and return slice pointing to it
  // 2nd argument is length
	a := make([]int, 5)
	printSlice("a", a)

  // 3rd argument is capacity
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

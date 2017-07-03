package main

import "fmt"

func main() {
  // `append` add new element to slice.
  // If array's capacity is not enough to store,
  // `append` create a new array and return pointer to it.
  // See also: https://golang.org/pkg/builtin/#append

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

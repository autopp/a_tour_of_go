package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

  // A slice has no contents
  // A slice likes a reference of array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

  // A modification slice refects original array and other slices
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

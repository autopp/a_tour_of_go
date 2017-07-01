package main

import "fmt"

func main() {
  // zero value of slice is `nil`
	var s []int
  // `nil` slice has 0 length and capacity
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

package main

import "fmt"

func main() {
  // A length of array is part of type
	var a [2]string // A array of string (length == 2)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // A array of int (length == 6)
	fmt.Println(primes)
}

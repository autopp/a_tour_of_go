package main

import "fmt"

func main() {
	var i interface{} = "hello"

  // `i.(string)` check whether `i` is a `string`, and get its value
	s := i.(string)
	fmt.Println(s) // "hello"

  // second value of assertion is boolean whether check is passed
	s, ok := i.(string)
	fmt.Println(s, ok) // "hello true"

  // When assertion is failed, first value is zero value, and second value is `false`
	f, ok := i.(float64)
	fmt.Println(f, ok) // "0 false"

  // When second value is not receive and assertion is failed, its become runtime error
	f = i.(float64) // panic
	fmt.Println(f)
}

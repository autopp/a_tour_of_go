package main

import "fmt"

func do(i interface{}) {
  // type `switch` is similar to regular switch
  // In each case, type of `v` is a matched type
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
  do(&Point{3, 4})
	do(true)
}

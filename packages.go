package main

import (
	"fmt" // Program can use package as `fmt`
	"math/rand" // Program can use package as `rand`
)

// all files of fmt should start with `package fmt`.
// all files of math/rand should start with `package rand`.

func main() {
  // rand.Intn(n) returns a integer in [0, n)
	fmt.Println("My favorite number is", rand.Intn(10))
}

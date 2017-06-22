package main

import (
	"fmt"
	"math"
)

func main() {
  // exported name from package should start with upper case
  fmt.Println(math.Pi)

  // This line shoud be compile error
  // fmt.Println(math.pi)
}

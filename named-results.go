package main

import "fmt"

// x and y are return values of split
// NOTE: named return values shoud be only used in small function for readablity
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

  // When named return value is declared, return statement can ommit value
	return
}

func main() {
	fmt.Println(split(17))
}

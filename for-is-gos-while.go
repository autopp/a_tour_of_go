package main

import "fmt"

func main() {
	sum := 1
  // semicolon are optional also
  // In Go, `for` is used as `while`
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
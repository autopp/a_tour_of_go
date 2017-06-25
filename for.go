package main

import "fmt"

func main() {
	sum := 0

  // In Go, parentheses are not needed
  // Braces are needed
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

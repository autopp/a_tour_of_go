package main

import "fmt"

func main() {
	sum := 1
  // initialization and post process are optional
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	pow := make([]int, 10)

  // `, value` is optional
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

  // By using `_`, you can discard index/value
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

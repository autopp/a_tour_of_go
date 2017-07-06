package main

import "fmt"

// `adder` returns function which receive and return one int
func adder() func(int) int {
	sum := 0

  // this func bind to `sum`
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

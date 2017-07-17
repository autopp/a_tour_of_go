package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

  // If you want to tell no more data to receiver, you can close a channel with `close`.
  // If you don't need, you shoud not to close channel
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

  // `for i := range c` repeat until channel is closed
	for i := range c {
		fmt.Println(i)
	}

  // You can know whether channel is closed by second value:
  // x, b <- c
  // NOTE: receiver should not close channel
}

package main

import "fmt"

func main() {
  // 2nd paramter is length of channel buffer
  // When buffer is max, sender is blocked
  // When buffer is empty, receiver is blocked
  ch := make(chan int, 2)
	ch := make(chan int, 2) // -> fatal error: all goroutines are asleep - deadlock!
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

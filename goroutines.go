package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
  // The function and arguments (`say`, and `"world"`) are evaluated at the caller goroutine (main)
  // `say("world")` is executed at the new goroutine
	go say("world")
	say("hello")
}

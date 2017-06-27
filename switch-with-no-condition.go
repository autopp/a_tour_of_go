package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

  // `switch` with no condition is equals to `switch true`.
  // This is convenient to replace very long `if`-`then`-`else`.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

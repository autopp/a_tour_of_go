package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
    // When `defer` is used multiple, there are stacked
    // At end of callee, defered calling are called by LIFO
    // More infomation: https://blog.golang.org/defer-panic-and-recover
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

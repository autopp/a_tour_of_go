package main

import (
	"fmt"
	"time"
)

// In Go, a error is represented by builtin `error` interface:
//
// type error interface {
//   Error() string
// }
//
// A functions useally return `error` as second value
// The caller handle error by checking second value (If not `nil`, error is occured!)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

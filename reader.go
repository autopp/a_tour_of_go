package main

import (
	"fmt"
	"io"
	"strings"
)

// `io.Reader` is interface for reading from a stream
//
// type Reader interface {
//   func Read(b []byte) (n int, err error)
// }
//
// `Read` stores read value to `b` and returns a size as first value
// When a stream is ending, it returns `io.EOF` as second value

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

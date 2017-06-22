package main

// There are basic types in Go:
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias of uint8
//
// rune // alias of int32 (It's represent a codepoint of Unicode)
//
// float32 float64
//
// complex64 complex128

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  // %T format print type
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

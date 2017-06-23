package main

import "fmt"

const Pi = 3.14

func main() {
  // A constant is declared with `const` prefix
  // cannot use `:=`
	const World = "world"
  // const World = "is not enough" <- compile error
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

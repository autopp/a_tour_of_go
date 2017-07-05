package main

import "fmt"

func main() {
	m := make(map[string]int)

  // Set value
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

  // Set value
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

  // Delete value
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

  // Second value is boolean for key existence
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

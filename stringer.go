package main

import "fmt"

// `Stringer` interface is defined at 'fmt'
// it is defined as
//
// type Stringer interface {
//   String() string
// }
//
// Many functions (E.g. `fmt.Println`) check a value whether implements `Stringer` or not
// for converting to string

type Person struct {
	Name string
	Age  int
}

// `Person` implements `Stringer`
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
  // `fmt.Println` use `Person.String()` to output
	fmt.Println(a, z)
}

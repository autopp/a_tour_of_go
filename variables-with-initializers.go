package main

import "fmt"

// var with initializers
var i, j int = 1, 2

// If omitted, the variable is initilized with default value (that is defiend for each type)

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

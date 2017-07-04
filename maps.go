package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// `map` relate key and value
// zero value of map is a `nil`
// `nil` map cannot store any pair
var m map[string]Vertex

func main() {
  // `make` returns initilized map
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m1 map[string]Vertex1

func main() {
	m1 = make(map[string]Vertex1)
	m1["Bell Labs"] = Vertex1{
		40.3232323, -88.2323,
	}
	fmt.Println(m1["Bell Labs"])
}

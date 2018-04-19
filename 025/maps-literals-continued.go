package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m = map[string]Vertex2{
	"Bell Labs": {34.111111, -45.22222},
	"Goolgle ":  {37.7777771, -46.22222},
}

func main() {
	fmt.Println(m)
}

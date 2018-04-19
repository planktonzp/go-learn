package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex3{
	"Bell Labs": Vertex3{
		43.1233, -55.1255,
	},
	"Google": Vertex3{
		37.4343, -88.23333,
	},
}

func main() {
	fmt.Println(m2)
}

package main

import (
	"fmt"
	"math"
)

type P2Vertex struct {
	X, Y float64
}

func (v P2Vertex) P2Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func P2AbsFunc(v P2Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := P2Vertex{3, 4}
	fmt.Println(v.P2Abs())
	fmt.Println(P2AbsFunc(v))

	p := P2Vertex{4, 3}
	fmt.Println(p.P2Abs())
	fmt.Println(P2AbsFunc(p))
}

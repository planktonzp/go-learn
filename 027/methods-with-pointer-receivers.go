package main

import (
	"fmt"
	"math"
)

type P4Vertex struct {
	X, Y float64
}

func (v *P4Vertex) Scale4(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *P4Vertex) Abs4() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &P4Vertex{3, 4}
	fmt.Printf("Before scaling: %+v ,Abs: %v, address: %p \n", v, v.Abs4(), &v)
	v.Scale4(5)
	fmt.Printf("After sclaing: %+v ,Abs: %v, address: %p \n", v, v.Abs4(), &v)
}

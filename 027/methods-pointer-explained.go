package main

import (
	"fmt"
)

type PVertex struct {
	X, Y float64
}

func (v *PVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *PVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := PVertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &PVertex{6, 8}
	p.Scale(5)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

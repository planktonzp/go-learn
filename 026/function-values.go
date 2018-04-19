package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	//compute作用就是返回一个含有俩双精度数的结构体fn？
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

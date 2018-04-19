package main

import (
	"fmt"
	"math"
)

type MVertex struct {
	X, Y float64
}

func (v MVertex) Abs() float64 {	//方法就是带有接收者参数的函数，Abs函数的接收者是v 类型是MVertex(结构体 内含两个浮点型变量)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := MVertex{3, 4}
	fmt.Println(v.Abs())
}

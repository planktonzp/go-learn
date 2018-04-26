package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

//	a = v	//Vertex does not implement Abser (Abs method has pointer receiver)

//  Abs方法使用值接收者时是为什么可以用传地址的方式传值？ go内置隐式转换？ 指针默认会转换成对应地址的值？
	fmt.Println(a.Abs())
}

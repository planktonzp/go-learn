package main

import "fmt"

func adder() func(int) int {	//函数adder的返回值是一个闭包 可以调用外部定义的sum变量 并将计算结果存进sum中
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()	//初始化两个变量并把adder函数(作为值)赋给他们，类似传参一样
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

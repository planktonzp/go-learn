package main

import (
	"fmt"
)

func fibonacci() func() []int { //声明一个切片返回类型，作为闭包的返回类型
	var s = []int{0, 1}
	return func() []int { //声明一个切片返回类型，作为闭包操作的外部变量的返回值，要和上层一致
		s = append(s, s[len(s)-1]+s[len(s)-2])
		return s
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

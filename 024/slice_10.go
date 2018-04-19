package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //range 返回两个值 第一个是元素下标第二个为元素的具体值
		fmt.Printf("2**%d=%d \n", i, v)
	}
}

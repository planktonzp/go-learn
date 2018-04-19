package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //左移位，指数增长
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value) //遍历数组时可以用下划线忽略
	}
}

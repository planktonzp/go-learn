package main

import "fmt"

func main() {
	a := make([]int, 5) //含义是设计一个长度为5,元素都是零的数组然后创建一个长度和数组一样的切片指向数组,即容量等于长度
	printSlice("a", a)

	b := make([]int, 0, 5) //含义是设计一个长度为5,元素都是零的数组然后创建一个长度为0的切片指向数组,即切片容量是0
	printSlice("b", b)

	c := b[:2] //切片截取的元素是切片指向的数组的内存地址
	printSlice("c", c)

	d := c[2:5] //但是这个地址是相对前一个数组或者切片而言的相对地址,即下标
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

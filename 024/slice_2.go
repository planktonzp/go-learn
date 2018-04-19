package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Pual",
		"George",
		"Ringo", //竖着写的时候结尾不是大括号就得加逗号
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	c := b[0:2]
	fmt.Println(a, b, c)

	b[0] = "xxx"
	fmt.Println(a, b, c)
	fmt.Println(names)
}

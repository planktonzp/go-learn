package main

import "fmt"

func main() {
	m3 := make(map[string]int)

	m3["Answer"] = 43
	fmt.Println(m3)

	m3["Answer"] = 666
	fmt.Println(m3)

	delete(m3, "Answer")
	fmt.Println(m3)

	//双赋值探测映射中健是否还存在
	v, ok := m3["Answer"]
	fmt.Println("The Values", v, "present", ok)
}

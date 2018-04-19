package main

import "fmt"

func main() {
	s := []int{}
	//var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("s is nil")
	}
}

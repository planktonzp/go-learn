package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice2(s)

	s = s[:0]
	printSlice2(s)

	s = s[:4]
	printSlice2(s)

	s = s[2:]
	printSlice2(s)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}

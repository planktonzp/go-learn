package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += 5
	}
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"math"
)

func Pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		Pow2(3, 2, 10),
		Pow2(3, 3, 10),
	)
}

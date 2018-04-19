//十次循環用牛頓法求x的平方根 z=(z^2-x)/(2*z)
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		if z*z-x != 0 {
			z -= (z*z - x) / (2 * z)
			fmt.Printf("z = %g\n", z)
		}
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(3),
	)
}

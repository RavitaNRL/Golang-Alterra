package main

import (
	"fmt"
)

func SimpleEquation(a, b, c int) {
	// perulangan untuk mencari nilai x, y, dan z
	for x := 1; x <= a/3; x++ {
		y := (a - x) / 2
		z := a - x - y
		if x*y*z == b && x*x+y*y+z*z == c {
			fmt.Println(x, y, z)
			return // keluar dari fungsi
		}
	}
	fmt.Println("No solution.")
}

func main() {
	SimpleEquation(1, 2, 3)    // No solution.
	SimpleEquation(6, 6, 14)   // 1 1 3
	SimpleEquation(12, 10, 30) // No solution.
}

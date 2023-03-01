package main

import (
	"fmt"
)

func getPrime(number int) int {
	if number < 1 {
		return 0
	}

	count := 0
	for z := 2; z <= 50; z++ {
		Prime := true
		for j := 2; j*j <= z; j++ {
			if z%j == 0 {
				Prime = false
				break
			}
		}
		if Prime {
			count++
			if count == number {
				return z
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(getPrime(1))  // 2
	fmt.Println(getPrime(5))  // 11
	fmt.Println(getPrime(8))  // 19
	fmt.Println(getPrime(9))  // 23
	fmt.Println(getPrime(10)) // 29
}

package main

import "fmt"

func fibona(bil int) int {
	if bil <= 1 {
		return bil
	}
	return fibona(bil-1) + fibona(bil-2)
}

func main() {
	fmt.Println(fibona(0))  // 0
	fmt.Println(fibona(1))  // 1
	fmt.Println(fibona(2))  // 1
	fmt.Println(fibona(3))  // 2
	fmt.Println(fibona(5))  // 5
	fmt.Println(fibona(6))  // 8
	fmt.Println(fibona(7))  // 13
	fmt.Println(fibona(9))  // 34
	fmt.Println(fibona(10)) // 55
}

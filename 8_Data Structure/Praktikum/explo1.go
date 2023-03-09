package main

import "fmt"

func main() {
	fmt.Println(diagonal([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	})) // 2
	fmt.Println(diagonal([][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	})) // 15
}

func diagonal(arr [][]int) int {
	z := len(arr)
	var left, right int
	for i := 0; i < z; i++ {
		left += arr[i][i]
		right += arr[i][z-i-1]
	}
	return abs(left - right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

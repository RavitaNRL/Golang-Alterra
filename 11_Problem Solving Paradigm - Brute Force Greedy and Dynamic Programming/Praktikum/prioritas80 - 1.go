package main

import "fmt"

func digitBiner(bilBulat int) []int {
	ans := make([]int, bilBulat+1) // make digunakan untuk menginisialisasikan slice dengan ukuran bilBulat+1.
	for i := 0; i <= bilBulat; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}

func main() {
	fmt.Println(digitBiner(2)) // [0 1 1]
	fmt.Println(digitBiner(8)) // [0 1 1 2 1 2 2 3 1]
}

package main

import (
	"fmt"
	"math"
)

func frogs(sum []int, N int) int {
	if N == 1 {
		return 0
	}
	if N == 2 {
		// untuk mengembalikan nilai selisi dari 2 selisih pada slice.
		//math.Abs digunakan untuk mengambil nilai absolut dari selisih tersebut.
		return int(math.Abs(float64(sum[0] - sum[1])))
	}
	return int(math.Min(float64(frogs(sum, N-1)+int(math.Abs(float64(sum[N-1]-sum[N-2])))), float64(frogs(sum, N-2)+int(math.Abs(float64(sum[N-1]-sum[N-3]))))))
}

func main() {
	fmt.Println(frogs([]int{20, 5, 15, 30}, 3))          // 5
	fmt.Println(frogs([]int{10, 30, 40, 20}, 4))         // 30
	fmt.Println(frogs([]int{30, 10, 60, 10, 60, 50}, 6)) // 40
}

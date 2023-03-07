//SOAL
// 1. Pada problem ini kamu harus menemukan total maksimum jumlah bilangan dari deret sebuah integer secara berurutan.

// Sample Test Case

// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

// Output: 6

// Penjelasan: 6 adalah hasil penambahan dari deret 4, -1, 2, 1

package main

import "fmt"

func MaxSequence(arr []int) int {
	maxBilangan := 0
	maxAkhirBilangan := 0

	for _, val := range arr {
		maxAkhirBilangan += val

		if maxAkhirBilangan < 0 {
			maxAkhirBilangan = 0
		}

		if maxBilangan < maxAkhirBilangan {
			maxBilangan = maxAkhirBilangan
		}
	}
	return maxBilangan
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}

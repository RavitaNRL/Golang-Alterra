package main

import "fmt"

func PairSum(arr []int, target int) []int { //pairsum yaitu mengembalikan index dari 2 angka pada array yang jumlahnya sama dengan target atau tujuan yang ditentunkan
	kiri, kanan := 0, len(arr)-1
	for kiri < kanan {
		jumlah := arr[kiri] + arr[kanan]
		if jumlah == target {
			return []int{kiri, kanan}
		} else if jumlah < target {
			kiri++
		} else {
			kanan--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}

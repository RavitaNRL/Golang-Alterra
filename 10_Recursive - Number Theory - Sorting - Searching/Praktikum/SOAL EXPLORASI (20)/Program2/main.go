package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) {
	sort.Ints(productPrice) // untuk mengurutkan productPrice dari yang terkecil ke terbesar.
	jumlah := 0             // dengan nilai awal pada variabel jumlah adalah 0.
	for r := 0; r < len(productPrice); r++ {
		if money >= productPrice[r] {
			money -= productPrice[r]
			jumlah++
		}
	}
	fmt.Printf("Input: money = %d, productPrice = %v\nOutput: %d\n\n", money+jumlah*productPrice[len(productPrice)-1], productPrice, jumlah)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // Input: money = 0, productPrice = [10000 14000 25000 25000] Output: 3
	MaximumBuyProduct(50000, []int{15000, 10000, 12000, 5000, 3000}) // Input: money = 5000, productPrice = [3000 5000 10000 12000 15000] Output: 5
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // Input: money = 2000, productPrice = [1000 2000 2000 3000 10000] Output: 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // Input: money = 4000, productPrice = [2000 2500 3000 7500] Output: 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // Input: money = 0, productPrice = [10000 30000] Output: 0
}

// SOAL
// - Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap i (0 <= i <= n), ans[i] adalah bilangan 1 dalam representasi biner dari i

// Input: n = 2

// Output: [0,1,10]

// Input: n = 3

// Output: [0,1,10, 11]

package main

import "fmt"

func digitBiner(bilanganBulat int) {
	for z := 0; z <= bilanganBulat; z++ {
		fmt.Printf("%b", z)
	}
}

func main() {
	digitBiner(2) // [0 1 1 0]
	fmt.Println()
	digitBiner(3) // [0 1 10 11]
}

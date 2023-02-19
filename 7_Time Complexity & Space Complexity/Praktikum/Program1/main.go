package main

import (
	"fmt"
	"math"
)

func main() {
	var angka int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)

	if PrimeNumber(angka) {
		fmt.Println("Termasuk bilangan prima ")
	} else {
		fmt.Println("Ini bukan bilangan prima, harap input kembali !!")
	}
}

func PrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(num))) //maxDivisior digunakan untuk menyimpan nilai akar dari num.
	for z := 2; z <= maxDivisor; z++ {         // math.sqrt untuk mengembalikan nilai dari int ke float64.
		if num%z == 0 {
			return false
		}
	}

	return true
}

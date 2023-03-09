// 2. Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk melakukan perhitungan perpangkatan (x^n dibaca x pangkat n). Time complexity dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time complexity yang optimal adalah logaritmik.

//Sample Test Cases

// Input : x = 2, n = 3

// Output : 8

// Input : x = 7, n = 2

// Output : 49

package main

import "fmt"

func main() {
	angka := 2
	pangkat := 3
	result := pow(angka, pangkat)
	fmt.Printf("Hasil %d pangkat %d yaitu %d\n", angka, pangkat, result)

	angka = 7
	pangkat = 2
	result = pow(angka, pangkat)
	fmt.Printf("Hasil %d pangkat %d yaitu %d\n", angka, pangkat, result)
}

func pow(angka, pangkat int) int {
	if pangkat == 0 {
		return 1
	}

	if pangkat%2 == 0 {
		z := pow(angka, pangkat/2)
		return z * z
	} else {
		r := pow(angka, (pangkat-1)/2)
		return angka * r * r
	}
}

//buatlah sebuah program untuk menghitung luas trapesium

package main

import "fmt"

func main() {
	var alas_a, alas_b, tinggi float64

	fmt.Print("Masukkan panjang alas a : ")
	fmt.Scan(&alas_a)

	fmt.Print("Masukkan panjang alas b: ")
	fmt.Scan(&alas_b)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	luas := (alas_a + alas_b) * tinggi / 2

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}

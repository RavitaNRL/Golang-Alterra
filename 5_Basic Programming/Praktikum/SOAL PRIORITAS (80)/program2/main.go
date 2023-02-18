// buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap

package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Printf("%d merupakan bilangan genap.\n", angka)
	} else {
		fmt.Printf("%d merupakan bilangan ganjil.\n", angka)
	}
}

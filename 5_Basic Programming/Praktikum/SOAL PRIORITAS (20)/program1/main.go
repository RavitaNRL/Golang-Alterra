//Buatlah sebuah program untuk mencetak segitiga asterik seperti dibawah ini!

package main

import "fmt"

func main() {
	var tinggi int

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	for z := 1; z <= tinggi; z++ {
		for k := 1; k <= tinggi-z; k++ {
			fmt.Print(" ")
		}
		for i := 1; i <= 2*z-1; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

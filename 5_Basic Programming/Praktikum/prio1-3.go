// - buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
//   - Nilai 80 - 100: A
//   - Nilai 65 - 79: B
//   - Nilai 50 - 64: C
//   - Nilai 35 - 49: D
//   - Nilai 0 - 34: E
//   - Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid'

package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai anda : ")
	fmt.Scan(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else if nilai >= 80 && nilai <= 100 {
		fmt.Println("Grade: A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("Grade: B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("Grade: C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}
}

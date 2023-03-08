// SOAL
//Angka Fibonacci adalah serangkaian angka di mana setiap angka adalah jumlah dari keduanya nomor sebelumnya. Beberapa angka Fibonacci pertama adalah: 0, 1, 1, 2, 3, 5, 8, ….
// Buatlah fungsi untuk menghitung angka Fibonacci ke-n (top-down)!

package main

import (
	"fmt"
)

func FiboTopDown(number int, listFibo map[int]int) int {
	if number <= 1 {
		return number
	}

	if value, ok := listFibo[number]; ok {
		return value
	}

	listFibo[number] = FiboTopDown(number-1, listFibo) + FiboTopDown(number-2, listFibo)
	return listFibo[number]
}
func main() {
	list := make(map[int]int)
	fmt.Println(FiboTopDown(0, list))  // 0
	fmt.Println(FiboTopDown(1, list))  // 1
	fmt.Println(FiboTopDown(2, list))  // 1
	fmt.Println(FiboTopDown(3, list))  // 2
	fmt.Println(FiboTopDown(5, list))  // 5
	fmt.Println(FiboTopDown(6, list))  // 8
	fmt.Println(FiboTopDown(7, list))  // 13
	fmt.Println(FiboTopDown(9, list))  // 34
	fmt.Println(FiboTopDown(10, list)) // 55

}

//  soal ! buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!

package main

import "fmt"

func Mapping(slice []string) map[string]int {
	data := make(map[string]int) // perintah untuk menyimpan data (qwe, asd, adi) ke dalam map

	// blok program untuk proses perhitungan string pada slice dan menambahkan data ke dalam map
	for _, s := range slice { // perulangan untuk mengecek setiap elemen pada slice
		data[s]++
	}
	return data
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"qwe", "asd", "adi", "qwe", "qwe"})) // map[qwe:3 asd:1 adi:1]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))               // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                  // map[]
}

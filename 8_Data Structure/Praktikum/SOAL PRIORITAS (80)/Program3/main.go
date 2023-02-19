package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	unique := make(map[int]int)
	var result []int

	// hitung kemunculan setiap angka
	for _, a := range angka {
		n, err := strconv.Atoi(string(a))
		if err != nil {
			continue
		}
		unique[n]++
	}

	// tambahkan angka ke hasil jika hanya muncul sekali
	for n, count := range unique {
		if count == 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1, 2, 3, 4, 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("872584"))     // [7, 2, 5, 4]
}

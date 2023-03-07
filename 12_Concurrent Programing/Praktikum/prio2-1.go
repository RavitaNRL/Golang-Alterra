//SOAL
// Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel (Bersamaan).

// Sample Test Cases
// Input:
// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
// aliqua
// Output:
// e: 1
// i: 1
// o: 1
// t: 1

package main

import (
	"fmt"
	"strings"
)

func hitungChar(text string) map[rune]int {
	// inisialisasi map untuk menyimpan jumlah kemunculan setiap karakter
	JumKarakter := make(map[rune]int)

	// konversi string menjadi slice rune agar dapat diiterasi karakter per karakter
	for _, r := range strings.ToLower(text) {
		// blok program untuk mengabaikan spasi dan karakter selain huruf.
		if r == ' ' || !('a' <= r && r <= 'z') {
			continue
		}
		JumKarakter[r]++
	}

	return JumKarakter
}

func main() {
	text := "Incheon salah satu kota di korea selatan"
	JumKarakter := hitungChar(text)

	for char, count := range JumKarakter {
		fmt.Printf("%c : %d\n", char, count)
	}
}

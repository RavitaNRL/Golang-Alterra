// SOAL
// Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!

package main

import (
	"fmt"
	"strings"
)

func Compare(str1, str2 string) string {
	var result string
	for k := 0; k < len(str1); k++ {
		for e := k + 1; e <= len(str1); e++ {
			// percabangan untuk mengecek str mana yang akan menjadi string atau substring.
			if strings.Contains(str2, str1[k:e]) {
				if len(str1[k:e]) > len(result) {
					result = str1[k:e]
				}
			}
		}
	}
	return result
}

func main() {

	// fungsi dari compare untuk membandingan dua string untuk mengahasilkan substring yang sama.fmt.Println(Compare("AKA", "AKASHI")) // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KIJANG", "KIJANG"))  // KIJANG
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}

// SOAL
// Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!

package main

import (
	"fmt"
	"strings"
)

func Compare(r, z string) string {
	var result strings.Builder // type data string builder memiliki performa lebih cepat dibandingkan dengan type data string. serta dapat membantu dalam melakukan substring.
	for k := 0; k < len(r); k++ {
		for j := 0; j < len(z); j++ {
			if r[k] == z[j] {
				result.WriteByte(r[k]) // writebyte untuk menambahkan byte ke string yg sedan dibangun.
				break
			}
		}
	}
	return result.String()
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA
	// fungsi dari compare untuk membandingan dua string untuk mengahasilkan substring yang sama.
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KIJANG", "KIJANG"))  // KIJANG
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}

//SOAL
// Caesar Cipher
// Buatlah sebuah method dengan parameter `offset` bertipe integer dan `input` bertipe string. Method ini menghasilkan sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan `offset` yang merupakan jumlah pergeseran hurufnya. String `input` diasumsikan berisi huruf kecil saja dan spasi. Sebagai contoh, ketika terdapat huruf `z` yang digeser dengan `offset` sebesar 3 maka huruf hasil pergeseran adalah `c.`

package main

import "fmt"

func caesar(offset int, input string) string {
	// pendeklarasian variabel untuk menampung setiap karakter yang sudah di shift.
	shifted := ""
	//blok program untuk melakukan pergeseran pada setiap karakter yang diberikan menggunakan range input.
	for _, karakter := range input {
		if karakter == ' ' {
			shifted += " "
			continue
		}
		hrufbaru := rune(int(karakter) + offset%26) // menggeser nilai karakter dalam integer dengan offsite yaitu jumlah pergesran nilai karakter yang tdk lebih dari 26.
		if hrufbaru > 'z' {
			hrufbaru = 'a' + hrufbaru - 'z' - 1
		}
		shifted += string(hrufbaru)
	}
	return shifted
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}

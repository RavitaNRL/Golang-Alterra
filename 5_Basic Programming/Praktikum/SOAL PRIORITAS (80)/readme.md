buatlah sebuah program untuk menghitung luas trapesium

package main

import "fmt"

func main() {
	var alas_a, alas_b, tinggi float64

	fmt.Print("Masukkan panjang alas a : ")
	fmt.Scan(&alas_a)

	fmt.Print("Masukkan panjang alas b: ")
	fmt.Scan(&alas_b)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	luas := (alas_a + alas_b) * tinggi / 2

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}

 buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap

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

 - buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
   - Nilai 80 - 100: A
   - Nilai 65 - 79: B
   - Nilai 50 - 64: C
   - Nilai 35 - 49: D
   - Nilai 0 - 34: E
   - Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid'

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

 buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz

package main

import "fmt"

func main() {
	for j := 1; j <= 100; j++ {
		if j%3 == 0 && j%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if j%3 == 0 {
			fmt.Print("Fizz ")
		} else if j%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", j)
		}
	}
}

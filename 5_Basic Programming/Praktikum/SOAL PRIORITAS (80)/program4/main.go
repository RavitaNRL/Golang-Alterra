//buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz

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

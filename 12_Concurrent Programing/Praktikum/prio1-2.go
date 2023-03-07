//SOAL
// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!

package main

import "fmt"

func main() {
	// Membuat channel untuk menyimpan bilangan kelipatan 3
	x := make(chan int)

	// Goroutine untuk mengirim bilangan kelipatan 3 ke channel
	go func() {
		for z := 0; z <= 25; z++ {
			if z%3 == 0 {
				x <- z
			}
		}
		close(x) // Menutup channel setelah selesai mengirim semua bilangan kelipatan 3
	}()

	// Menerima dan mencetak bilangan kelipatan 3 dari channel
	fmt.Println("Berikut ini bilangan kelipatan 3 : ")
	for v := range x {
		fmt.Println(v)
	}
}

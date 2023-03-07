package main

import "fmt"

func main() {
	// Membuat channel dengan buffer sebesar 10 untuk menyimpan bilangan kelipatan 3
	r := make(chan int, 10) // ini merupakan buffer channel dengan buffer sebesar 10

	// blok program goroutine untuk mengirim bilangan kelipatan 3 ke channel
	go func() {
		for y := 0; y <= 20; y++ {
			if y%3 == 0 {
				// Mengirimkan bilangan kelipatan 3 ke channel
				r <- y
			}
		}

		// Menutup channel setelah selesai mengirim semua bilangan kelipatan 3
		close(r)
	}()

	// blok program untuk menerima dan mencetak bilangan kelipatan 3 dari channel
	fmt.Println("berikut bilangan kelipatan 3 dengan buffer channel : ")

	for bil := range r {
		// Mencetak bilangan kelipatan 3 yang diterima dari channel
		fmt.Println(bil)
	}
}

package main

import (
	"fmt"
	"time"
)

// func printKelipatanInputan menerima parameter x berupa bilangan bulat positif
// dan mencetak semua angka kelipatan X dari 0 hingga 20
func printKelipatanInputan(X int) {
	for z := 0; z <= 20; z += X {
		fmt.Printf("%d ", z)
	}
}

func main() {
	// set interval waktu 3 detik
	interval := time.Second * 3

	// memanggil func printKelipatanInputan menggunakan goroutine
	go printKelipatanInputan(3)

	// tunggu selama interval waktu yang telah ditentukan
	time.Sleep(interval)
}

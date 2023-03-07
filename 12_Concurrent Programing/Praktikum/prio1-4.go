// soal
// Buatlah  program yang yang menerapkan mutex! Jenis program yang dibuat bebas (contoh: perhitungan faktorial).

package main

import (
	"fmt"
	"sync"
)

var (
	//untuk menentukan jumlah goroutine yang akan dijalankan.
	y = 5
	r sync.Mutex
)

func faktorial(k int) int {
	if k == 1 {
		return 1
	}
	return k * faktorial(k-1)
}

func main() {
	// variabel wg bertype data sync.WaitGroup => untuk menunggu goroutine yang sedang berjalan sampai selesai sebelum program berakhir.
	var wg sync.WaitGroup

	// menambahkan nilai pada variabel y ke waitgroup yang akan dijalankan oleh goroutine.
	wg.Add(y)

	for z := 1; z <= y; z++ {
		go func(number int) {
			defer wg.Done()
			r.Lock()
			result := faktorial(number)
			fmt.Printf("Factorial of %d is %d\n", number, result)
			r.Unlock()
		}(z)
	}
	wg.Wait()
}

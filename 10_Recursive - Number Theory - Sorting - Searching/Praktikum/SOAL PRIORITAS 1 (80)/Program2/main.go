//SOAL
// Buatlah program di Golang yang dapat mengurutkan barang berdasarkan jumlah kemunculannya. Jika ada barang yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan barang tersebut!

package main

import (
	"fmt"
	"sort"
)

//pendaklarisan struct dengan nama struct yaitu pair.
type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// membuat sebuah map untuk menampung atau menghitung jumlah kemunculan setiap barang.
	// nama map yaitu jumlah.
	jumlah := make(map[string]int)
	// akan melakukan perulangan sebanyak jumlah item yang ada.
	for _, item := range items {
		jumlah[item]++
	}

	//membuat slice dari struct pair dengan nama itemCounts untuk menampung nama barang dan jumlah kemunculannya.
	itemCounts := make([]pair, 0, len(jumlah))
	for name, count := range jumlah {
		itemCounts = append(itemCounts, pair{name, count})
	}

	// mengurutkan slice berdasarkan jumlah kemunculan dari yang terbesar
	sort.Slice(itemCounts, func(z, k int) bool {
		return itemCounts[z].count < itemCounts[k].count
	})

	return itemCounts
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// output: [{golang 1} {ruby 2} {js 4}]

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// output: [{C 1} {D 2} {B 3} {A 4}]

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// output: [{football 1} {basketball 1} {tenis 1}]

}

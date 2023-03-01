//SOAL
// Buatlah program di Golang yang dapat mengurutkan barang berdasarkan jumlah kemunculannya. Jika ada barang yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan barang tersebut!

package main

import (
	"fmt"
	"sort"
)

//pendaklarasian struct
type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	itemCount := make(map[string]int)
	for _, item := range items {
		if _, ok := itemCount[item]; ok {
			itemCount[item]++
		} else {
			itemCount[item] = 1
		}

		pairs := make([]pair, 0, len(itemCount))
		j := 0
		for k, v := range itemCount {
			pairs = append(pairs, pair{k, v})
			j++

		}

		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})

		return pairs
	}
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	//golang -> 1 ruby -> 2 js -> 4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "A", "B", "A", "D", "D"}))
	//A -> 5 B -> 3 C -> 1 D -> 2
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))

}

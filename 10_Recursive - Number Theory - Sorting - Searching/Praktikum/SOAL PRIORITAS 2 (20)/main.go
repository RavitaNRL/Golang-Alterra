package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, kartu := range cards {
		//menggunakan percabangan untuk mengecek apakah kartu yang ada di tangan sama dengan kartu yang ada di deck.
		if kartu[0] == deck[0] || kartu[1] == deck[0] || kartu[0] == deck[1] || kartu[1] == deck[1] {
			return kartu // jika ada kartu yang sama antara di tangan dengan deck maka akan mengulangi kartu dengan ketentuan output [x,y]
		}
	}
	//jika tidak ada kartu yang maka akan tercetak "tutup kartu"
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}, []int{3, 6}}, []int{4, 3}))
	// [6, 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	// "tutup kartu"
}

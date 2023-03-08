package main

import "fmt"

func digitBiner(bilanganBulat int) {
	for z := 0; z <= bilanganBulat; z++ {
		fmt.Printf("%b", z)
	}
}

func main() {
	digitBiner(2) // [0 1 1 0]
	fmt.Println()
	digitBiner(3) // [0 1 10 11]
}

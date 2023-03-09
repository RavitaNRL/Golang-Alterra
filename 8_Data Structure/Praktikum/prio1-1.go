package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	isName := make(map[string]bool) // merupakan map yang digunakan untuk menyimpan nama pada array A dan B
	mergedArray := []string{}       // perintah untuk menggabungkan array A dan B

	//blok program untuk array A
	for _, name := range arrayA {
		if !isName[name] {
			isName[name] = true
			mergedArray = append(mergedArray, name)
		}
	}

	//blok program untuk array B
	for _, name := range arrayB {
		if !isName[name] {
			isName[name] = true
			mergedArray = append(mergedArray, name)
		}
	}
	return mergedArray
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{"hwoarang"}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Masukkan kata: ")
	fmt.Scanln(&input)
	if iniPalindrome(input) {
		fmt.Printf("%s ini merupakan palindrome.\n", input)
	} else {
		fmt.Printf("%s ini bukan palindrome.\n", input)
	}
}

func iniPalindrome(kata string) bool {
	kata = strings.ToLower(kata)
	for z := 0; z < len(kata)/2; z++ {
		if kata[z] != kata[len(kata)-z-1] {
			return false
		}
	}
	return true
}

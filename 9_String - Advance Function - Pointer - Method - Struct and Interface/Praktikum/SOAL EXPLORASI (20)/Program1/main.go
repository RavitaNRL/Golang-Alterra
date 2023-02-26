package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	kunci := 5 // Kunci pergeseran untuk algo substitusi chipper
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			// Pergeseran karakter sebanyak kunci
			char = 'a' + (char-'a'+rune(kunci))%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+rune(kunci))%26
		}
		nameEncode += string(char)
	}
	s.nameEncode = nameEncode // s.nameEncode untuk menyimpan hasil dari karakter pada variabel nameEncode.
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	key := 3 // Kunci pergeseran
	for _, char := range s.nameEncode {
		if char >= 'a' && char <= 'z' {
			// Pergeseran karakter sebanyak kunci
			char = 'a' + (char-'a'-rune(key)+26)%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'-rune(key)+26)%26
		}
		nameDecode += string(char)
	}
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}

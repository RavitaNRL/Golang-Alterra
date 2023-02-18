  Buatlah sebuah program untuk mencetak segitiga asterik seperti dibawah ini!

package main

 import "fmt"

 func main() {
 	var tinggi int

 	fmt.Print("Masukkan tinggi segitiga: ")
 	fmt.Scan(&tinggi)

 	for z := 1; z <= tinggi; z++ {
 		for k := 1; k <= tinggi-z; k++ {
 			fmt.Print(" ")
 		}
 		for i := 1; i <= 2*z-1; i++ {
 			fmt.Print("*")
 		}
 		fmt.Println()
 	}
 }

 buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka,
 input 26
 output 1 , 2 ,13, 26

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	fmt.Println("Faktor bilangan dari", num, "adalah: ")

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}

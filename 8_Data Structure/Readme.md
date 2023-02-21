array merupakan struktur data yang terdapat kumpulan element didalamnya dimana indeks pada array selalu dimulai dari 0, selain itu array hanya dapat berisi satu jenis variabel dengan ukuruan indeks yang tetap. 
sehingga array dapat dikatakan sangat penting dalam sebuah program karena digunakan untuk melakukan pengolahan data yang efesien dan efektif. pada array juga terdapat looping.

jenis data pada array sama seperti umumnya yaitu : 
a. numerik
b. boolean
c. string

contoh program looping array : 

	 program keempat = looping array
	primes := [5]int{2, 3, 5}
	//cara pertama
	 for indeks := 0; indeks < len(primes); indeks++ {
		fmt.Println(primes[indeks])

	//cara kedua
	// for indeks, element := range primes {
	// 	fmt.Println(indeks, "=>", element)
	// }


dengan ouput :

cara pertama : 
PS C:\Users\Ravita\Desktop\8_Data Structure\Praktikum> go run "c:\AKU2\STUPEN - GOLANG PROGRAMMING - ALTERRA\MINGGU 2\PERT 5 -- SECTION 8\PROGRAM DI VIDEO\array\main.go"
2
3
5
0
0
PS C:\Users\Ravita\Desktop\8_Data Structure\Praktikum> 

cara kedua : 
PS C:\Users\Ravita\Desktop\8_Data Structure\Praktikum> go run "c:\AKU2\STUPEN - GOLANG PROGRAMMING - ALTERRA\MINGGU 2\PERT 5 -- SECTION 8\PROGRAM DI VIDEO\array\main.go"
0 => 2
1 => 3
2 => 5
3 => 0
4 => 0
PS C:\Users\Ravita\Desktop\8_Data Structure\Praktikum> 

slice merupakan struktur data yang hampir sama dengan array, yang berbeda hanya pada ukuran indeks artinya slice hanya dapat menampung satu jenis variabel dan ukuran dari slice ini dapat berubah ubah 

contoh program slice : 

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var primes = [5]int{2, 3, 5, 7, 11}

	var part_primes []int = primes[1:4]

	fmt.Println(reflect.ValueOf(part_primes))
	fmt.Println(part_primes)
}

output program : 

PS C:\AKU2\STUPEN - GOLANG PROGRAMMING - ALTERRA\MINGGU 2\PERT 5 -- SECTION 8\PROGRAM DI VIDEO> go run "c:\AKU2\STUPEN - GOLANG PROGRAMMING 
- ALTERRA\MINGGU 2\PERT 5 -- SECTION 8\PROGRAM DI VIDEO\slice\main.go"

[3 5 7]
[3 5 7]
PS C:\AKU2\STUPEN - GOLANG PROGRAMMING - ALTERRA\MINGGU 2\PERT 5 -- SECTION 8\PROGRAM DI VIDEO> 


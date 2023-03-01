Rangkuman section 9 - minggu ke-3 

pada golang string dapat digunakan untuk :

1. len string => untuk menghitung panjang string pada sebuah variabel. 
len tidak hanya bisa berupa dalam bentuk string namun juga bisa dalam bentuk array, slice dan map.

2. compare string => untuk membandingkan string 1 dengan string lainnya
contoh :

package main 

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	str2 := "World"
	str3 := "Hello"
	str4 := "Hello"
	fmt.Println(str1 == str2) // akan mencetak false karena str1 tidak sama dengan str2
	fmt.Println(str3 == str4) // akan mencetak true karena str1 sama dengan str3

}

dari code diatas ketika program di eksekusi maka akan membandingan data pada variabel 1 dengan variabel lainnya. 

3. contains string => hampir sama dengan compare string yang membedakan yaitu cara kerja pada constains string yaitu membandingan string 1 apakah bagian dari string yang kedua.

selain string pada golang juga ada advance function ada beberapa jenis yaitu :

1. variadic function => untuk memanggil paramter dengan jumlah yang berbeda-beda.
dengan menggunakan variadic function ini penulisan code menjadi lebih simple dan dinamis daripada 
menggunakan slice. 
perbandingan penulisan code program jika menggunakan slice dan variadic function :

>> penggunaan slice :

package main 

import "fmt"

func main() {
 fmt.Println(sum([]int{12, 4}))
 fmt.Println(sum([]int{12, 4, 3}))
 fmt.Println(sum([]int{12, 4, 3, 5}))
}

func sum(data []int) int {
	result := 0
	for _, v := range data { // untuk perulangan pada setiap data.
		result += v // untuk menambahkan iterasi dengan menambahkan result lama dgn yng baru.
	}
	return result
}

>> penggunaan variadic function :
dengan menggunakan jenis ini pada function sum tanda slice atau [] diganti dengan titik 3 (...)

package main 

import "fmt"

func main() {
fmt.Println(sum(12, 4))
fmt.Println(sum(12, 4, 3))
fmt.Println(sum(12, 4, 3, 5))
}

func sum(data ...int) int {
	result := 0
	for _, v := range data { // untuk perulangan pada setiap data.
		result += v // untuk menambahkan iterasi dengan menambahkan result lama dgn yng baru.
	}
	return result
}


tidak hanya itu pada golang juga terdapat struct. struct pada golang merupakan call objek atau pemanggilan objek.
dimana pemanggilan ini merupakan pengganti OPP. karena pada golang tidak menggunkan OOP maka untuk melakukan gaya penulisan code program 
seperti OOP dapat menggunakan struct. 

cara pendaklarasian struct pada golang :

type student struct { // student merupakan nama struct yang digunakan.
	FirstName string // fisrtname => field pada struct / variabel.
	LastName  string
	Age       int
}

untuk gaya penulisan struct ada 3 yaitu :

>> cara pertama : 

var Person1 student // deklarasi variabel Person1 dengan tipe data struct student.
	Person1.FirstName = "ravita"
	Person1.LastName = "asmi"
	Person1.Age = 20
	fmt.Println("cara pertama : ", Person1) // maka akan mencetak data pada field atau variabel di person1.

pada cara pertama ini dimulai dengan menuliskan kata var pada variabel. kemudian untuk tiap field diberi nama struct lalu titik. 

>> cara kedua :

var Person2 = student{
		FirstName: "ravita",
		LastName:  "nurul",
		Age:       17,
	}
	fmt.Println("cara kedua : ", Person2)

pada cara kedua ini masih sama seprrti cara pertama dengan dimulai menuliskan kata var didepan variabel yang kita gunakan.
namun pada tiap field struct tidak perlu diberi nama struct. 

>> cara ketiga :

Person3 := student{"eza", "ravita", 21}
	fmt.Println("cara ketiga : ", Person3)

pada cara ketiga ini tidak perlu diberi kata var pada awal variabel dan tidak
menuliskan field yang digunakan pda struct. namun kita harus mengingat urutan tiap field yang sudah dibuat
jika menggunaakn cara ketiga. 
walaupun mudah dengan cara ketiga terdpat kelemahan, jika kita nemambahkan field baru pada struct 
maka akan error jika tidak di tambahkan data pada field baru tersebut. namun jika menggunakan cara pertama atau kedua
walaupun tidak menambahkan data pada field baru tersebut tidk akan terditeksi error. 

 



 


// soal
// Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut receiver kepada struct Car yang memiliki property type dan fuelIn

package main

import "fmt"

//pendaklarasian struct
type Mobil struct {
	Type   string
	fuelin float64 //  untuk menampung bahan bakar yang ada dengan satuan liter
}

func (mobil Mobil) PerhitunganJarakTempuh() float64 { // PerhitunganJarakTempuh merupakan nama method untuk menghitung jarak yang ditempuh dengan bahan bakar 1.5 liter
	JarakTempuh := mobil.fuelin * 1.5
	return JarakTempuh // variabel untuk menyimpan perhitungan jarak yang dapat ditempuh
}

func main() {
	myMobil := Mobil{
		Type:   "Station Wagon",
		fuelin: 15}
	JarakTempuh := myMobil.PerhitunganJarakTempuh()

	fmt.Println("Jarak yang dapat ditempuh oleh mobil", myMobil.Type, "dengan bahan bakar yang sedang terisi 1,5 Liter adalah", JarakTempuh, "km")
	//kegunaan dari myMobil.Type adalah untuk memanggil type mobil agar tercetak pada bagian output program.

}

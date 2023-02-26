// soal
// Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut receiver kepada struct Car yang memiliki property type dan fuelIn

package main

import "fmt"

type Mobil struct {
	Type   string
	FuelIn float64 // variabel FuelIn untuk menampung bahan bakar yang ada dengan satuan liter
}

func (mobil Mobil) CalculateMileage() float64 { // method CalculateMileage menghitung jarak yang ditempuh dengan bahan bakar 1.5 liter
	mileage := mobil.FuelIn * 1.5
	return mileage // variabel untuk menyimpan perhitungan jarak yang dapat ditempuh
}

func main() {
	myMobil := Mobil{Type: "Station Wagon", FuelIn: 30.5}
	mileage := myMobil.CalculateMileage()

	fmt.Println("Jarak yang dapat ditempuh oleh mobil", myMobil.Type, "dengan bahan bakar yang sedang terisi 1,5 Liter adalah", mileage, "km")
	//kegunaan dari myMobil.Type adalah untuk memanggil type mobil agar tercetak pada bagian output program.

}

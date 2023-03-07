// SOAL
//Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!

package main

import "fmt"

// struct dibawah ini merepresentasikan data kendaraan yang terdiri dari 2 field yaitu total roda dan kecepatan per jam dengan masing masing field bertype data int.
type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

// struct dibawah ini menjelaskan data mobil dengan menggunakan fitur embedding (memasukan type data struct kedalam type struct lainnya)
type mobil struct {
	// hasil embedding dari struct kendaraan.
	kendaraan
}

//func mulai dijalankan dengan menambahkan kecepatan sebanyak 10.
func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

//blok program dibawah ini untuk menambahkan kecepatan mobil dengan menambahkan kecepatan baru.
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

//blok program dibawah ini untuk menampilkan kecepatan mobil.
func main() {
	//membuat object atau variabel baru nama mobilcepat dengan tipe data mobil.
	mobilcepat := mobil{}

	//memanggil func yang berjalan pada object mobilcepat sebanyak 3kali.
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	//membuat object atau variabel baru nama mobillamban dengan tipe data mobil
	mobillamban := mobil{}

	//memanggil func yang berjalan pada object mobillamban sebanyak 1kali.
	mobillamban.berjalan()

	// tambahkan code program untuk menampilkan kecepatan mobilcepat dan mobillamban
	//menjadi :
	fmt.Println("kecepatan yang dihasilkan mobilcepat :", mobilcepat.kecepatanperjam)
	fmt.Println("kecepatan yang dihasilkan mobillamban :", mobillamban.kecepatanperjam)

}

//NOTE :
// pada code program diatas hanya kurang pemberian komentar pada setiap fungsi atau struktu data untuk menjelaskan maksud dari code program tersebut.

//SOAL
// Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
//     - Berapa banyak kekurangan dalam penulisan kode tersebut?
//     - Bagian mana saja terjadi kekurangan tersebut?
//     - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.

package main

import "fmt"

type user struct {
	id int
	// username int
	// password int
	// seharus nya pada kedua variabel diatas bertype data string
	//menjadi :
	username string
	password string
}

// untuk penaman struct lebih baik di kombinasikan dengan huruf kapital agar lebih mudah untuk dibaca.
// type userservice struct
//menjadi :
type userService struct {
	// t []user
	// sebaiknya pada variabel t diberi nama yang lebih jelas yaitu
	// menjadi variabel dataUser, karena berisikan data mengenai nama pengguna.
	dataUser []user
}

// untuk penaman func lebih baik di kombinasikan dengan huruf kapital agar lebih mudah untuk dibaca.
// func (u userservice) getallusers() []user
//menjadi :
func (u userService) getAllUsers() []user {
	return u.dataUser
}

// untuk penaman func lebih baik di kombinasikan dengan huruf kapital agar lebih mudah untuk dibaca.
// func (u userservice) getuserbyid(id int) user
//menjadi :
func (u userService) getUserByID(id int) user {
	// for _, r := range u.dataUser {
	// 	if id == r.id {
	// 		return r
	// 	}
	// }

	// return user{}
	// sebaiknya variabel r diganti menjadi users, karena berisikan data mengenai nama pengguna dan lebih deskriptif.
	//manjadi :
	for _, pengguna := range u.dataUser {
		if id == pengguna.id {
			return pengguna
		}
	}

	return user{}
}

// tambahkan func main agar dapat memanggil func yang telah dibuat dan memberikan output

func main() {
	// membuat slice user baru
	users := []user{
		{1, "onyourbi", "password"},
		{2, "matchaw", "secret"},
		{3, "mee", "267191"},
	}

	// membuat instance dari userService dan memasukkan slice user baru
	userService := userService{users}

	// mencetak seluruh user
	fmt.Println("Seluruh pengguna:", userService.getAllUsers())

	// mencetak user dengan id = 2
	fmt.Println("pengguna dengan user id 2 :", userService.getUserByID(2))
	fmt.Println("pengguna dengan user id 1 :", userService.getUserByID(1))
}

// NOTE :
// jangan lupa untuk memberikan komen pada setiap fungsi atau struktu data untuk menjelaskan maksud dari code program tersebut.
// sehingga orang lain mudah untuk membaca nya.

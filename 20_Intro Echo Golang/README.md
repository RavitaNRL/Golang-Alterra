MATERI ECHO GOLANG
**************************

THIRD PARTY LIBRARY GOLANG
***************************
thrid party / library
=> suatu kumpulan code yang memiliki fungsi tertentu, dimana fungsi tersebut dapat dipanggil oleh program lain.
positif :
=> menghemat waktu dalam membangun sistem 
negatif :
=> jika banyak menggunakan library dari luar ketika library yang kita pakai itu mengalami update maka program kita jga harus diupadate. 
 library golang yang populer :
1. echo 
=> untuk membantu dalam pengembangan web / rest api
2. go kit
3. gorm.io
=> untuk orm (oprasi databse)


ECHO GOLANG
************
echo golang
=> web framework golang yang punya performa tinggi, dapat dipasang-pasangkan dan minimalist. 

kelebihan :
1. optimaze router
=> dalam penggunaan echo router nya powerfull dalam penggunaan alokasi memori
2. middleware
=> contohnya gui. sehingga kita dapat kastem.
3. data rendering
=> memberikan suatu respon pada sebuah request dengan mudah. 
4. scalable
5. data binding


menggunakan restAPi dengan framework echo 
********************************************
1. buat folder
2. open folder didalam vscode
3. membuat inisialisasi agar library dapat diakses dengan mudah
	1. buka terminal
`	2. ketikan pada terminal " go mod init namafolder". go mod init ( menginisialisasi)
	3. ketika berhasil maka akan muncul file go mod. 

4. import library echo. dapat diakses di website resmi echo golang. 
	1. gunakan perintah go get. (ada di website echo : go get github.com/labstack/echo/v4). nanti letak library echo ada di gopath.
5. ketika import libary echo berhasil maka akan ada tambahan code pada file go mod dan ada tambahan file yaitu go sum.
6. buat file, misal main.go ( untuk membangun restapi sederhana, sebelum dilakukan pemecahan).
7. membuat routing ( jalur untuk mengakses url, misal localhost:8000/user )  
maka nanti project akan mengakses atau memanggil controler dengan method GET. 
8. ketika code program telah selesai dibuat maka ketik di terminal " go run main.go atau go run nama file.go "
9. maka akan muncul tulisan echo pada halaman terminal, artinya program dapat dilakukan testing.
10. untuk testing dapat gunakan postmant
11. rubah mathod menjadi GET lalu ketik port :8000 pada url (http://localhost:8000/hello). 
12. sebelum melakukan testing dengan postman, move folder ke folder src di folder go.








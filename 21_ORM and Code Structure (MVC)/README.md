MATERI ORM 
**************

DEFINISI ORM
**************
orm
=> membantu dalam konversi data, misal pada saat mengakses database.

kelebihan : 
1. less repetitive query
=> mudah dalam mendapatkan data pada sebuah database. dengan cara membuat struct yang field nya sama seperti field pada tabel di database.

2. fetch data
=> otomatis terkonvert pada struct yang sudah dibuat dari database.

3. screening data 
=> mempermudah validasi data pada saat memasukan data di struct dari database. 

4. cache query
=> dapat mengakses data yang sama kemudian disimpan pada penyimpanan sementara pada cache dan ketika orm butuh datanya tinggal manggil data yang sudah tersimpan di cache sehingga waktu yang dibutuhkan lebih cepat. 


kelemahan :
1.  cost the overhead process
=> harus menggunakan library pihak ketiga sehingga menambahkan code khusus dari library tersebut.

2. un-necessary reletinship data
=> dapat melakukan eksekusi query yang tidak diinginkan, sehingga menambahkan cost.

3. specific sql function
=> menjadi ketergantungan. 

GORM
******
gorm
=> fungsi orm yang RDMS atau sql pada go dan dapat dipasangkan dengan framework yang sudah dibuat oleh develop.

koneksi database dengan GORM
*******************************
1. membutuhkan user name dan password sesuai database. 


DATABASE MIGRATION
************************
=> untuk jalan untuk apa yang ditambahkan di versi yang baru. misal ingin menambahkan field pada database melalui struct.

INSTAL GORM 
*******************
1. install orm library go melalui website gorm (https://gorm.io/docs/)
=>
	1. buka terminal kemudian ketikan " go get -u gorm.io/gorm " akan mendownload package yang akan dimasukan ke goroot project.
	2. jika sudah install kembali driver gorm " go get -u gorm.io/driver/sqlite " karena menggunakan mysql maka /sqlite ganti dengan /mysql
	3. go get -u gorm.io/driver/mysql

KONEKSI GORM KE DATABASE 
**************************
1. dapat dilihat di dokumentasi GORM pilih menu CONNECTING TO DATABASE.
2. buat database dulu dengan quey sql, bisa menggunakan phpmyadmin.
3. melakukan test connecting yaitu lihat hostname dan port serta password.
4. jika sudah membuat database, melakukan koneksi dengan perintah yang ada di dokumentasi GORM.
kemudian samakan data nya seperti hostname, port dan password. dengan membuat func initDb
5. jika sudah buat struct di file main.go, lakukan migrtion database (dengan type auto migartion).
dengan cara mebuat func migration. 
6. jika code sudah lengkap lakukan testing dengan postman.


MVC
**********
MVC (MODEL, VIEW, CONTROLER)
=> 
1. MODEL
=> berisi seperti struct (modeling database yang akan diolah).
2. CONTROLER
=> isi dari logika object project atau pengolahan.
3. view
=> bagian paling awal. jika di restful API tidak terlalu banyak, seperti main. 


note : 
1. pada inisialisai database membuat folder baru dengan nama config. jadi code program untuk inisialisasi database ditempatkan di folder config. 
2. buat route untuk bagian router. 










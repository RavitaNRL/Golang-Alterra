MATERI MIDDLEWARE
*********************

APA ITU MIDDLEWARE
*******************
MIDDLEWARE
=> suatu entitas yang dipasangkan dalam sebuah server.
konsep nya ketika client membutuhkan respon maka server butuh layer atau middleware.
dimana layer tersebut berisi fungsi-fungsi khusus untuk membantu komukasi data dari sisi client ke sis server. 

IMPLEMENTASI 
****************
pada saat terjadi komunikasi tersebut yaitu ketika client dari web akan mengakses sebuah rest API dalam server. ada beberapa layer sebelum client berhasil mengakses api server itu :
1. logging Middleware
=> untuk mencatat log apa saja yang keluar masuk ke dalam server.
2. session Middleware
=> status apakah client masih aktif dalam mengakses server atau tidak.
3. auth Middleware
=> untuk autentification apakah client berhak mengakses server atau tidak dengan alamat url tersebut.
4. custom Middleware
=> untuk membuat custom sebelum melakukan komunikasi data antara client dengan server.

TOOLS MIDDLEWARE
*******************
1. negroni
2. echo
3. interpose
4. alice

note : diperbolehkan membuat library sendiri.


ECHO MIDDLEWARE
******************
isi dari echo middleware :
1. basic auth 
2. key auth
3. JWT
4. logger
5. dst (dapat dilihat di dokumentasi echo middleware "  echo.labstack.com/middleware/ " )

jenis middleware pada echo :
1. echo#Pre() 
=> digunakan ketika akan mensetup sebuah middleware maka akan tereksekusi sebelum root masuk kedalam controller.

2. echo#Use()
=> digunakan ketika middleware akan dieksekusi setelah mendapatkan akses dari conteks echo.

contoh mengimplementasikan echo#Pre() dengan https redirect :
=> code program 
e := echo.New()
e.Pre(middleware.HTTPSRedirect()) 

logika 
=> 
e.pre 
=> pre menandakan sebelum masuk ke dalam urutan proses. 

contoh lainnya :
ketika mengakses http://labstack.com
maka akan di redirected oleh middleware maka akan terjadi 
https://labstack.com.

note : untuk saat ini restAPI jika bisa di protocol https, karena lebih aman / tingkat keamanan tinggi dari pada menggunakan protocol http.


LOG MIDDLEWARE
****************
log middleware
=> salah satu jenis middleware yang berfungsi untuk mencatat informasi apa saja yang terjadi didalam server terutama didalam request.
dapat digunakan untuk pendeteksi serangan sistem, dll.


IMPLEMENTASI LOG MIDDLEWARE
********************************
1. buka dokumentasi dari log middleware lalu pilih menu logger.
2. untuk mengintegrasi tambahankan " e.Use(middleware.logger()) "
3. di vscode buat folder middleware, kemudian buat file log_middleware.go (nama file bebas).
4. tambahkan package sesuai dengan nama folder (package middleware).
5. buat func logMiddleware ( untuk menerima echo yang diatur atau set di folder route sehingga dibuatkan parameter).
6. kemudain dibuatkan pointer func logMiddleware (e *echo.Echo) ( untuk menunjukan alamat memori di router yang digunakan, dimana nantinya akan dikirim ke logmiddleware).
7. kemudian tambahkan middleware log " e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
})) " 
8. pada bagian format bisa dirubah atau disesuaikan, karna bagian format digunakan untuk mengambil data. data apa saja yang dapat di ambil dapat dilihat pada dokumentasi " https://echo.labstack.com/middleware/logger/ " bagian configurasi.
9. memasangkan logmiddleware kedalam route program dengan cara memanggil fungsi yang ada pada logmiddleware. dengan cara mengimport package logmiddleware kedalam file route.go dengan cara " namamodul/namapackage " ( untuk nama modul dapat dilihat pada file go.mod ).
10. kemudain tambahkan inisialisasi disamping import logmiddleware. 
11. jika sudah panggil logmiddleware dengan perintah " m.logMiddleware(e) ".
12. testing dilakukan di postman.
13. jika berhasil maka di menu output vscode akan tercatat atau terpanggil perintah didalam format. 


AUTH MIDDLEWARE
*******************
auth middleware
=> suatu proses pada API dalam mengauntentifikasi suatu user. dengan auth middleware dapat menentukan 
apakah request dari client dapat / berhak mengakses API atau tidak.

contoh :
ketika user ingin membuka aplikasi maka harus login. pada saat login server akan mengirimkan token untuk memberikan penanda bahwa sudah login ke sistem
hal tersebut gambaran cara kerja dari auth middleware. 

ada 2 authentication middleware :
1. basic authentication
2. JSON web token


BASIC AUNTHENTICATION
**************************
basic authentication
=> suatu proses dalam rest API dengan mengirimkan data username dan password melalui data headers.
aturan :
1. mengirimkan key Authorization dengan value basic + base64encode dari username:password.
2. akan muncul hasil header generate berupan authorization : basic ( hasil input base64encode ).
3. encode dari password dan username yang nantinya akan dikirimkan di headers pada postman.


IMPLEMENTASI BASIC AUTHENTICATION
***********************************
1. tambahkan file baru pada folder middleware di vscode (auth_basicMiddleware.go).
2. tambahkan package middleware.
3. buat func basicAuthDB(username, password string, c echo.Context). untuk paramater sesuai dengan dokumentasi pada web echo. (username, password string, c echo.Context).
4. melakukan pengecekan username dan password pada database. sebelum itu cek terlebih dahulu pada bagian struct di folder models. jika belum ada field username dan password pada struct tersebut maka tambahkan.
5. jika sudah ditambahkan di struct maka running terlebih dahulu agar field yang ditambahkan di struct juga bertambah di databse nya.
6. lalu lakukan cek pada file auth_basicMiddleware apakah sudah sama username dan password di databse nya. cara nya :
	1. buat var ( var user models.Users )
	2. mengakses database didalam config ( config.DB.Where("username = ? AND password = ?", username, password)First(&user). note : untuk paramater di DB.Where sama kan dengan kolom di databse)
	3. panggil error dengan membuat variabel err. variabel err ditambahkan pada saat mengakses datasabe. 
	4. buat percabangan if (err != nill) { retrun false, err }. (jadi jika terdetect tidak teraunthentication maka tidak akan bisa mengakses API).
	5. akan mereturn true, nil. ( berhasil dan ada sebuah data yang masuk ke data user).
7. lakukan integrasi kedalam route.go 
8. membuat gruop baru untuk auth basic (eAuthBasic := e.Group("/auth")) 
9. gunakan route dari authbasic (eAuthBasic.GET("/users", controller.GetUserController)
10. perintah untuk menggunakan authbasic (eAuthBasic.Use(midd.BasicAuth(midd.BasicAuthDB)) (nanti akan muncul otomatis di bagian import "github.com/labstack/echo/middleware" dan disamping ini bikin inisial yaitu mid)
11. run program dan testing dengan postman.
12. agar berhak megakses API yaitu menambahkan header atau token berupa encode.
13. cara menambahkan header di postman
	1. masukan token berupa encode
	2. key yang digunakan Authorization
	3. value yang digunakan Basic 
	4. masukan token encode setelah basic yaitu Basic SPASI token encode.
	5. untuk mendapatkan password isi manual dlu di databse nya kemudian covert ke encode melalui website encode64 online.



JWT WEB TOKEN 
*******************
jwt web token
=> authentication seperti basic auth namun memiliki algoritma yang berbeda. tidak perlu mengakses dari database cukup mengakses dari server guna untuk memvalidasi apakah JWT ini valid atau tidak. 

struktur dari jwt :
1. terdiri dari beberapa token dengan penulisan tokennya dipisahkan dengan titik. 
2. untuk masing-masing token yang digunakan berasal dari encode64.

hal yang diperhatikan ketika ingin melakukan testing pada postmant :
1. key yang digunakan Authorization dengan value Bearer (spasi) token.

IMPLEMENTASI JWT 
********************
1. membuat key yang akan digunakan untuk melakukan algoritma enkripsi menggunakan jwt.
2. sebelum membuat key, buat folder baru " constants ". jika sudah buat file baru " constants.go ". 
3. jika sudah buat package " package constants " (nama package sesuai dengan nama folder).
4. buat secret jwt dengan perintah " const SECRET_JWT = "123" ( jika kode secret bocor ke public maka orang luar dapat mengakses secara bebas ke server).
5. kemudian buat file di folder middleware untuk jwt " jwt_middleware.go ". (untuk mengisi / mengenerate token dan mempilot data (mengambil user id / name )).
	1. buatlah package " package middleware ".
	2. buatlah func untuk melakukan generate token "func CreateToken". (akan berisi pilot data (identitas bahwa suatu token itu milik userID berapa (pilot data tdk bole berisi data penting seperti , nohp, password, email)). 
	3. pada func CreateToken berisi 2 paramter untuk generate " func CreateToken(userId int, name string) "
	4. akan melakukan retrun token " func CreateToken(userId int, name string) (string, error) { "
	5. menyimpan data pilot ke dalam jwt menggunakan jwt claims, dimana jwt berasa dari library.
	6. kemudian buat variabel untuk menyimpan jwt claim " claims := jwt.MapClaims{} "
	7. berisikan data apa saja yang akan disematkan " claims["userId"] = userId "
	8. lakukan hal yang sama seperti lanngkah 7 " claims["name"] = name "
	9. tambahkan sebuah exp " claims["exp"] = time.Now().Add(time.Hour * 1).Unix() " (untuk masa berlaku token yang digunakan dimaan nantinya akan otomatis, jdi nanti dibagian sistem atau echo akan melakukan eksekusi jika token sudah melebihi waktu yang diberikan maka akan meretrunt exp jadi token tdk bisa digunakan).
	perintah pada no 9 untuk mengatur jangka waktu token hanya 1 jam.
	10. untuk header membuat variabel " token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) "
	11. melakukan generate token "return token.SignedString([]byte(constans.SECRETE_JWT))" 
	12. lalu import "nama module/constans"
	13. note : create token pada file jwt_middleware tdk perlu dikirimkan ke database karna nnti akan otomatis dikriimkan ke server.
6. jika sudah selesai pada file jwt_middleware. cek folder route dan tambahkan data baru mengenai jwt tadi.
7. penambahan data yaitu "eJwt := e.Group("/jwt") " dan "eJwt.GET("/users", controllers.GetUserController)
8. jika sudah buatkan func baru mengenai login. dan pada bagian file route.go tambahkan " e.POST("/login", controllers.LoginController) ".
9. pada bagian file controler tambahkan func LoginControler dengan mengirimkan echo context dan mereturn error.
10. pada func login lakukan pengecekan ke database. 
11. jika sudah maka pada file controler lakukan mengenerate di jwt dengan membuat variabel" token, err := middleware.CreateToken(user.ID, user.Name)
12. lakukan import "namamodul/middleware " pada file controller. 
13. pada file models, tambahkan token (karena pada praktik itu menyimpan di databse maka dilakukan pembuatan struct baru untuk menyimpan token "type UserResponse Struct { Token string ..... " 
14. kemudian pada bagian file controler dibuatkan variabel yang nantinya akan mengambil dari struct UserResponse " userResponse := midels.UserResponse(user.ID, user.Name, user.Email, token) "
15. lakukan testing di postman
16.jika sudah berhasil, buatlah batasan akses dibagian route. eJwt.Use(mid.JWT([]byte(consntan.SECRETE_JWT)))"



note : jika menggunakan echo jwt pada sebuah program dibutuhkan library tambahan yaitu dapat dicek pada dokumentasi " jwt golang " 
untuk mendapatkan library tambahan tersebut lakukan go get pada terminal " go get github.com/dgrijalva/jwt-go "
(harus nya udah default pada saat go get echo)

note : saat penambahan data pada file route.go jangan gunakan group untuk mengerjakan praktik. jadi dibagian gruop dikosingin ("/")
cukup taro /jwt di yang make aja. 

note : middleware tidak bisa digunakan untuk echo versi 4.

















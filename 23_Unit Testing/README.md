MATERI UNIT TESTING
#############################

UNIT TESTING
*******************
unit testing
=> suatu proses untuk menganalisa suatu software / sistem untuk menditeksi fitur saat ini dan fitur yang diinginkan oleh user.

tujuan 
=> untuk mengetahui apakah sistem sudah berjalan dengan baik tanpa muncul bug atau kerusakan sistem.
1. preventing regression
=> kondisi dimana sistem sudah berjalan dgn baik, namun jadi salah karna ada suatu perubahan. (penambahan fitur), maka dri itu lakukan testing ulang agar mengetahui 
bagianmana yang error atau rusak.
2. confidence level in refactoring
=> suatu proses dimana kita mengubah sistem tanpa mengubah fungsionalitas.
3. improve code design
=> memikirkan bagaimana bagian yang dirubah dapat ditesting agar berjalan dengan baik
4. code documentation
=> berisikan apakah sistem sudah berjalan dgn baik dlam segi output.
5. scalling the team
=> dapat dipahami oleh tim lainnya. 


level testing
=> 
1. UI
=> test yang dilakukan melalui user interface aplikasi. 
2. INTEGRATION
=> test terhadap spesifik modul atau subsistem API.
3. UNIT
=> testing yang dilakukan terhadap unit terkecil dalam sebuah sistem. testing terhadapt func atau method.


FRAMEWORK
*********************
framework
=> sebuah perangkatt kerja atau library dalam melakukan unit testing, sehingga tdk perlu membuat library untuk melakukan testing.

kenapa butuh ?
=> krna dpat lebih fokus dalam menghasilkan testing yang baik dan efesien tanpa fokus membangun library tersebut. 

list framework
=> 
1. testify (GO)
2. mocha dan jest (JS)


STRUKTUR UNIT TESTING
*************************
struktur 
=> strategi untuk diimplementasikan dalam program. cara untuk meletakkan file test dalam aplikasi code test.

pattren yang sering digunakan :
=> 
1. cetralize
=> membuat atau meletakan semua folder test kedalam satu folder yang terfokus dan biasa nay dibuat folder nama nya test.
2. save test file together
=> meletakan file testing di samping file yang akan di lakukan testing.

kelebihan dari kedua pttren diatas :
=> 
1. direktori dari aplikasi lebih clean krna menaruh semua file kedalam satu folder
2. lebih mudah dalam mencari file test dari function atau komponen yang akan ditest, krna file test diletakan berdampingan aatu disampingnya dengn file yang di test tersebut.


komponen dari unit testing :
=>
1. test file
=> berisikan code test dibuat dan didalamnya terdapat kumpulan test suites.
2. test suites
=> koleksi dari test cases.
3. test fixtures
=> suatu proses untuk memastikan apa yang kita pakai dalam melakukan testing itu konsisten. sehingga ketika melakukan test ulang maka hasil nya akan sama dengan test sebelumnya.
4. test cases
=> kumpulan kondisi yang dijalankan pda saat testing yang berisikan input, proses dan output.



RUNNER
**************
runner
=> software yang dirancang untuk menjalankan testing. 


MOCKING
*************
mocking
=> suatu yang dibuat untuk sebuah tiruan dalam hal testing. 


hal tdk disarankan untuk dilakukan testing
********************************************
=>
1. party modules
2. databse
3. party api or other extrenals system
4. object class


COVERAGE
**************
coverage
=> suatu alat ukur untuk menunjukan bagian mana yang sdah diexecute padaa saaat dilakukan testing.
coverage muncul pada saat testing sudah berjalan kemudian dicatat. 
bagian yang dianlisa :
=>
1. function
2. statement
3. branch
4. lines

hasil/report format dari testing coverage :
=>
1. cli
2. xml
3. html
4. lcov


SONARQUBE
*****************
sonarqube
=> 
tools khusus untuk menganalisa code yang dibuat atau code yang ada disistem. selain code juga dapat menganalisa bug maupun security.

 

IMPLEMENTASI UNIT TESTING
****************************************
1. buatlah file testing dengan format GO. 
dengan kriteria : 
	1.diharuskan nama library yang akan di test dan ditambahkan dengan _test.go (ex : user_test.go)
	2. lokasi file di package atau folder.
2. jika sudah buat file dan func. maka tulisakan satu function untuk melakukan test terhadap func yang lain. 
dengan syarat :
	1. nama func diawali dengan test dan huruf besar
	2. harus mengimplementasikan func Test(nama func yang mau ditest) kemdian membuat paramter (t *testing.T)
3. jika sudah menuliskan test maka pahami apa yang dituliskan. 
4. jika sudah semua lakukan running atau test dengan cara go test .lib/calculate -cover (go test (nama folder package) -cover).
5. jika ingin mengimplementasikan testing dan melihat report dengan cara go test .(nama folder package) coverprofile=cover.out && go tool cover -  html=cover.out.


TESTING SCENARIO
*************************
test scenario
=> testing mengenia gambaran umum mengani apa yang akan ditest. bisa disebut high level test.

contoh : 
pertama, saat craate new acctount di cek fungsionalitasnya. 
kedua, pada saat melakukan cek fungsionalitas login pada sebuah website.


TEST CASE
**************
test case 
=> kumpulan langkah-langkah dari sisi positif atau negatif dari sebuah test skenario.

contoh :
terdapat scenario mengenai cek fungsionalitas dari login dan terdapaat 4 test case
1. memasukan data email dan password dimana masing-masing field harus valid.
2. memasukan data email dan password dimana untuk field email itu valid dan password tdk valid.
3. memasukan data email dan password dimana untuk field email tdk valid dan password valid
4. memasukan data email dan password dimana untuk kedua feild itu tidak valid.

hal diatas merupaakn kumpulan test case. 

perbedaan kedua nya
********************

test scenario 
=> membicarakan apa yang aakn dites missal login / register

test case
=> membicarakan bagaimaana cara mengeetesnya dimana berisi kumpulan test untuk melakukan testing. 









	



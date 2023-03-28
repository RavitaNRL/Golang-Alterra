MATERI CLEAN AND HEXAGONAL ARCHITECTURE
********************************************
clean architecture
=> 

arsitektur yang baik itu terpisahkan menjadi beberapa layer dan dapat diperburui, diperbaiki serta dapat di testing.

MVVM
=> dimana view dengan model salalu terhubung.

hexagonal architecture dibagi 3
*********************************
1. API interface
=> berisikan controller, framework, cli.
2. domain service
=> penghubung antara kedua nya.
3. SPI interface
=> beriskan databse nya (sql, mongodb, redist. terparty).

3 layer arsitektur 
***********************
1. entitas layer
=> berisi bisnis objek.
2. use case / domain layer
=> berisi logika dari objek atau semua proses berada disana
3. controller / presention layer
=> berisikan framework nya atau tempat mengimplementasikan framework.
4. drivers / data layer
=> mengatur aplikasi data, seperti menerima data dri jaringan, membaca database, dll.








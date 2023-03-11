database
=> kumpulan data yang disimpan secara sistematis dalam komputer yang dapat diolah atau di manipulasi menggunakan perangkat lunak sehingga menghasilkan informasi.

relasi pada database dibagi menjadi 3 jenis, antara lain :

1. one to one
=> contoh, tiap penduduk harus memiliki 1 KTP; tiap KTP harus dimiliki oleh 1 penduduk.

2. one to many
=> contoh, satu fakultas dapat mengatur banyak jurusan; banyak jurusan hanya dapat diatur oleh 1 fakultas.

3. many to many
=> contoh, banyak supplier menyediakan banyak gudang; banyak gudang disediakan oleh banyak supplier. 


pada database terdapat istilah RDMS. RDMS singkatan dari Relational Database Model yang merupakan software untuk memanipulasi suatu data. contohnya MySql.
pada MySql menggunakan query SQL. 
jenis perintah pada sql yaitu :

1. DDL (data definition language)
=> perintah yang ada pada DDL, yaitu :

- Create (membuat)
- Drop (menghapus / delete)
- Rename (mengganti)
- Alter (mengubah)

2. DML (data manipulation language)
=> perintah yang ada pada DML, yaitu :

- Inser (memasukan/input)
- Select (menampilkan)
- Update (mengupdate)
- Delete (menghapus)
- Like / Between
- And /  Or
- Order By (descending/ascending)
- limit

pada query sql terdapat perintah join yaitu untuk menggabungkan. hal tersebut dibagi menjadi beberapa, sebagai berikut :

1. inner join
=> untuk mengembalikan baris dari dua tabel atau lebih.
query :
SELECT t.message FROM users u INNER JOIN tweets t ON u.id = t.user_id;

2. left join
=> untuk mengembalikan seluruh baris tabel dari sisi kiri yang dikenai kondisi ON dan hanya baris dari tabel sisi kanan yang memenuhi kondisi join.
query :
SELECT u.username, t.message FROM users u LEFT JOIN tweets t ON u.id = t.user_id;

3. right join
=> kebalikan dari left join, artinya untuk mengembalikan seluruh baris tabel dari sisi kanan yang dikenai kondisi ON dan hanya baris dari tabel sisi kiri yang memenuhi kondisi join.
query :
SELECT u.username, t.message FROM users u RIGHT JOIN tweets t ON u.id = t. user_id;









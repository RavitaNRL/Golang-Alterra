perintah join pada query sql dibagi menjadi 3 yaitu : 
********************************************************
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

perintah atau fungsi agregasi pada sql dibagi menjadi beberapa, sebagai berikut :
***********************************************************************************
1. MIN 
=> untuk mencari nilai terkecil dari suatu data di dalam tabel. 
perintah untuk menjalannya fungsi MIN yaitu 

SELECT MIN(id) AS id FROM users;
logika :
akan menampilkan data terkecil pada field id dari tabel users. 

2. MAX
=> untuk menentukan nilai terbesar dari suatu data pada sebuah tabel. 
perintah untuk menjalankan fungsi MAX yaitu :

SELECT MAX(created_at) FROM users;
logika :
akan menampilkan data pada field created_at yang terbesar dari tabel users.

3. SUM 
=> untuk menjumlahkan total nilai dari sebuah data di dalam tabel. 
perintah untuk menjalankan fungsi SUM yaitu :

SELECT SUM(favorit_num) FROM tweets WHERE user_id = 2;
logika :
akan menampilkan jumlah total pada field favorit_num dari tabel tweet dimana user_id = 2.

4. AVG 
=> untuk mencari rata-rata pada sebuah data di dalam tabel. 
perintah untuk menjalankan fungsi AVG yaitu :

SELECT AVG(nilai_uts) FROM ipk WHERE npm = 51420063;
logika :
akan menampilkan nilai rata-rata pada field nilai_uts dari tabel ipk dengan npm 51420063.

5. COUNT
=> untuk mencari jumlah dari sebuah data pada tabel. 
perintah untuk menjalankan fungsi COUNT yaitu : 

SELECT COUNT(2) FROM tweets WHERE id = 5;
logika :
akan menampilkan jumlah data dari tabel tweets dimana id nya adalah 5.

6. HAVING 
=> untuk menyelesaikan data berdasarkan kriteria tertentu. 
perintah untuk menjalankan fungsi HAVING yaitu : 

SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(favorit_count) >2;
logika :
akan menampilkan data dari tabel tweets dengan jumlah total field favorit_count per user lebih dari 2.

SUBQUERY
***********

subquery yaitu fungsi yang digunakan untuk mengembalikan data yang akan digunakan dalam query utama dengan syarat untuk membatasi data yang akan diambil. 
fungsi ini dapat dijalankan dengan perintah :
- SELECT
- INSERT
- UPDATE
- DELETE
dengan dikombinaskan dengan operator beetween, in, dst.


FUNCTION
*********

function yaitu kumpulan statemen yang akan mengembalikan sebuah nilai balik pada pemanggilnya.
contoh :

DELIMITER $$
CREATE FUNCTION sf_count_tweet_peruser
 (user_id_p int) RETURNS INT DETERMINISTIC
 BEGIN
DECLARE total INT;
SELECT COUNT (*) INTO total FROM tweets
WHERE user_id = user_id_p AND
type= 'tweets';
 RETURN total;
ENDS$
DELIMITER;





 

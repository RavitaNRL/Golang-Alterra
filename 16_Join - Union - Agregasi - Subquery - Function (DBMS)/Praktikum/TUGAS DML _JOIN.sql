-- MEMBUAT QUERY SQL DARI ERD YANG ADA DI TUGAS 1 (GAMBAR)
-------------------------------------------------------------

-- table product_types
-- pk id int(11)
--     name varchar(255)
--     create_at timestamp
--     update_at timestamp

CREATE TABLE product_types (
  id INT(11) PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);

-- table operators
-- pk id int(11)
--     name varchar(255)
--     create_at timestamp
--     update_at timestamp

CREATE TABLE operators (
  id INT(11) PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);


-- table products_description
-- pk id int(11)
--     description text
--     create_at timestamp
--     update_at timestamp

CREATE TABLE products_description (
  id INT(11) PRIMARY KEY,
  description TEXT,
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);

-- table payment_methods
-- pk id int(11)
--     name varchar(255)
--     status smallint
--     create_at timestamp
--     update_at timestamp

CREATE TABLE payment_methods (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);


-- table users
-- pk id int(11)
--     status smallint
--     dob date
--     gender char(1)
--     create_at timestamp
--     update_at timestamp

CREATE TABLE users (
    id INT(11) PRIMARY KEY,
    status SMALLINT,
    dob DATE,
    gender char(1),
    created_at TIMESTAMP NULL DEFAULT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL
);


-- table transactions
--  pk id int(11)
--  fk user_id int(11)
--  fk payment_methods int(11)
--  status varchar(10)
--  total_qty int(11)
--  total_price numberic(25,2)
--  create_at timestamp
--  update_at timestamp

CREATE TABLE transactions (
  id INT(11) PRIMARY KEY,
  user_id INT(11),
  payment_methods INT(11),
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25,2),
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);

-- membuat foreign key pada field user_id ditebel transactions

ALTER TABLE transactions ADD FOREIGN KEY (user_id) REFERENCES users(id);

-- membuat foreign key pada field payment_methods ditebel transactions

ALTER TABLE transactions ADD FOREIGN KEY (payment_methods) REFERENCES payment_methods(id);

-- transactions_details
--  pk transactions_id int(11)
--  pk product_id int(11)
--  status varchar(10)
--  qty int(11)
--  price numberic(25,2)
--  create_at timestamp
--  update_at timestamp

CREATE TABLE transactions_details (
  transactions_id INT(11) PRIMARY KEY,
  product_id INT(11),
  status VARCHAR(10),
  qty INT(11),
  price NUMERIC(25,2),
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);

-- membuat foreign key pada field product_id di tebel transactions_details.

ALTER TABLE transactions_details ADD FOREIGN KEY (product_id) REFERENCES products(id);

-- table products
--  pk id int(11)
--  fk product_types_id int(11)
--  fk operators_id int(11)
--  code varchar(255)
--  name varchar(255)
--  status smallint
--  create_at timestamp
--  update_at timestamp

CREATE TABLE products (
  id INT(11) PRIMARY KEY,
  product_types_id INT(11),
  operators_id INT(11),
  code VARCHAR(255),
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP NULL DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL
);

-- membuat foreign key pada field product_types_id di tebel products.

ALTER TABLE products ADD FOREIGN KEY (product_types_id) REFERENCES product_types(id);

-- membuat foreign key pada field operators_id di tebel products.

ALTER TABLE products ADD FOREIGN KEY (operators_id) REFERENCES operators(id);

----------------------------------------------------------------------------------------

---------------------------------------------------
-- TUGAS DML 
----------------------------------------------------

-- 1. INSERT 
--      a. Insert 5 operators pada table operators.

INSERT INTO operators (id, name, created_at, updated_at) VALUES 
(1, 'SARAH', '2023-03-01 00:05:07', '2023-05-23 09:08:01'),
(2, 'DEWI', '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(3, 'MEE', '2021-03-01 00:05:07', '2023-05-23 09:08:01'),
(4, 'JENY', '2022-03-01 00:05:07', '2023-10-23 09:08:01'),
(5, 'Smartfren', '2021-03-01 00:05:07', '2022-05-23 09:08:01');

--      b. Insert 3 product type.

INSERT INTO product_types (id, name, created_at, updated_at) VALUES 
(1, 'ELECTRONIC', '2021-01-20 18:05:17', '2022-05-23 09:08:01'),
(2, 'FASHION', '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(3, 'FOOD', '2021-03-01 00:05:07', '2023-05-23 09:08:01');

--      c. Insert 2 product dengan product type id = 1, dan operators id = 3.

INSERT INTO products (id, product_types_id, operators_id, code, name, status, created_at, updated_at) VALUES 
(101, 1, 3, 'NEW', 'KULKAS', 4, '2023-03-01 00:05:07', '2023-03-10 00:05:07'),
(102, 1, 3, 'OLD', 'MESIN CUCI', 1, '2023-01-20 07:05:07', '2023-03-10 12:05:07');

--      d. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (id, product_types_id, operators_id, code, name, status, created_at, updated_at) VALUES 
(103, 2, 1, 'NEW', 'TV', 4, '2020-09-08 14:00:07', '2022-03-10 07:05:18'),
(104, 2, 1, 'OLD', 'AC', 1, '2023-01-20 07:05:07', '2023-10-10 12:05:07'),
(105, 2, 1, 'OLD', 'KIPAS ANGIN', 1, '2021-04-20 11:00:18', '2022-01-29 15:15:00');

--      e. Insert 3 product dengan product type id = 3, dan operators id = 4.

INSERT INTO products (id, product_types_id, operators_id, code, name, status, created_at, updated_at) VALUES 
(106, 3, 4, 'NEW', 'LAPTOP', 4, '2021-10-24 09:00:37', '2022-09-10 11:01:07'),
(107, 3, 4, 'OLD', 'HANDPHONE', 1, '2023-01-20 07:05:07', '2024-01-06 15:00:27'),
(108, 3, 4, 'OLD', 'MEJA', 1, '2023-01-20 07:05:07', '2023-03-10 12:05:07'); 

--      f. Insert product description pada setiap product.

INSERT INTO products_description (id, description, created_at, updated_at) VALUES 
(101, 'KULKAS BARU', '2023-03-01 00:05:07', '2023-03-10 00:05:07'),
(102, 'MESIN CUCI BARU', '2023-01-20 07:05:07', '2023-03-10 12:05:07'),
(103, 'TV BARU', '2020-09-08 14:00:07', '2022-03-10 07:05:18'),
(104, 'AC BARU', '2023-01-20 07:05:07', '2023-10-10 12:05:07'),
(105, 'KIPAS ANGIN BARU', '2021-04-20 11:00:18', '2022-01-29 15:15:00'),
(106, 'LAPTOP BARU', '2021-10-24 09:00:37', '2022-09-10 11:01:07'),
(107, 'HANDPHONE BARU', '2023-01-20 07:05:07', '2024-01-06 15:00:27'),
(108, 'MEJA BARU', '2023-01-20 07:05:07', '2023-03-10 12:05:07');

--      g. Insert 3 payment methods

INSERT INTO payment_methods (id, name, status, created_at, updated_at) VALUES 
(1, 'grey', '1', '2021-01-20 18:05:17', '2022-05-23 09:08:01'),
(2, 'joy', '2', '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(3, 'jems', '4', '2021-03-01 00:05:07', '2023-05-23 09:08:01');

--      h. Insert 5 user pada tabel user.

INSERT INTO users (id, status, dob, gender, created_at, updated_at) VALUES 
(001, 6, '2021-01-20', 'P', '2021-01-20 18:05:17', '2022-05-23 09:08:01'),
(002, 5, '2022-03-01', 'L', '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(003, 1, '2021-03-01', 'P', '2021-03-01 00:05:07', '2023-05-23 09:08:01'),
(004, 7, '2021-03-01', 'L', '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(005, 5, '2021-03-01', 'P', '2021-03-01 00:05:07', '2023-05-23 09:08:01');

--      i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1).

INSERT INTO transactions (id, users_id, payment_methods_id, status, total_qty, total_price, created_at, updated_at) VALUES 
(1, 001, 1, '1', 5, 1000, '2021-01-20 18:05:17', '2022-05-23 09:08:01'),
(2, 002, 2, '2', 6, 2000, '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(3, 003, 3, '3', 2, 3000, '2021-03-01 00:05:07', '2023-05-23 09:08:01');

-- masih error pas diimplementasiin ke mysql!!!!
--------------------------------------------------

--      i. Insert 3 product di masing-masing transaksi.

INSERT INTO transactions_details (id, transactions_id, products_id, qty, price, created_at, updated_at) VALUES 
(1, 1, 101, 1, 1000, '2021-01-20 18:05:17', '2022-05-23 09:08:01'),
(2, 2, 102, 2, 2000, '2022-03-01 00:05:07', '2023-05-24 07:08:01'),
(3, 3, 103, 3, 3000, '2021-03-01 00:05:07', '2023-05-23 09:08:01');

INSERT INTO product(id, product_types_id, operators_id, code, name, status, created_at, updated_at) VALUES 
(1, 1, 1, 'NEW', 'KULKAS', 1, '2023-03-01 00:05:07', '2023-03-10 00:05:07'),
(2, 1, 1, 'OLD', 'MESIN CUCI', 1, '2023-01-20 07:05:07', '2023-03-10 12:05:07'),
(3, 1, 1, 'OLD', 'TV', 1, '2020-09-08 14:00:07', '2022-03-10 07:05:18');


-- 2. SELECT 
--      a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.

SELECT nama FROM user WHERE gender = 'L' or gender = 'M';

--      b. Tampilkan product dengan id = 3.

SELECT * FROM product WHERE id = 3;

--      c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.

SELECT * FROM user WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND nama LIKE '%a%';

--      d. Hitung jumlah user / pelanggan dengan status gender Perempuan.

SELECT COUNT(*) FROM user WHERE gender = 'P';

--      e. Tampilkan data pelanggan dengan urutan sesuai nama abjad

SELECT * FROM user ORDER BY nama ASC;

--      f. Tampilkan 5 data pada data product

SELECT * FROM product LIMIT 5;


-- 3. UPDATE
--      a. Ubah data product id 1 dengan nama ‘product dummy’.

UPDATE products SET name = 'product dummy' WHERE id = 1;

--      b. Update qty = 3 pada transaction detail dengan product id = 1.

UPDATE transaction_detail SET qty = 3 WHERE product_id = 1;


-- 4. DELETE
--      a. Delete data pada tabel product dengan id = 1.

DELETE FROM products WHERE id = 1;

--     b. Delete pada pada tabel product dengan product type id 1.

DELETE FROM products WHERE product_types_id = 1;

-------------------------------------------------------------------------------------

---------------------------------------------
-- TUGAS JOIN
----------------------------------------------

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.

-- SELECT * FROM transactions WHERE users_id = 1 OR users_id = 2;


-- 2. Tampilkan jumlah harga transaksi user id 1.

SELECT SUM(total_price) FROM transactions WHERE users_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2.

SELECT COUNT(*) FROM transactions_details WHERE product_types_id = 2;

-- 4. Tampilkan semua field table product dan field name table products_type yang saling berhubungan.

SELECT product.*, product_type.name 
FROM product 
INNER JOIN product_type ON product.product_type_id = product_type.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT transactions.*, product.name, user.name 
FROM transactions
INNER JOIN product ON transactions.product_id = product.id
INNER JOIN user ON transactions.user_id = user.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

DELETE FROM transactions  
WHERE transaction_id = <transaction_id> 
AND DELETE FROM transaction_details 
WHERE transaction_id = <transaction_id> 

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

CREATE FUNCTION UpdateTotalQty() 
RETURNS VOID 
AS
BEGIN 
  UPDATE Total_Qty SET qty = 
  (SELECT SUM(qty) 
  FROM TransactionDetails
  WHERE transaction_id NOT IN (SELECT transaction_id FROM DeletedTransaction))
END;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

SELECT * FROM products WHERE id NOT IN (SELECT products_id FROM transactions_details);














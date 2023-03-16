-------------------------
-- PERINTAH INSERT
-------------------------

-- Insert 5 operators pada table operators.

INSERT INTO operator(name) VALUES
('TOKO CERIA'),
('TOKO SEJAHTERA'),
('TOKO ALI'),
('TOKO MERPATI'),
('TOKO MAJU JAYA');

-- Insert 3 product type.

INSERT INTO product_types(name) VALUES
('Alat Tulis'),
('ELEKTRONIK'),
('MAKANAN');

-- Insert 2 product dengan product type id = 1, dan operators id = 3.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(1, 3, 'MKN001', 'Buku Sidu', 2),
(1, 3, 'MKN002', 'Crayon', 1);

-- Insert 3 product dengan product type id = 2, dan operators id = 1.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(2, 1, 'MNM001', 'Lampu Jalan', 4),
(2, 1, 'MNM002', 'Televisi', 3),
(2, 1, 'MNM003', 'Kabel', 1);

-- Insert 3 product dengan product type id = 3, dan operators id = 4.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(3, 4, 'MNM004', 'Biskuit', 2),
(3, 4, 'MNM005', 'Mie Instan', 5),
(3, 4, 'MNM006', 'Chocolate', 1);

-- Insert product description pada setiap product.

INSERT INTO product_descriptions(description) VALUES
('buku sidu untuk menulis'),
('crayon untuk mewarnai gambar'),
('lampu jalan untuk menerangi jalan'),
('televisi untuk menonton'),
('kabel untuk menghubungkan'),
('biskuit cemilan ringan'),
('mie instan untuk makan'),
('cemilan bersalut chocolate');

-- Insert 3 payment methods.

INSERT INTO payment_methods(name, status) VALUES
('BCA', 1),
('Mandiri', 2),
('Gopay', 3);

-- Insert 5 user pada tabel user.

INSERT INTO users(name, status, dob, gender) VALUES
('Budi', 1, '1990-01-01', 'L'),
('Susi', 2, '1991-02-02', 'P'),
('Rudi', 3, '1992-03-03', 'L'),
('Sari', 4, '1993-04-04', 'P'),
('Rina', 5, '1994-05-05', 'P');

-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

INSERT INTO transaksi(user_id, payments_method_id, status, total_qty, total_price) VALUES
(1, 1, 'Cashless', 2, 20000),
(2, 2, 'Cash', 3, 30000),
(3, 3, 'Credit', 4, 40000);

-- Insert 3 product di masing-masing transaksi.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(1, 1, 'MKN001', 'Buku Sidu', 2),
(2, 2, 'MKN002', 'Crayon', 1),
(3, 3, 'MNM001', 'Lampu Jalan', 4);

-------------------------
-- PERINTAH SELECT
-------------------------

-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.

SELECT name FROM users WHERE gender = 'L' or gender = 'M';

-- Tampilkan product dengan id = 3.

SELECT * FROM product WHERE id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.

SELECT * FROM users WHERE create_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan.

SELECT COUNT(*) FROM users WHERE gender = 'P';

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad

SELECT * FROM users ORDER BY name ASC;

-- Tampilkan 5 data pada data product

SELECT * FROM product LIMIT 5;

-------------------------
-- PERINTAH UPDATE
-------------------------

-- Ubah data product id 1 dengan nama ‘product dummy’.

UPDATE product SET name = 'product dummy' WHERE id = 1;

-- Update qty = 3 pada transaction detail dengan product id = 1.

UPDATE transaksi_details SET qty = 3 WHERE product_id = 1;

-------------------------
-- PERINTAH DELETE
-------------------------

-- Delete data pada tabel product dengan id = 1.

DELETE FROM product WHERE id = 1;

-- Delete pada pada tabel product dengan product type id 1.

DELETE FROM product WHERE product_type_id = 1;




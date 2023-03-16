-- Gabungkan data transaksi dari user id 1 dan user id 2.

SELECT * FROM transaksi WHERE user_id = 1 OR user_id = 2;

-- Tampilkan jumlah harga transaksi user id 1.

SELECT SUM(total_price) FROM transaksi WHERE user_id = 1;

-- Tampilkan semua field table product dan field name table products_type yang saling berhubungan.

SELECT product.*, product_types.name FROM product JOIN product_types ON product.product_type_id = product_types.id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT transaksi.*, product.name, users.name 
FROM transactions
INNER JOIN product ON transaksi.product_id = product.id
INNER JOIN user ON transaksi.user_id = user.id;



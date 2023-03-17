-- Gabungkan data transaksi dari user id 1 dan user id 2.

SELECT * FROM transaksi WHERE user_id = 1 OR user_id = 2;

-- Tampilkan jumlah harga transaksi user id 1.

SELECT SUM(total_price) FROM transaksi WHERE user_id = 1;

-- Tampilkan total transaksi dengan product type 2.

SELECT SUM(total_price) as total_price FROM transaksi WHERE user_id = 1 group by user_id;

-- Tampilkan semua field table product dan field name table products_type yang saling berhubungan.

SELECT product.*, product_types.name FROM product 
JOIN product_types ON product.product_type_id = product_types.id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT t.*, p.name, u.name as users FROM transaksi_details td 
JOIN transaksi t ON t.id = td.transaksi_id
JOIN users u ON u.id = t.user_id
JOIN product p ON p.id = td.product_id
WHERE t.user_id = 1;

 -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

delimiter $$
CREATE trigger delete_transaksi_details
BEFORE DELETE ON transaksi for each row
BEGIN
    DECLARE trans_id INT;
    set trans_id = old.id;
    DELETE FROM transaksi_details WHERE transaksi_id = trans_id;
END $$

DELETE FROM transaksi WHERE id = 4; 

SELECT * FROM transaksi_details;

-- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

delimiter $$
CREATE trigger update_transaksi_id 
AFTER DELETE ON transaksi_details for each row
BEGIN
    DECLARE trans_id INT;
    set trans_id = old.transaksi_id;
    UPDATE transaksi set total_qty = (SELECT  SUM (qty) FROM transaksi_details WHERE transaksi_id = trans_id) 
    WHERE id = trans_id;
END $$ 

SELECT * FROM transaksi_details;

SELECT * FROM transaksi_details WHERE transaksi_id = 2;

DELETE FROM transaksi_details WHERE transaksi_id = 2 AND product_id = 9;

-- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

SELECT * FROM product WHERE id NOT IN (SELECT product_id FROM transaksi_details);

SELECT * FROM product;




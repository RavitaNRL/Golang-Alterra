-- membuat tabel customer
CREATE TABLE `customer` (
  `customer_id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `customer_nama` varchar(20) NOT NULL,
  `gender` enum('Laki-Laki','Perempuan') NOT NULL,
  `status_user` enum('baru','lama') NOT NULL,
  `address` varchar(30) NOT NULL,
  `create_data` date NOT NULL,
  `update_data` date NOT NULL
);

-- membuat tabel pemesanan
CREATE TABLE `pemesanan` (
  `no_pemesanan` int(11) PRIMARY KEY AUTO_INCREMENT,
  `product_name` varchar(25) NOT NULL,
  `customer_name` varchar(20) NOT NULL,
  `address` varchar(30) NOT NULL,
  `tangga_pemesanan` date NOT NULL,
  `tanggal_pengiriman` date NOT NULL
);

-- membuat tabel product
CREATE TABLE `product` (
  `product_id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `product_name` varchar(20) NOT NULL,
  `operator` varchar(20) NOT NULL,
  `product_type` varchar(20) NOT NULL,
  `payment_method` enum('cash','cashless') NOT NULL,
  `product_deskripsi` varchar(25) NOT NULL
);

-- membuat tabel transaksi
CREATE TABLE `transaksi` (
  `id_transaksi` int(11) PRIMARY KEY AUTO_INCREMENT,
  `customer_name` varchar(20) NOT NULL,
  `product_id` int(11) NOT NULL,
  `product_price` int(11) NOT NULL,
  `many_product` int(11) NOT NULL,
  `payment_methods` enum('cash','cashless') NOT NULL
); 

-- membuat foreign key untuk fielf product_id pada tabel transaksi
ALTER TABLE transaksi ADD FOREIGN KEY (product_id) REFERENCES product(product_id);

-- membuat tabel transaksi details
CREATE TABLE `transaksi_details` (
  `transaksi_details_id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `no_pemesanan` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `customer_name` varchar(20) NOT NULL,
  `product_type` varchar(20) NOT NULL,
  `many_product` int(11) NOT NULL,
  `payment_date` date NOT NULL,
  `payment_methods` enum('cash','cashless') NOT NULL,
  `tanggal_pemesanan` date NOT NULL,
  `tanggal_pengiriman` date NOT NULL
);

-- membuat foreign key untuk field no_pemesanan pada tabel transaksi_details
ALTER TABLE transaksi_details ADD FOREIGN KEY (no_pemesanan) REFERENCES pemesanan(no_pemesanan);

-- membuat foreign key untuk field customer_id pada tabel transaksi_details
ALTER TABLE transaksi_details ADD FOREIGN KEY (customer_id) REFERENCES customer(customer_id);

-- membuat tabel kurir
CREATE TABLE `kurir` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `nama` varchar(20) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL
);

-- menambahkan field ongkos_dasar pada tabel kurir
ALTER TABLE kurir ADD ongkos_dasar int(11) NOT NULL;

-- merubah nama tabel kurir menjadi shipping
ALTER TABLE kurir RENAME TO shipping;

-- menghapus tabel shipping karena tidak terpakai
DROP TABLE shipping;

CREATE DATABASE AltaOnline_Shop;

-- TABEL product_types

CREATE TABLE product_types (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(20),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABEL operator

CREATE TABLE operator (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABEL product_descriptions

CREATE TABLE product_descriptions (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(255),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABEL payment_methods

CREATE TABLE payment_methods (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    status SMALLINT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABEL users

CREATE TABLE users (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    status VARCHAR(20),
    dob DATE,
    gender CHAR(1),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABEL payment_methods_details

CREATE TABLE payment_methods_details (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    payment_method_id INT(11),
    detail VARCHAR(50),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- TABEL alamat

CREATE TABLE alamat (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    user_id INT(11),
    nama_jalan VARCHAR(30),
    kode_post INT(11),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- TABEL kurir

CREATE TABLE kurir (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    alamat_id INT(11),
    nama_penerima VARCHAR(30),
    no_tlpn INT(11),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (alamat_id) REFERENCES alamat(id)
);

-- TABEL user_payment_methods_details

CREATE TABLE user_payment_methods_details (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    user_id INT(11),
    payment_method_detail_id INT(11),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_detail_id) REFERENCES payment_methods_details(id)
);

-- TABEL transaksi

CREATE TABLE transaksi (
	id INT(11) PRIMARY KEY AUTO_INCREMENT,
    user_id INT(11),
	payments_method_details_id INT,
	transaksi_details_id INT,
	status VARCHAR(10),
	total_qty INT,
	total_price INT,
	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(payments_method_details_id) REFERENCES payment_methods_details(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- membuat foreign key untuk field transaksi_details_id pada tabel transaksi

ALTER TABLE transaksi ADD FOREIGN KEY (transaksi_details_id) REFERENCES transaksi_details(id);

-- TABLE transaksi_details

CREATE TABLE transaksi_details (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    transaksi_id INT(11),
    product_id INT(11),
    status VARCHAR(10),
    qty INT,
    price NUMERIC(25,2),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(transaksi_id) REFERENCES transaksi(id)
);

-- membuat foreign key pada tabel transaksi_details untuk kolom product_id

ALTER TABLE transaksi_details ADD FOREIGN KEY (product_id) REFERENCES product(id);

-- TABEL product

CREATE TABLE product (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    product_type_id INT(11),
    operator_id INT(11),
    transaction_details_id INT(11),
    name VARCHAR(255),
    price NUMERIC(25,2),
    stock INT(11),
    status VARCHAR(255),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (transaction_details_id) REFERENCES transaksi_details(id),
    FOREIGN KEY (operator_id) REFERENCES operator(id)
);


-- menambahkan field ongkos_dasar pada tabel kurir

ALTER TABLE kurir ADD ongkos_dasar int(11) NOT NULL;

-- merubah nama tabel kurir menjadi shipping

ALTER TABLE kurir RENAME TO shipping;

-- menghapus tabel shipping karena tidak terpakai

DROP TABLE shipping;


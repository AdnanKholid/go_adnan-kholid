--nomer 2.1
CREATE DATABASE alta_online_shop;
CREATE DATABASE alta;


--nomer 2.2.a
CREATE TABLE user (
    id INT NOT NULL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    alamat TEXT NOT NULL,
    tanggal_lahir DATE NOT NULL,
    status_user int NOT NULL,
    gender char (1) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL);

--nomer 2.2.b
CREATE TABLE Product (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name_product VARCHAR(255) NOT NULL,
    product_type_id INT NOT NULL,
    operator_id INT NOT NULL,
    FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    FOREIGN KEY (operator_id) REFERENCES operator(id));

CREATE TABLE product_type (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name_product_type VARCHAR(255) NOT NULL);

CREATE TABLE operators (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name_operator VARCHAR(255) NOT NULL);

CREATE TABLE product_description (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    name_product_desciption VARCHAR(255) NOT NULL,
    price int NOT NULL,
    description TEXT,
    FOREIGN KEY (product_id) REFERENCES product(id));

CREATE TABLE payment_method(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name_payment_method VARCHAR(255) NOT NULL);

--nomer 2.2.c
CREATE TABLE transaction(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    payment_method_id int NOT NULL,
    tanggal_transaksi DATETIME NOT NULL,
    sub_total INT NOT NULL,
    total_item INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES payment_method(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id));

CREATE TABLE transaction detail(
    id INT PRIMARY KEY AUTO_INCREMENT,
    transaksi_id INT NOT NULL,
    product_id INT NOT NULL,
    total_product INT NOT null,
    FOREIGN KEY (transaksi_id) REFERENCES transaksi(id),
    FOREIGN KEY (product_id) REFERENCES product(id));

--nomer 2.3
CREATE TABLE kurir(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL);

--nomer 2.4
ALTER TABLE kurir
    ADD ongkos_dasar INT NOT NULL;


--nomer 2.5
ALTER TABLE kurir
    RENAME TO shipping;


--nomer 2.6
DROP TABLE shipping;


--nomer 2.7.a
CREATE TABLE payment_method_description (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nama_perusahaan VARCHAR(255) NOT NULL,
    payment_method_id int NOT NULL,
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id));

--nomer 2.7.b
CREATE TABLE alamat(
    id INT PRIMARY KEY AUTO_INCREMENT,
    alamat TEXT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id));

--nomer 2.7.c(
CREATE TABLE user_payment_method_id detail ( 
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    payment_method_id int NOT NULL,
    tanggal_payment DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id));



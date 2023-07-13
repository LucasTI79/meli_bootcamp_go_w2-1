CREATE DATABASE IF NOT EXISTS melisprint;
USE melisprint;

DROP 
  TABLE IF EXISTS products;
CREATE TABLE products(
  `id` INT PRIMARY KEY AUTO_INCREMENT, 
  `description` TEXT NOT NULL, expiration_rate FLOAT NOT NULL, 
  freezing_rate FLOAT NOT NULL, height FLOAT NOT NULL, 
  lenght FLOAT NOT NULL, netweight FLOAT NOT NULL, 
  product_code TEXT NOT NULL, recommended_freezing_temperature FLOAT NOT NULL, 
  width FLOAT NOT NULL, id_product_type INT NOT NULL, 
  id_seller INT NOT NULL
);

DROP 
  TABLE IF EXISTS employees;
CREATE TABLE employees(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  card_number_id TEXT NOT NULL, first_name TEXT NOT NULL, 
  last_name TEXT NOT NULL, warehouse_id INT NOT NULL
);

DROP 
  TABLE IF EXISTS countries;
CREATE TABLE countries(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  country_name TEXT NOT NULL
);

DROP 
  TABLE IF EXISTS provinces;
CREATE TABLE provinces(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  province_name TEXT NOT NULL, 
  country_id INT, 
  FOREIGN KEY(country_id) REFERENCES melisprint.countries(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS localities;
CREATE TABLE localities(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  locality_name TEXT NOT NULL, 
  province_id INT, 
  FOREIGN KEY(province_id) REFERENCES melisprint.provinces(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS warehouses;
CREATE TABLE warehouses(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `address` VARCHAR(255) NULL,
  telephone VARCHAR(255) NULL,
  warehouse_code VARCHAR(255) NULL UNIQUE,
  minimum_capacity INT NULL,
  minimum_temperature INT NULL,
  locality_id INT NOT NULL,
  FOREIGN KEY(locality_id) REFERENCES `melisprint`.`localities` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS sections;
CREATE TABLE sections(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  section_number INT NOT NULL, current_temperature INT NOT NULL, 
  minimum_temperature INT NOT NULL, 
  current_capacity INT NOT NULL, minimum_capacity INT NOT NULL, 
  maximum_capacity INT NOT NULL, warehouse_id INT NOT NULL, 
  id_product_type INT NOT NULL
);

DROP 
  TABLE IF EXISTS sellers;
CREATE TABLE sellers(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  cid INT NOT NULL UNIQUE, 
  company_name TEXT NOT NULL, 
  `address` TEXT NOT NULL, 
  telephone TEXT(15) NOT NULL,
  `locality_id` INT NOT NULL, 
  FOREIGN KEY (`locality_id`) REFERENCES `melisprint`.`localities` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS buyers;
CREATE TABLE buyers(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  card_number_id TEXT NOT NULL, first_name TEXT NOT NULL, 
  last_name TEXT NOT NULL
);

DROP 
  TABLE IF EXISTS product_types;
CREATE TABLE product_types (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  description TEXT NOT NULL
);

ALTER TABLE 
  products 
ADD 
  FOREIGN KEY(id_product_type) REFERENCES product_types(id) ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE 
  products 
ADD 
  FOREIGN KEY(id_seller) REFERENCES sellers(id) ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE 
  sections 
ADD 
  FOREIGN KEY(id_product_type) REFERENCES product_types(id) ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE 
  sections 
ADD 
  FOREIGN KEY(warehouse_id) REFERENCES warehouses(id) ON DELETE NO ACTION ON UPDATE NO ACTION;

DROP 
  TABLE IF EXISTS product_batches;
CREATE TABLE product_batches(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  `batch_number` INT NOT NULL, 
  `initial_quantity` INT NOT NULL, 
  `current_quantity` INT NOT NULL, 
  `current_temperature` DECIMAL(19, 2) NOT NULL, 
  `due_date` DATETIME NOT NULL, 
  `manufacturing_date` DATETIME NOT NULL, 
  `manufacturing_hour` INT NOT NULL, 
  `minimum_temperature` DECIMAL(19, 2) NOT NULL, 
  `product_id` INT NOT NULL, 
  `section_id` INT NOT NULL, 
  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(section_id) REFERENCES sections(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS product_records;
CREATE TABLE product_records(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  `last_update_date` DATETIME NOT NULL, 
  `purchase_price` DECIMAL(19, 2) NOT NULL, 
  `sale_price` DECIMAL(19, 2) NOT NULL, 
  `product_id` INT NOT NULL, 
  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS carriers;
CREATE TABLE carriers(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  `cid` VARCHAR(255) NOT NULL, 
  `company_name` VARCHAR(255) NOT NULL, 
  `address` VARCHAR(255) NOT NULL, 
  `telephone` VARCHAR(255) NOT NULL, 
  `locality_id` INT NOT NULL, 
  FOREIGN KEY(locality_id) REFERENCES localities(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS order_status;
CREATE TABLE order_status (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  description TEXT NOT NULL
);

DROP 
  TABLE IF EXISTS purchase_orders;
CREATE TABLE IF NOT EXISTS melisprint.purchase_orders (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  `order_number` VARCHAR(255) NOT NULL, 
  `order_date` DATETIME(6) NOT NULL, 
  `tracking_code` VARCHAR(255) NOT NULL, 
  `buyer_id` INT NOT NULL, 
  `carrier_id` INT NULL, 
  `order_status_id` INT NOT NULL, 
  `warehouse_id` INT NULL, 
  `product_record_id` INT NOT NULL, 
  FOREIGN KEY(buyer_id) REFERENCES buyers(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(carrier_id) REFERENCES carriers(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(order_status_id) REFERENCES order_status(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(warehouse_id) REFERENCES warehouses(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(product_record_id) REFERENCES product_records(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS order_details;
CREATE TABLE IF NOT EXISTS melisprint.order_details (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
  `clean_liness_status` TEXT NOT NULL, 
  `quantity` INT DEFAULT 1, 
  `temperature` DECIMAL(19, 2) NOT NULL, 
  `product_record_id` INT NOT NULL, 
  `purchase_order_id` INT NOT NULL, 
  FOREIGN KEY(purchase_order_id) REFERENCES purchase_orders(id) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  FOREIGN KEY(product_record_id) REFERENCES product_records(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

DROP 
  TABLE IF EXISTS inbound_orders;
CREATE TABLE IF NOT EXISTS `melisprint`.`inbound_orders` (
  `id` INT NOT NULL AUTO_INCREMENT, 
  `order_date` DATETIME(6) NOT NULL, 
  `order_number` VARCHAR(255) NOT NULL, 
  `employee_id` INT NOT NULL, 
  `product_batch_id` INT NOT NULL, 
  `warehouse_id` INT NOT NULL, 
  PRIMARY KEY (`id`), 
  INDEX `employee_id_idx` (`employee_id` ASC) VISIBLE, 
  INDEX `product_batch_id_idx` (`product_batch_id` ASC) VISIBLE, 
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE, 
  CONSTRAINT `fk_employee_inbound_orders` FOREIGN KEY (`employee_id`) REFERENCES `melisprint`.`employees` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  CONSTRAINT `fk_product_batch_inbound_orders` FOREIGN KEY (`product_batch_id`) REFERENCES `melisprint`.`product_batches` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION, 
  CONSTRAINT `fk_warehouse_inbound_orders` FOREIGN KEY (`warehouse_id`) REFERENCES `melisprint`.`warehouses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB;

DROP 
  TABLE IF EXISTS roles;
CREATE TABLE IF NOT EXISTS `melisprint`.`roles` (
  `id` INT NOT NULL AUTO_INCREMENT, 
  `description` VARCHAR(255) NOT NULL, 
  `rol_name` VARCHAR(255) NOT NULL, 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

DROP 
  TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS `melisprint`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT, 
  `passoword` VARCHAR(255) NOT NULL, 
  `username` VARCHAR(255) NOT NULL, 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

DROP 
  TABLE IF EXISTS user_rol;
CREATE TABLE IF NOT EXISTS `melisprint`.`user_rol` (
  `usuario_id` INT NOT NULL AUTO_INCREMENT, 
  `rol_id` INT NOT NULL, 
  INDEX `usuario_id_idx` (`usuario_id` ASC) VISIBLE, 
  INDEX `rol_id_idx` (`rol_id` ASC) VISIBLE, 
  CONSTRAINT `fk_usuario_user_rol` FOREIGN KEY (`usuario_id`) REFERENCES `melisprint`.`users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION, 
  CONSTRAINT `fk_rol_user_rol` FOREIGN KEY (`rol_id`) REFERENCES `melisprint`.`roles` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB;

DROP 
  TABLE IF EXISTS logs;
CREATE TABLE IF NOT EXISTS `melisprint`.`logs` (
  `id` INT NOT NULL AUTO_INCREMENT, 
  `method` VARCHAR(255) NOT NULL, 
  `label` VARCHAR(255) NOT NULL, 
  `level` VARCHAR(255) NOT NULL, 
  `message` VARCHAR(255) NOT NULL, 
  `status` INT NOT NULL, 
  `insert_date` DATETIME(6) NOT NULL, 
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('Brazil');
INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('United States');

INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('São Paulo', 1);
INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('California', 2);

INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('São Paulo City', 1);
INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('Los Angeles', 2);

INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('123456789', 'Seller 1', 'Address 1', '123456789', 1);
INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('987654321', 'Seller 2', 'Address 2', '987654321', 2);

INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 1');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 2');

INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `lenght`, `netweight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `id_product_type`, `id_seller`) VALUES ('P001', 'Product 1', '10', 5.5, 8.2, 100.25, 0.8, -18, 0.5, 1, 1);
INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `lenght`, `netweight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `id_product_type`, `id_seller`) VALUES ('P002', 'Product 2', '7.5', 3.2, 6.7, 75.5, 0.9, -15, 0.3, 2, 2);

INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES ('Warehouse 1 Address', '111111111', 'W001', 100, -20, 1);
INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES ('Warehouse 2 Address', '222222222', 'W002', 150, -18, 2);

INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `id_product_type`) VALUES (1, -18, -20, 50, 20, 100, 1, 1);
INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `id_product_type`) VALUES (2, -15, -18, 60, 30, 150, 2, 2);

INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (1, 200, -18, '2023-07-31 00:00:00', 300, '2023-07-01 00:00:00', 8, -20, 1, 1);
INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (2, 150, -15, '2023-08-15 00:00:00', 200, '2023-07-10 00:00:00', 9, -18, 2, 2);

INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 10.50, 15.00, 1);
INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 8.75, 12.50, 2);

INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('987654321', 'John', 'Doe');
INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('123456789', 'Jane', 'Smith');

INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('111111', 'Carrier 1', 'Carrier Address 1', '111111111', 1);
INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('222222', 'Carrier 2', 'Carrier Address 2', '222222222', 2);

INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Pending');
INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Processing');

INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO001', '2023-07-01 10:00:00', 'TRACK001', 1, 1, 1, 1, 1);
INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO002', '2023-07-02 11:00:00', 'TRACK002', 2, 2, 2, 2, 2);

INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Clean', 10, -18, 1, 1);
INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Not clean', 20, -15, 2, 2);

INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('123456', 'John', 'Smith', 1);
INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('654321', 'Jane', 'Doe', 2);

INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-05 14:00:00', 'INB001', 1, 1, 1);
INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-06 15:00:00', 'INB002', 2, 2, 2);

INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Administrator', 'admin');
INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Employee', 'employee');

INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password1', 'user1');
INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password2', 'user2');

INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (1, 1);
INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (2, 2);

INSERT INTO `melisprint`.`logs` (`method`, `label`, `level`, `message`, `status`, `insert_date`) VALUES ('GET', 'API Request', 'Info', 'API request received', 200, '2023-07-05 16:00:00');
INSERT INTO `melisprint`.`logs` (`method`, `label`, `level`, `message`, `status`, `insert_date`) VALUES ('POST', 'Data Update', 'Warning', 'Data update failed', 500, '2023-07-05 17:00:00');
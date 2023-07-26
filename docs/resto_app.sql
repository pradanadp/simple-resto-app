CREATE TABLE `customers` (
  `id` integer PRIMARY KEY,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `phone_number` varchar(20) UNIQUE NOT NULL,
  `email` varchar(50) UNIQUE NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `menu_items` (
  `id` integer PRIMARY KEY,
  `item_name` varchar(255) UNIQUE NOT NULL,
  `description` varchar(1000) NOT NULL,
  `price` decimal(10,2),
  `categories` ENUM ('appetizers', 'main_course', 'desserts') NOT NULL,
  `availability` ENUM ('in_stock', 'out_of_stock') NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `orders` (
  `id` integer PRIMARY KEY,
  `customer_id` interger,
  `total_amount` decimal(10, 2) NOT NULL,
  `payment_status` ENUM ('paid', 'pending', 'cancelled') NOT NULL,
  `order_status` ENUM ('pending', 'preparing', 'ready', 'delivered') NOT NULL,
  `delivery_address` varchar(255),
  `contact_info` varchar(255),
  `notes` varchar(1000),
  `created_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `order_items` (
  `id` integer PRIMARY KEY,
  `order_id` integer,
  `item_id` integer,
  `quantity` integer NOT NULL,
  `subtotal` decimal(10,2) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `purchase_receipts` (
  `id` integer PRIMARY KEY,
  `order_id` integer UNIQUE,
  `paymend_method` ENUM ('credit_card', 'cash', 'transfer', 'e_wallet', 'e_money', 'qr_code') NOT NULL,
  `receipt_total` decimal(10, 2) NOT NULL,
  `receipt_date` date,
  `additional_details` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `carts` (
  `id` integer PRIMARY KEY,
  `customer_id` integer,
  `item_id` integer,
  `quantity` integer NOT NULL,
  `created_at` timestamp DEFAULT (now())
);

ALTER TABLE `orders` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`item_id`) REFERENCES `menu_items` (`id`);

ALTER TABLE `purchase_receipts` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `carts` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

ALTER TABLE `carts` ADD FOREIGN KEY (`item_id`) REFERENCES `menu_items` (`id`);

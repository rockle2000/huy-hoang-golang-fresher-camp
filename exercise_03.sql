CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `restaurant_id` int NOT NULL,
  `image` json,
  `price` int,
  `rating` float 
  `description` varchar(255),
  `status` tinyint DEFAULT 1,
  `created_at` timestamp DEFAULT current_timestamp(0),
  `updated_at` timestamp DEFAULT current_timestamp(0) ON UPDATE CURRENT_TIMESTAMP(0)
);

CREATE TABLE `restaurants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `address` varchar(255),
  `avatar` json,
  `rating` float
  `is_verified` tinyint,
  `status` tinyint DEFAULT 1,
  `created_at` timestamp DEFAULT current_timestamp(0)
  `updated_at` timestamp DEFAULT current_timestamp(0) ON UPDATE CURRENT_TIMESTAMP(0)
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `image` json,
  `status` tinyint DEFAULT 1,
  `created_at` timestamp DEFAULT current_timestamp(0),
  `updated_at` timestamp DEFAULT current_timestamp(0) ON UPDATE CURRENT_TIMESTAMP(0)
);

CREATE TABLE `product_category` (
  `product_id` int NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`product_id`, `category_id`)
);

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(50) NOT NULL,
  `firstname` varchar(50) NOT NULL,
  `lastname` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phonenumber` varchar(50) NOT NULL,
  `address` varchar(50) NOT NULL,
  `avatar` json
  `status` tinyint DEFAULT 1,
  `created_at` timestamp DEFAULT current_timestamp(0),
  `updated_at` timestamp DEFAULT current_timestamp(0) ON UPDATE CURRENT_TIMESTAMP(0)
  PRIMARY KEY (`product_id`, `category_id`)
);

ALTER TABLE `products` ADD CONSTRAINT FK_Product_Restaurant FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`);

ALTER TABLE `product_category` ADD CONSTRAINT FK_ProdCate_Product FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `product_category` ADD CONSTRAINT FK_ProdCate_Category FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);


CREATE DATABASE IF NOT EXISTS blog DEFAULT CHARSET utf8mb4 ;
USE blog;
SET NAMES utf8mb4;

CREATE TABLE `person` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int DEFAULT NULL,
    `username` varchar(60) DEFAULT NULL,
    `sex` varchar(10) DEFAULT NULL,
    `email` varchar(20) DEFAULT NULL,
    PRIMARY KEY (`id`, `user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
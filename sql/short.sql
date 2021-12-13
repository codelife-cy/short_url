CREATE TABLE `short` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `long_url` varchar(10240) DEFAULT NULL,
     `short_url` varchar(11) DEFAULT NULL,
     `note` varchar(100) DEFAULT NULL,
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4
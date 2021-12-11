CREATE TABLE `short` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `long_url` varchar(100) DEFAULT NULL,
     `short_url` varchar(100) DEFAULT NULL,
     `note` varchar(100) DEFAULT NULL,
     `create_time` datetime DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4
CREATE TABLE `sequence` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `stub` varchar(32) NOT NULL,
    `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uniq_sequence_stub` (`stub`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8
CREATE TABLE `todo` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `title` varchar(200) DEFAULT NULL,
    `description` varchar(1024) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`id`)
);

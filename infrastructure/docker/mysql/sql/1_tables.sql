USE myretrocollection;

CREATE TABLE `platforms` (
    `platform_uuid` varchar(40) NOT NULL DEFAULT '',
    `name` varchar(200) NOT NULL DEFAULT '',
    `slug` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`platform_uuid`),
    UNIQUE KEY `name,slug` (`slug`,`name`),
    KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
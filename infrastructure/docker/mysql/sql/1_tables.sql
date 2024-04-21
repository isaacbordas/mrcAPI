USE myretrocollectiondb;

CREATE TABLE `platforms` (
    `platform_uuid` varchar(40) NOT NULL DEFAULT '',
    `platform_id` int(11) unsigned NOT NULL,
    `platform_name` varchar(200) NOT NULL DEFAULT '',
    `platform_slug` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`platform_uuid`),
    UNIQUE KEY `platform_name,platform_slug` (`platform_slug`,`platform_name`),
    KEY `platform_name` (`platform_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `genres` (
     `genre_uuid` varchar(40) NOT NULL DEFAULT '',
     `genre_id` int(11) unsigned NOT NULL,
     `genre_name` varchar(200) NOT NULL DEFAULT '',
     PRIMARY KEY (`genre_uuid`),
     UNIQUE KEY `genre_name` (`genre_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `developers` (
    `developer_uuid` varchar(40) NOT NULL DEFAULT '',
    `developer_id` int(11) unsigned NOT NULL,
    `developer_name` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`developer_uuid`),
    UNIQUE KEY `developer_name` (`developer_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `publishers` (
    `publisher_uuid` varchar(40) NOT NULL DEFAULT '',
    `publisher_id` int(11) unsigned NOT NULL,
    `publisher_name` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`publisher_uuid`),
    UNIQUE KEY `publisher_name` (`publisher_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `games` (
    `game_uuid` varchar(40) NOT NULL DEFAULT '',
    `game_id` int(11) unsigned NOT NULL,
    `game_name` varchar(200) NOT NULL DEFAULT '',
    `release_date` varchar(10) DEFAULT '',
    `platform` int(11) unsigned DEFAULT NULL,
    `players` int(11) DEFAULT NULL,
    `overview` varchar(200) DEFAULT NULL,
    `developers` varchar(200) DEFAULT NULL,
    `genres` varchar(200) DEFAULT NULL,
    `publishers` varchar(200) DEFAULT NULL,
    `alternates` varchar(200) DEFAULT NULL,
    `rating` varchar(200) DEFAULT NULL,
    `region` int(2) unsigned DEFAULT NULL,
    PRIMARY KEY (`game_uuid`),
    KEY `name_platform` (`game_name`, `platform`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `regions` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` text NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
INSERT INTO `regions` VALUES (1,'NTSC'),(2,'NTSC-U'),(3,'NTSC-C'),(4,'NTSC-J'),(5,'NTSC-K'),(6,'PAL'),(7,'PAL-A'),(8,'PAL-B'),(9,'Other');
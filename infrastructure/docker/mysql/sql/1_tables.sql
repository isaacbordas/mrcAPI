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
    `game_slug` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`game_uuid`),
    UNIQUE KEY `game_name,game_slug` (`game_slug`,`game_name`),
    KEY `game_name` (`game_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
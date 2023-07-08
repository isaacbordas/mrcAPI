USE myretrocollectiondb;

CREATE TABLE `platforms` (
    `platform_uuid` varchar(40) NOT NULL DEFAULT '',
    `platform_name` varchar(200) NOT NULL DEFAULT '',
    `platform_slug` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`platform_uuid`),
    UNIQUE KEY `platform_name,platform_slug` (`platform_slug`,`platform_name`),
    KEY `platform_name` (`platform_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `games` (
    `game_uuid` varchar(40) NOT NULL DEFAULT '',
    `game_name` varchar(200) NOT NULL DEFAULT '',
    `game_slug` varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`game_uuid`),
    UNIQUE KEY `game_name,game_slug` (`game_slug`,`game_name`),
    KEY `game_name` (`game_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
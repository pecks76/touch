CREATE TABLE `client`
(
    `id` int NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `pot`
(
    `id`       int        NOT NULL AUTO_INCREMENT,
    `name`     varchar(1) NOT NULL,
    `clientId` int        NOT NULL,
    `depositId` int        NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `account`
(
    `id`           int         NOT NULL AUTO_INCREMENT,
    `wrapper_type` varchar(16) NOT NULL,
    `potId`        int         NOT NULL,
    `amount`       int         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `deposit`
(
    `id`       int NOT NULL AUTO_INCREMENT,
    `clientId` int NOT NULL,
    `nominal`  int NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `instruction`
(
    `id`           int         NOT NULL AUTO_INCREMENT,
    `depositId`    int         NOT NULL,
    `potName`      varchar(1)  NOT NULL,
    `wrapper_type` varchar(16) NOT NULL,
    `amount`       int         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
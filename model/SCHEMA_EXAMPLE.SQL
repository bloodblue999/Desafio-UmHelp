CREATE DATABASE IF NOT EXISTS um_help;
USE um_help;

DROP TABLE IF EXISTS `tb_money_transaction`;
DROP TABLE IF EXISTS `tb_wallet`;
DROP TABLE IF EXISTS `tb_user_account`;
DROP TABLE IF EXISTS `tb_currency`;

CREATE TABLE `tb_user_account`
(
    `user_account_id` BIGINT AUTO_INCREMENT,
    `public_id`       VARCHAR(36)  NOT NULL DEFAULT (UUID()),
    `first_name`      VARCHAR(30)  NOT NULL,
    `last_name`       VARCHAR(255) NOT NULL,
    `document`        VARCHAR(14)  NOT NULL UNIQUE,
    `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`      DATETIME     NULL,

    PRIMARY KEY (`user_account_id`)
) ENGINE = innodb;


CREATE TABLE `tb_money_transaction`
(
    `money_transaction_id` BIGINT AUTO_INCREMENT,
    `money_value`          BIGINT   NOT NULL,
    `processing_data`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `sender_id`            BIGINT   NOT NULL,
    `receiver_id`          BIGINT   NOT NULL,

    PRIMARY KEY (`money_transaction_id`),
    FOREIGN KEY (`sender_id`) REFERENCES tb_user_account (`user_account_id`),
    FOREIGN KEY (`receiver_id`) REFERENCES tb_user_account (`user_account_id`)

) ENGINE = innodb;

CREATE TABLE `tb_currency`
(
    `currency_id` BIGINT AUTO_INCREMENT,
    `code`        VARCHAR(4) NOT NULL UNIQUE,
    `symbol`      VARCHAR(4) NOT NULL UNIQUE,
    `created_at`  DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  DATETIME   NULL,

    PRIMARY KEY (`currency_id`)
) ENGINE = innodb;

CREATE TABLE tb_wallet
(
    `wallet_id`   BIGINT AUTO_INCREMENT,
    `balance`     BIGINT      NOT NULL DEFAULT 0,
    `alias`       VARCHAR(45) NOT NULL,
    `owner_id`    BIGINT      NOT NULL,
    `currency_id` BIGINT      NOT NULL,
    `created_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  DATETIME    NULL,

    PRIMARY KEY (`wallet_id`),
    FOREIGN KEY (`owner_id`) REFERENCES `tb_user_account` (`user_account_id`),
    FOREIGN KEY (`currency_id`) REFERENCES `tb_currency` (`currency_id`)

) ENGINE = innodb;

INSERT INTO tb_currency (`code`, `symbol`)
VALUES ('BRL', 'R$');

INSERT INTO `tb_currency` (`code`, `symbol`)
VALUES ('USD', '$');

INSERT INTO `tb_currency` (`code`, `symbol`)
VALUES ('EUR', '€');
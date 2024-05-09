CREATE DATABASE IF NOT EXISTS um_help;
USE um_help;

DROP TABLE IF EXISTS tb_user_account;

CREATE TABLE tb_user_account
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(30)  NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    document   VARCHAR(14)  NOT NULL,
    balance    BIGINT       NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME     NULL
) ENGINE = innodb;


CREATE TABLE tb_money_transaction
(
    id              BIGINT PRIMARY KEY AUTO_INCREMENT,
    money_value     BIGINT   NOT NULL,
    processing_data DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sender_id       BIGINT   NOT NULL,
    receiver_id     BIGINT   NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES tb_user_account (id),
    FOREIGN KEY (receiver_id) REFERENCES tb_user_account (id)

) ENGINE = innodb;
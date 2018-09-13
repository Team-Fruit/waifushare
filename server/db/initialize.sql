CREATE USER 'waifushare'@'%';
GRANT INSERT,SELECT,UPDATE,DELETE ON `waifushare_db`.* TO 'waifushare'@'%';

CREATE DATABASE IF NOT EXISTS `waifushare_db`;

CREATE TABLE IF NOT EXISTS `waifushare_db`.`user` (
    `id`                INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `user_id`           VARCHAR(15)     NOT NULL,
    `display_name`      TINYTEXT,
    `password_hash`     TINYTEXT,
    `saved_image_id`    INT UNSIGNED,
    `account_status`    INT             NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `user_id` ( `user_id` )
);

CREATE TABLE IF NOT EXISTS `waifushare_db`.`twimg` (
    `id`            INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `twitter_id`    VARCHAR(255)    NOT NULL,
    `file_name`     TINYTEXT        NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `twitter_id` ( `twitter_id` )
);

CREATE TABLE IF NOT EXISTS `waifushare_db`.`image` (
    `id`            INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `created_at`    DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by`    INT UNSIGNED    NOT NULL,
    `twitter_id`    TINYTEXT,

    PRIMARY KEY ( `id` ),

    CONSTRAINT `fk__image__user__id`
        FOREIGN KEY ( `created_by` )
        REFERENCES `waifushare_db`.`user` ( `id` )
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `waifushare_db`.`user__image` (
    `user_id`   INT UNSIGNED    NOT NULL,
    `image_id`  INT UNSIGNED    NOT NULL,
    `is_like`   INT             NOT NULL,

    PRIMARY KEY ( `user_id`, `image_id` ),

    CONSTRAINT `fk__user__image__user__id`
        FOREIGN KEY ( `user_id` )
        REFERENCES `waifushare_db`.`user` ( `id` )
        ON DELETE CASCADE,
    CONSTRAINT `fk__user__image__image__id`
        FOREIGN KEY ( `image_id` )
        REFERENCES `waifushare_db`.`image` ( `id` )
        ON DELETE CASCADE
);


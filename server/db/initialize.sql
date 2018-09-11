CREATE USER 'waifushare'@'%';
GRANT INSERT,SELECT,UPDATE,DELETE ON `waifushare_db`.* TO 'waifushare'@'%';

CREATE DATABASE IF NOT EXITS 'waifushare_db';

CREATE TABLE IF NOT EXITS `waifushare_db`.`user` (
    `id`                INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `user_id`           TINYTEXT        NOT NULL,
    `display_name`      TINYTEXT        NOT NULL,
    `password_hash`     TINYTEXT        NOT NULL,
    `saved_image_id`    INT UNSIGNED    NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `user_id` ( `user_id` )
);

CREATE TABLE IF NOT EXITS `waifushare_db`.`twimg` (
    `id`            INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `twitter_id`    TINYTEXT        NOT NULL,
    `file_name`     TINYTEXT        NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `twitter_id` ( `twitter_id` )
);

CREATE TABLE IF NOT EXITS `waifushare_image`.`image` (
    `id`            INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `created_at`    DATETIME        NOT NULL CURRENT_TIMESTAMP,
    `created_by`    INT UNSIGNED    NOT NULL,
    `twitter_id`    TINYTEXT        NOT NULL,

    PRIMARY KEY ( `id` ),

    CONSTRAINT `fk__image__user__id`
        FOREIGN KEY ( `created_by` )
        REFERENCES `waifushare_db`.`user` ( `id` )
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXITS `waifushare_image`.`user__image` (
    `user_id`   INT UNSIGNED    NOT NULL,
    `image_id`  INT UNSIGNED    NOT NULL,
    `is_like`   INT             NOT NULL,

    CONSTRAINT `fk__user__image__user__id`
        FOREIGN KEY ( `user_id` )
        REFERENCES `waifushare_db`.`user` ( `id` )
        ON DELETE CASCADE,
    CONSTRAINT `fk__user__image__image__id`
        FOREIGN KEY ( `image_id` )
        REFERENCES `waifushare_db`.`image` ( `id` )
        ON DELETE CASCADE
);


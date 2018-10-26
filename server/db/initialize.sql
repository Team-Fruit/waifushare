CREATE USER 'waifushare'@'%';
GRANT INSERT,SELECT,UPDATE,DELETE ON `waifushare_db`.* TO 'waifushare'@'%';

CREATE DATABASE IF NOT EXISTS `waifushare_db`;

CREATE TABLE IF NOT EXISTS `waifushare_db`.`user` (
    `id`               INT UNSIGNED    NOT NULL AUTO_INCREMENT,
    `screen_name`      VARCHAR(15)     NOT NULL,
    `password_hash`    TINYTEXT        NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `screen_name` ( `screen_name` )
);

CREATE TABLE IF NOT EXISTS `waifushare_db`.`twimg` (
    `id`                     INT UNSIGNED       NOT NULL AUTO_INCREMENT,
    `filename`               VARCHAR(255)       NOT NULL,
    `tweet_id`               BIGINT UNSIGNED    NOT NULL,
    `twitter_screen_name`    VARCHAR(15)        NOT NULL,
    `twimg_filename`         TINYTEXT           NOT NULL,

    PRIMARY KEY ( `id` ),
    UNIQUE KEY `filename` ( `filename` )
);

CREATE TABLE IF NOT EXISTS `waifushare_db`.`user__image` (
    `user_id`       INT UNSIGNED    NOT NULL,
    `image_id`      INT UNSIGNED    NOT NULL,
    `created_at`    DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY ( `user_id`, `image_id` ),

    CONSTRAINT `fk__user__image__user__id`
        FOREIGN KEY ( `user_id` )
        REFERENCES `waifushare_db`.`user` ( `id` )
        ON DELETE CASCADE,
    CONSTRAINT `fk__user__image__image__id`
        FOREIGN KEY ( `image_id` )
        REFERENCES `waifushare_db`.`twimg` ( `id` )
        ON DELETE CASCADE
);


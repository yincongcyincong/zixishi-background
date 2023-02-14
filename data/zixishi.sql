CREATE DATABASE `study_room`;
CREATE TABLE `seatinfo`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `sid`          bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '自习室id',
    `seat_type_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '座位类型id',
    `create_time`  int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`  int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='座位表';

CREATE TABLE `study_room`
(
    `id`    bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `intro` text                not null DEFAULT '' COMMENT '介绍',
    `name`  varchar(255)        not null DEFAULT '' COMMENT '自习室名称',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='自习室表';

CREATE TABLE `seat_type`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `sid`         bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '自习室id',
    `price_intro` text                not null DEFAULT '' COMMENT '价格介绍，json格式',
    `intro`       text                not null DEFAULT '' COMMENT '介绍',
    `name`        varchar(255)        not null DEFAULT '' COMMENT '类型名称',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='自习室类型表';


CREATE TABLE `buy_record`
(
    `id`       bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `sid`      bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '自习室id',
    `uid`      varchar(255)        not null DEFAULT '' COMMENT '用户微信凭证',
    `uname`    varchar(255)        not null DEFAULT '' COMMENT '用户名',
    `price`    int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '购买，单位：分',
    `end_time` int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '到期时间',
    `buy_time` int(10) unsigned    NOT NULL DEFAULT 0 COMMENT '购买时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='购买记录表';

INSERT INTO study_room
values (1, '这是一个自习室');
INSERT INTO seat_type
values (1, 1, '{"按日缴费": 9,"按日缴费": 59,"按日缴费": 269}', '休闲舒适区[图片图片]', '休闲舒适区');
INSERT INTO seatinfo
values (1, 1, 1, 1676292591, 1676292591);
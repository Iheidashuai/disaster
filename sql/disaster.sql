CREATE TABLE role
(
    id            INT auto_increment PRIMARY KEY NOT NULL COMMENT '用户 id',
    attack_power  INT UNSIGNED                   NOT NULL DEFAULT 0 COMMENT '攻击力',
    defense       INT UNSIGNED                   NOT NULL DEFAULT 0 COMMENT '防御力',
    life          BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '生命值',
    hat           BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '帽子 ',
    clothing      BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '衣服',
    shoes         BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '鞋子',
    weapon        BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '武器',
    pants         BIGINT UNSIGNED                NOT NULL DEFAULT 0 COMMENT '裤子',
    level         INT UNSIGNED                   NOT NULL DEFAULT 1 COMMENT '等级',
    military_rank INT UNSIGNED                   NOT NULL DEFAULT 0 COMMENT '军衔',

    ctime         TIMESTAMP                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    mtime         TIMESTAMP                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)
CREATE TABLE user
(
    id         int PRIMARY KEY auto_increment,
    uid        int             NOT NULL,
    username   VARCHAR(50)     NOT NULL,
    password   VARCHAR(50)     NOT NULL,
    salt       VARCHAR(50)     NOT NULL,
    email      VARCHAR(50)     NOT NULL,
    phone      VARCHAR(50)     NOT NULL,
    age        INT             NOT NULL,
    sex        INT             NOT NULL,
    face       VARCHAR(50)     NOT NULL,
    experience INT             NOT NULL,
    vip        INT             NOT NULL,
    country    INT             NOT NULL,
    level      BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级',

    ctime      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    mtime      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);

CREATE TABLE user_attr
(
    id             int PRIMARY KEY auto_increment,
    uid            int             NOT NULL,
    attack_power   INT UNSIGNED    NOT NULL DEFAULT 0 COMMENT '攻击力',
    defense        INT UNSIGNED    NOT NULL DEFAULT 0 COMMENT '防御力',
    vitality       BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '生命值',
    armor          BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '护甲',
    armor_piercing BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '穿甲',
    crit           BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '暴击',

    ctime          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    mtime          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
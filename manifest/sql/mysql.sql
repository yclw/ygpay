-- 创建数据库
CREATE DATABASE IF NOT EXISTS ygpay DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 切换数据库
USE ygpay;

-- 用户信息表
CREATE TABLE t_member_info (
    id                   BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    uid                  VARCHAR(36)  DEFAULT '' NOT NULL COMMENT '用户UID',
    username             VARCHAR(20)  DEFAULT '' NOT NULL COMMENT '帐号',
    password_hash        VARCHAR(255) DEFAULT '' NOT NULL COMMENT '密码哈希',
    avatar               VARCHAR(255) DEFAULT '' NULL COMMENT '头像',
    sex                  TINYINT(1)   DEFAULT 1  NULL COMMENT '性别: 1男 2女 3未知',
    email                VARCHAR(60)  DEFAULT '' NULL COMMENT '邮箱',
    mobile               VARCHAR(20)  DEFAULT '' NULL COMMENT '手机号码',
    address              VARCHAR(100) DEFAULT '' NULL COMMENT '联系地址',
    last_active_at       DATETIME                NULL COMMENT '最后活跃时间',
    remark               VARCHAR(255)            NULL COMMENT '备注',
    sort                 INT          DEFAULT 0  NOT NULL COMMENT '排序',
    status               TINYINT(1)   DEFAULT 1  NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at           DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at           DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    -- 索引
    UNIQUE INDEX uniq_uid (uid),
    UNIQUE INDEX uniq_username (username),
    INDEX idx_email (email),
    INDEX idx_mobile (mobile),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) COMMENT '用户信息表' CHARSET=utf8mb4;

-- 系统配置表
CREATE TABLE t_sys_config (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '配置ID',
    `group`     VARCHAR(128) DEFAULT 'default' NOT NULL COMMENT '配置分组',
    `key`       VARCHAR(100) NOT NULL COMMENT '参数键名',
    value       VARCHAR(2000) NOT NULL COMMENT '参数值',
    description VARCHAR(200) NULL COMMENT '配置描述',
    sort        INT          DEFAULT 0 NOT NULL COMMENT '排序',
    status      TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at  DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE INDEX uniq_key (`key`),
    INDEX idx_group (`group`),
    INDEX idx_status (status)
) COMMENT '系统配置表' CHARSET=utf8mb4;

-- 角色信息表
CREATE TABLE t_role_info (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID',
    name       VARCHAR(32)  NOT NULL COMMENT '角色名称',
    `key`      VARCHAR(128) NOT NULL UNIQUE COMMENT '角色权限字符串',
    remark     VARCHAR(255) NULL COMMENT '备注',
    sort       INT          DEFAULT 0 NOT NULL COMMENT '排序',
    status     TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_status (status)
) COMMENT '角色信息表' CHARSET=utf8mb4;

-- API信息表
CREATE TABLE t_api_info (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'API ID',
    name        VARCHAR(128) NOT NULL COMMENT 'API名称',
    path        VARCHAR(255) NOT NULL COMMENT 'API路径',
    method      VARCHAR(10)  NOT NULL COMMENT 'API方法',
    group_name  VARCHAR(64)  DEFAULT '' COMMENT 'API分组',
    description VARCHAR(255) NULL COMMENT 'API描述',
    need_auth   TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '是否需要认证: 0否 1是',
    rate_limit  INT          DEFAULT 0 NOT NULL COMMENT '限流次数/分钟',
    sort        INT          DEFAULT 0 NOT NULL COMMENT '排序',
    status      TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at  DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_path (path),
    INDEX idx_method (method),
    INDEX idx_group_name (group_name),
    INDEX idx_status (status),
    INDEX idx_method_path (method, path)
) COMMENT 'API信息表' CHARSET=utf8mb4;

-- 菜单信息表
CREATE TABLE t_menu_info (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '菜单ID',
    name        VARCHAR(128) NOT NULL COMMENT '菜单名称',
    path        VARCHAR(255) NOT NULL COMMENT '菜单路径',
    icon        VARCHAR(128) DEFAULT '' COMMENT '菜单图标',
    title       VARCHAR(128) DEFAULT '' COMMENT '菜单标题',
    showParent  TINYINT(1)   DEFAULT 0 NOT NULL COMMENT '是否显示父菜单: 0是 1否',
    component   VARCHAR(200) DEFAULT '' COMMENT '组件路径',
    noShowingChildren TINYINT(1)   DEFAULT 0 NOT NULL COMMENT '是否显示子菜单: 0是 1否',
    value           VARCHAR(128) DEFAULT '' COMMENT '菜单值',
    showTooltip     TINYINT(1)   DEFAULT 0 NOT NULL COMMENT '是否显示提示: 0是 1否',
    parentId        BIGINT       NULL COMMENT '父菜单ID',
    redirect        VARCHAR(255) DEFAULT '' COMMENT '重定向',
    description VARCHAR(255) NULL COMMENT '菜单描述',
    sort        INT          DEFAULT 10 NOT NULL COMMENT '排序',
    status      TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at  DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_path (path),
    INDEX idx_status (status)
) COMMENT '菜单信息表' CHARSET=utf8mb4;

-- 登录日志表
CREATE TABLE t_log_login (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '登录日志ID',
    member_id       BIGINT       NULL COMMENT '用户ID',
    username        VARCHAR(20)  DEFAULT '' NOT NULL COMMENT '登录账号',
    ip_address      VARCHAR(45)  DEFAULT '' NOT NULL COMMENT 'IP地址',
    user_agent      VARCHAR(500) DEFAULT '' COMMENT '用户代理',
    login_location  VARCHAR(100) DEFAULT '' COMMENT '登录地点',
    browser         VARCHAR(50)  DEFAULT '' COMMENT '浏览器',
    os              VARCHAR(50)  DEFAULT '' COMMENT '操作系统',
    login_status    TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '登录状态: 0失败 1成功',
    login_message   VARCHAR(255) DEFAULT '' COMMENT '登录信息',
    login_time      DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
    INDEX idx_member_id (member_id),
    INDEX idx_username (username),
    INDEX idx_ip_address (ip_address),
    INDEX idx_login_status (login_status),
    INDEX idx_login_time (login_time),
    INDEX idx_member_login_time (member_id, login_time)
) COMMENT '登录日志表' CHARSET=utf8mb4;

-- 角色菜单关联表
CREATE TABLE t_role_menu (
    role_id BIGINT NOT NULL COMMENT '角色ID',
    menu_id BIGINT NOT NULL COMMENT '菜单ID',
    PRIMARY KEY (role_id, menu_id)
) COMMENT '角色菜单关联表' CHARSET=utf8mb4;

-- 用户角色关联表
CREATE TABLE t_member_role (
    member_id BIGINT NOT NULL COMMENT '用户ID',
    role_id   BIGINT NOT NULL COMMENT '角色ID',
    PRIMARY KEY (member_id, role_id)
) COMMENT '用户角色关联表' CHARSET=utf8mb4;

-- 角色API关联表
CREATE TABLE t_role_api (
    role_id BIGINT NOT NULL COMMENT '角色ID',
    api_id  BIGINT NOT NULL COMMENT 'API ID',
    PRIMARY KEY (role_id, api_id)
) COMMENT '角色API关联表' CHARSET=utf8mb4;

-- 菜单树关系表
CREATE TABLE t_menu_tree (
    id  BIGINT NOT NULL COMMENT '菜单ID',
    pid BIGINT NOT NULL COMMENT '父菜单ID',
    PRIMARY KEY (id, pid)
) COMMENT '菜单树关系表' CHARSET=utf8mb4;

-- 角色树关系表
CREATE TABLE t_role_tree (
    id  BIGINT NOT NULL COMMENT '角色ID',
    pid BIGINT NOT NULL COMMENT '父角色ID',
    PRIMARY KEY (id, pid)
) COMMENT '角色树关系表' CHARSET=utf8mb4;

-- Casbin规则表
CREATE TABLE t_casbin_rule (
    id    INT AUTO_INCREMENT PRIMARY KEY,
    ptype VARCHAR(255) DEFAULT '' NOT NULL,
    v0    VARCHAR(255) DEFAULT '' NOT NULL,
    v1    VARCHAR(255) DEFAULT '' NOT NULL,
    v2    VARCHAR(255) DEFAULT '' NOT NULL,
    v3    VARCHAR(255) DEFAULT '' NOT NULL,
    v4    VARCHAR(255) DEFAULT '' NOT NULL,
    v5    VARCHAR(255) DEFAULT '' NOT NULL,
    INDEX idx_ptype (ptype),
    INDEX idx_v0 (v0),
    INDEX idx_v1 (v1)
) COMMENT 'Casbin规则表' CHARSET=utf8mb4;
-- 用户信息表
create table t_member_info
(
    id                   bigint auto_increment comment '用户ID' primary key,
    username             varchar(20)  default '' not null comment '帐号',
    password_hash        char(32)     default '' not null comment '密码',
    salt                 char(16)                not null comment '密码盐',
    password_reset_token varchar(150) default '' null comment '密码重置令牌',
    role_id              bigint       default 10 null comment '角色ID',
    avatar               char(150)    default '' null comment '头像',
    sex                  tinyint(1)   default 1  null comment '性别',
    email                varchar(60)  default '' null comment '邮箱',
    mobile               varchar(20)  default '' null comment '手机号码',
    address              varchar(100) default '' null comment '联系地址',
    last_active_at       datetime                null comment '最后活跃时间',
    remark               varchar(255)            null comment '备注',
    status               tinyint(1)   default 1  null comment '状态',
    created_at           DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at           DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
) comment '用户信息表';


-- 系统配置表
CREATE TABLE t_sys_config (
  id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '配置ID',
  `group`     VARCHAR(128) NOT NULL DEFAULT 'default' COMMENT '配置分组',
  `key`       VARCHAR(100) NOT NULL COMMENT '参数键名（唯一）',
  value       VARCHAR(2000) NOT NULL COMMENT '参数值',
  description VARCHAR(200) NULL COMMENT '配置描述',
  sort        INT NOT NULL DEFAULT 0 COMMENT '排序',
  status      tinyint(1)   default 1  null comment '状态',
  created_at  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE INDEX uniq_key (`key`),               -- 键名唯一
  INDEX idx_group (`group`)                    -- 按分组查询
) COMMENT '系统配置表';

-- 角色信息表
CREATE TABLE t_role_info (
  id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID',
  name VARCHAR(32) NOT NULL COMMENT '角色名称',
  `key` VARCHAR(128) NOT NULL UNIQUE COMMENT '角色权限字符串',
  pid BIGINT DEFAULT 0 COMMENT '上级角色ID',
  level INT NOT NULL DEFAULT 1 COMMENT '关系树等级',
  tree VARCHAR(1000) NULL COMMENT '关系树路径',
  remark VARCHAR(255) NULL COMMENT '备注',
  sort INT NOT NULL DEFAULT 0 COMMENT '排序',
  status TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_tree (tree(100)),
  INDEX idx_pid_level (pid, level)
) COMMENT '角色信息表';

-- 角色菜单关联表
CREATE TABLE t_role_menu (
  role_id BIGINT NOT NULL COMMENT '角色ID',
  menu_id BIGINT NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (role_id, menu_id),
) COMMENT '角色菜单关联表';


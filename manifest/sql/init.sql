-- 配置表

-- 登录配置
INSERT INTO t_sys_config (`group`, `key`, value, description) VALUES ('login', 'loginRegisterSwitch', '1', '登录注册开关');
INSERT INTO t_sys_config (`group`, `key`, value, description) VALUES ('login', 'loginCaptchaSwitch', '1', '登录验证码开关');
INSERT INTO t_sys_config (`group`, `key`, value, description) VALUES ('login', 'loginCaptchaType', '1', '登录验证码类型');
INSERT INTO t_sys_config (`group`, `key`, value, description) VALUES ('login', 'loginProtocol', '1', '登录协议');
INSERT INTO t_sys_config (`group`, `key`, value, description) VALUES ('login', 'loginPolicy', '1', '登录政策');

-- 角色表
-- 超级管理员
INSERT INTO t_role_info (id, name, `key`, remark) VALUES (1, '超级管理员', 'super_admin', '超级管理员');

-- 用户表
-- 超级管理员
INSERT INTO t_member_info (id, uid, username, password_hash, last_active_at, remark) VALUES (1, '1', 'admin', '123456', NOW(), '超级管理员');
INSERT INTO t_member_role (member_id, role_id) VALUES (1, 1);


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
    meta          JSON         DEFAULT '{}' COMMENT '菜单元数据',
    showTooltip     TINYINT(1)   DEFAULT 0 NOT NULL COMMENT '是否显示提示: 0是 1否',
    parentId        BIGINT       NULL COMMENT '父菜单ID',
    redirect        VARCHAR(255) DEFAULT '' COMMENT '重定向',
    description VARCHAR(255) NULL COMMENT '菜单描述',
    sort        INT          DEFAULT 0 NOT NULL COMMENT '排序',
    status      TINYINT(1)   DEFAULT 1 NOT NULL COMMENT '状态: 0禁用 1启用',
    created_at  DATETIME     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_path (path),
    INDEX idx_status (status)
) COMMENT '菜单信息表' CHARSET=utf8mb4;
-- 菜单表
INSERT INTO t_menu_info (id, name, path,title, icon, sort) VALUES (1, 'Permission', '/permission','menus.purePermission', 'ep:lollipop', 10);
INSERT INTO t_menu_tree (id, pid) VALUES (1, 0);

INSERT INTO t_menu_info (id, name, path ,title, component) VALUES (2, 'PermissionButtonRouter', '/permission/button/router','menus.purePermissionButtonRouter', 'permission/button/index');
INSERT INTO t_menu_tree (id, pid) VALUES (2, 1);

-- {
--       path: "/permission/page/index",
--       name: "PermissionPage",
--       meta: {
--         title: "menus.purePermissionPage",
--       }
    -- },
INSERT INTO t_menu_info (id, name, path ,title, component) VALUES (3, 'PermissionPage', '/permission/page/index','menus.purePermissionPage', 'permission/page/index');
INSERT INTO t_menu_tree (id, pid) VALUES (3, 1);

INSERT INTO t_role_menu (role_id, menu_id) VALUES (1, 1);
INSERT INTO t_role_menu (role_id, menu_id) VALUES (1, 2);
INSERT INTO t_role_menu (role_id, menu_id) VALUES (1, 3);

-- api信息表

-- api列表
INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('api列表', '/api/list', 'GET', 'api', '获取api列表', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('api详情', '/api/one', 'GET', 'api', '获取api详情', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('api创建', '/api/create', 'POST', 'api', '创建api', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('api更新', '/api/update', 'PUT', 'api', '更新api', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('api删除', '/api/delete', 'DELETE', 'api', '删除api', 1, 1000, 1, NOW(), NOW());

-- 角色列表
INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('角色列表', '/role/list', 'GET', 'role', '获取角色列表', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('角色详情', '/role/one', 'GET', 'role', '获取角色详情', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('角色创建', '/role/create', 'POST', 'role', '创建角色', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('角色更新', '/role/update', 'PUT', 'role', '更新角色', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('角色删除', '/role/delete', 'DELETE', 'role', '删除角色', 1, 1000, 1, NOW(), NOW());

-- 用户列表

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('用户列表', '/member/list', 'GET', 'member', '获取用户列表', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('用户详情', '/member/one', 'GET', 'member', '获取用户详情', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('用户创建', '/member/create', 'POST', 'member', '创建用户', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('用户更新', '/member/update', 'PUT', 'member', '更新用户', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('用户删除', '/member/delete', 'DELETE', 'member', '删除用户', 1, 1000, 1, NOW(), NOW());

-- 菜单列表

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('菜单列表', '/menu/list', 'GET', 'menu', '获取菜单列表', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('菜单详情', '/menu/one', 'GET', 'menu', '获取菜单详情', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('菜单创建', '/menu/create', 'POST', 'menu', '创建菜单', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('菜单更新', '/menu/update', 'PUT', 'menu', '更新菜单', 1, 1000, 1, NOW(), NOW());

INSERT INTO t_api_info 
(name, path, method, group_name, description, need_auth, sort, status, created_at, updated_at) VALUES 
('菜单删除', '/menu/delete', 'DELETE', 'menu', '删除菜单', 1, 1000, 1, NOW(), NOW());

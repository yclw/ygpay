create table t_api_info
(
    id          bigint auto_increment comment 'API ID'
        primary key,
    api_uid     varchar(36)                            not null,
    name        varchar(128)                           not null comment 'API名称',
    path        varchar(255)                           not null comment 'API路径',
    method      varchar(10)                            not null comment 'API方法',
    group_name  varchar(64)  default ''                null comment 'API分组',
    description varchar(500)                           null,
    need_auth   tinyint(1)   default 1                 not null comment '是否需要认证: 0否 1是',
    rate_limit  int unsigned default '0'               not null,
    sort        int          default 0                 not null comment '排序',
    status      tinyint(1)   default 1                 not null comment '状态: 0禁用 1启用',
    created_at  datetime     default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  datetime     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uniq_api_uid
        unique (api_uid),
    constraint uniq_method_path
        unique (method, path)
)
    comment 'API信息表';

create index idx_group_name
    on t_api_info (group_name);

create index idx_method
    on t_api_info (method);

create index idx_method_path
    on t_api_info (method, path);

create index idx_need_auth
    on t_api_info (need_auth);

create index idx_path
    on t_api_info (path);

create index idx_status
    on t_api_info (status);

create table t_casbin_rule
(
    id    bigint auto_increment
        primary key,
    ptype varchar(100) default '' not null,
    v0    varchar(100) default '' not null,
    v1    varchar(100) default '' not null,
    v2    varchar(100) default '' not null,
    v3    varchar(100) default '' not null,
    v4    varchar(100) default '' not null,
    v5    varchar(100) default '' not null,
    constraint uniq_ptype_v0_v1_v2_v3_v4_v5
        unique (ptype, v0, v1, v2, v3, v4, v5)
)
    comment 'Casbin';

create index idx_ptype
    on t_casbin_rule (ptype);

create index idx_v0
    on t_casbin_rule (v0);

create index idx_v1
    on t_casbin_rule (v1);

create index idx_v2
    on t_casbin_rule (v2);

create index idx_v3
    on t_casbin_rule (v3);

create index idx_v4
    on t_casbin_rule (v4);

create index idx_v5
    on t_casbin_rule (v5);

create table t_log_login
(
    id             bigint auto_increment comment '登录日志ID'
        primary key,
    member_id      bigint                                  null comment '用户ID',
    username       varchar(20)   default ''                not null comment '登录账号',
    ip_address     varchar(45)   default ''                not null comment 'IP地址',
    user_agent     varchar(1000) default ''                null,
    login_location varchar(100)  default ''                null comment '登录地点',
    browser        varchar(50)   default ''                null comment '浏览器',
    os             varchar(50)   default ''                null comment '操作系统',
    login_status   tinyint(1)    default 1                 not null comment '登录状态: 0失败 1成功',
    login_message  varchar(255)  default ''                null comment '登录信息',
    login_time     datetime      default CURRENT_TIMESTAMP null comment '登录时间'
)
    comment '登录日志表';

create index idx_ip_address
    on t_log_login (ip_address);

create index idx_login_status
    on t_log_login (login_status);

create index idx_login_time
    on t_log_login (login_time);

create index idx_member_id
    on t_log_login (member_id);

create index idx_member_login_time
    on t_log_login (member_id, login_time);

create index idx_member_status
    on t_log_login (member_id, login_status);

create index idx_username
    on t_log_login (username);

create table t_member_info
(
    id             bigint auto_increment comment '用户ID'
        primary key,
    uid            varchar(36)  default ''                not null comment '用户UID',
    username       varchar(20)  default ''                not null comment '帐号',
    password_hash  varchar(255) default ''                not null comment '密码哈希',
    avatar         varchar(500) default ''                null,
    sex            tinyint(1)   default 1                 null comment '性别: 1男 2女 3未知',
    email          varchar(100) default ''                null,
    mobile         varchar(20)  default ''                null comment '手机号码',
    address        varchar(200) default ''                null,
    last_active_at datetime                               null comment '最后活跃时间',
    remark         varchar(500)                           null,
    sort           int          default 0                 not null comment '排序',
    status         tinyint(1)   default 1                 not null comment '状态: 0禁用 1启用',
    created_at     datetime     default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at     datetime     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    nickname       varchar(20)  default ''                not null comment '昵称',
    constraint uniq_uid
        unique (uid),
    constraint uniq_username
        unique (username)
)
    comment '用户信息表';

create index idx_created_at
    on t_member_info (created_at);

create index idx_email
    on t_member_info (email);

create index idx_last_active
    on t_member_info (last_active_at);

create index idx_mobile
    on t_member_info (mobile);

create index idx_nickname
    on t_member_info (nickname);

create index idx_status
    on t_member_info (status);

create table t_member_role
(
    member_id  bigint                             not null comment '用户ID',
    role_id    bigint                             not null comment '角色ID',
    created_at datetime default CURRENT_TIMESTAMP null,
    primary key (member_id, role_id)
)
    comment '用户角色关联表';

create table t_menu_info
(
    id         bigint auto_increment comment '菜单ID'
        primary key,
    menu_uid   varchar(36)                          not null,
    pid        bigint     default 0                 not null comment '父菜单ID',
    type       int        default 1                 not null comment '菜单类型: 0目录 1菜单 2外链',
    name       varchar(128)                         not null comment '菜单名称',
    path       varchar(255)                         not null comment '菜单路径',
    title      varchar(128)                         not null comment '菜单标题',
    icon       varchar(128)                         null comment '菜单图标',
    sort       int        default 10                not null comment '排序',
    showParent tinyint(1) default 0                 not null comment '是否显示父菜单: 0是 1否',
    showLink   tinyint(1) default 1                 not null comment '是否显示该菜单: 0是 1否',
    keepAlive  tinyint(1) default 0                 not null comment '是否缓存: 0是 1否',
    redirect   varchar(255)                         null comment '重定向',
    component  varchar(200)                         null comment '组件路径',
    frameSrc   varchar(255)                         null comment '内嵌地址',
    url        varchar(255)                         null comment '外部链接',
    status     tinyint(1) default 1                 not null comment '状态: 0禁用 1启用',
    created_at datetime   default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uniq_menu_uid
        unique (menu_uid),
    constraint uniq_name
        unique (name),
    constraint uniq_path
        unique (path)
)
    comment '菜单信息表';

create index idx_path
    on t_menu_info (path);

create index idx_pid
    on t_menu_info (pid);

create index idx_pid_sort
    on t_menu_info (pid, sort);

create index idx_sort
    on t_menu_info (sort);

create index idx_status
    on t_menu_info (status);

create index idx_type
    on t_menu_info (type);

create table t_role_api
(
    role_id    bigint                             not null comment '角色ID',
    api_id     bigint                             not null comment 'API ID',
    created_at datetime default CURRENT_TIMESTAMP null,
    primary key (role_id, api_id)
)
    comment '角色API关联表';

create index idx_created_at
    on t_role_api (created_at);

create table t_role_info
(
    id         bigint auto_increment comment '角色ID'
        primary key,
    role_uid   varchar(36)                          not null,
    pid        bigint     default 0                 not null comment '父角色ID',
    name       varchar(32)                          not null comment '角色名称',
    `key`      varchar(128)                         not null comment '角色权限字符串',
    remark     varchar(500)                         null,
    sort       int        default 0                 not null comment '排序',
    status     tinyint(1) default 1                 not null comment '状态: 0禁用 1启用',
    created_at datetime   default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint `key`
        unique (`key`),
    constraint uniq_role_uid
        unique (role_uid)
)
    comment '角色信息表';

create index idx_name
    on t_role_info (name);

create index idx_pid
    on t_role_info (pid);

create index idx_pid_sort
    on t_role_info (pid, sort);

create index idx_status
    on t_role_info (status);

create table t_role_menu
(
    role_id    bigint                             not null comment '角色ID',
    menu_id    bigint                             not null comment '菜单ID',
    created_at datetime default CURRENT_TIMESTAMP null,
    primary key (role_id, menu_id)
)
    comment '角色菜单关联表';

create index idx_created_at
    on t_role_menu (created_at);

create table t_sys_config
(
    id          bigint auto_increment comment '配置ID'
        primary key,
    `group`     varchar(128) default 'default'         not null comment '配置分组',
    `key`       varchar(100)                           not null comment '参数键名',
    value       varchar(2000)                          not null comment '参数值',
    description varchar(500)                           null,
    sort        int          default 0                 not null comment '排序',
    status      tinyint(1)   default 1                 not null comment '状态: 0禁用 1启用',
    created_at  datetime     default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  datetime     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uniq_key
        unique (`key`)
)
    comment '系统配置表';

create index idx_group
    on t_sys_config (`group`);

create index idx_group_status
    on t_sys_config (`group`, status);

create index idx_status
    on t_sys_config (status);


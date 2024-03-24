SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS `rooms` (
  `room_id` bigint(20) NOT NULL COMMENT '房间号',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '登记者',
  `create_time` datetime DEFAULT NULL COMMENT '登记时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '使用者',
  `update_time` datetime DEFAULT NULL COMMENT '使用时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8 COMMENT='房间表';

BEGIN;
INSERT INTO `rooms` VALUES (001, '王泽林', '12345678910', '0', '0', 'admin', '2024-01-01 00:00:01', '张建宇', '2024-01-01 00:00:01','');
INSERT INTO `rooms` VALUES (002, '王泽林', '12345678910', '0', '0', 'admin', '2024-01-01 00:00:01', '张建宇', '2024-01-01 00:00:01','');
COMMIT;

CREATE TABLE  IF NOT EXISTS `task` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `job_detail` varchar(500) NOT NULL COMMENT '任务详情',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='员工任务进度表';

BEGIN;
INSERT INTO `task` VALUES (001, '季度KPI', '前台', '完成999万单订单', '0', 'admin', '2024-01-01 00:00:01', '', NULL, '已完成500万');
COMMIT;

CREATE TABLE IF NOT EXISTS `login_infor` (
  `info_id` bigint(20) NOT NULL auto_increment COMMENT '访问ID',
  `user_name` varchar(50) DEFAULT '' COMMENT '用户账号',
  `status` char(1) DEFAULT '1' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='系统访问记录';

BEGIN;
COMMIT;

CREATE TABLE IF NOT EXISTS `hotel_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='岗位信息表';

BEGIN;
INSERT INTO `hotel_post` VALUES (1, 'ceo', '董事长', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_post` VALUES (2, 'leader', '总经理', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_post` VALUES (3, 'money', '会计', 2, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_post` VALUES (4, 'user_996', '行政', 3, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_post` VALUES (5, 'user_007', '前台员工', 4, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_post` VALUES (6, 'clean', '保洁', 4, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
COMMIT;

CREATE TABLE IF NOT EXISTS `hotel_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `role_passwd` varchar(30) NOT NULL COMMENT '角色密码',
  `phone` char(30) NOT NULL COMMENT '电话号/微信号',
  `role_key` varchar(100) NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：会计数据权限 3：行政数据权限 4：前台及以下数据权限）',
  `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色信息表';

BEGIN;
INSERT INTO `hotel_role` VALUES (1, '王泽林', 'zl20040203','12345678910','ceo', 1, '1', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_role` VALUES (2, '王泽林', 'zl20040203', '12345678910', 'leader', 2, '1', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_role` VALUES (3, '王泽林', 'zl20040203', '12345678910', 'money', 3, '2', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_role` VALUES (4, '王泽林', 'zl20040203', '12345678910', 'user_996', 4, '3', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_role` VALUES (5, '王泽林', 'zl20040203', '12345678910', 'user_007', 5, '4', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `hotel_role` VALUES (6, '王泽林', 'zl20040203', '12345678910', 'clean', 6, '4', 1, '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
COMMIT;

CREATE TABLE IF NOT EXISTS `sleep_user` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `phone` char(30) NOT NULL COMMENT '电话号/微信号',
   `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='登记客户信息表';

BEGIN;
INSERT INTO `sleep_user` VALUES (1, '王泽林', '12345678910', '0', '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
INSERT INTO `sleep_user` VALUES (2, '王泽林', '12345678910', '0', '0', 'admin', '2024-01-01 00:00:01', '', NULL, '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE  IF NOT EXISTS `menu` (
    `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账单ID',
    `serve_name` varchar(64) NOT NULL DEFAULT '' COMMENT '服务名称',
    `room` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '房间名称',
    `money` varchar(500) NOT NULL COMMENT '消费详情',
    `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1完成）',
    `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
    `create_time` datetime DEFAULT NULL COMMENT '创建时间',
    `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
    `update_time` datetime DEFAULT NULL COMMENT '更新时间',
    `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
    PRIMARY KEY (`menu_id`,`serve_name`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='客户消费表';

BEGIN;
INSERT INTO `menu` VALUES (114514, '288精油开背', '8888', '288', '1', 'admin', '2024-01-01 00:00:01', '', NULL, '已完成');
COMMIT;

CREATE TABLE  IF NOT EXISTS hotel_update (
    `update_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账单ID',
    `serve_name` varchar(64) NOT NULL DEFAULT '' COMMENT '服务名称',
    `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1完成）',
    `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
    `create_time` datetime DEFAULT NULL COMMENT '创建时间',
    `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
    `update_time` datetime DEFAULT NULL COMMENT '更新时间',
    `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
    PRIMARY KEY (`update_id`,`serve_name`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='更新表';

BEGIN;
INSERT INTO hotel_update VALUES (null, '删除客户', '1','admin', '2024-01-01 00:00:01', '', NULL, '已完成');
COMMIT;
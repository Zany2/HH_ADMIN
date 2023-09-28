/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MariaDB
 Source Server Version : 100521
 Source Host           : 127.0.0.1:3306
 Source Schema         : hh_admin

 Target Server Type    : MariaDB
 Target Server Version : 100521
 File Encoding         : 65001

 Date: 17/09/2023 20:30:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (10, 'g', '11111111111', 'admin_post', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'admin_post', '/api/v1/system/auth/post/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', 'admin_post', '/api/v1/system/auth/post/delete', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', 'admin_post', '/api/v1/system/auth/post/edit', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', 'admin_post', '/api/v1/system/auth/post/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'admin_post', '/api/v1/system/auth/post/query', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1, 'p', 'admin_post', '/system/auth', 'GET', '', '', '');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '部门状态（1正常 2停用）',
  `del_flag` tinyint(1) NOT NULL DEFAULT 1 COMMENT '删除标志（1代表存在 2代表删除）',
  `create_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 205 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, 0, '0', '行行科技', 0, 'daixk', '13888888888', 'aa@qq.com', 1, 1, '1', '2023-09-08 17:46:02', '0', NULL);
INSERT INTO `sys_dept` VALUES (2, 1, '0,1', '青岛总公司', 1, 'daixk', '13888888888', 'aa@qq.com', 1, 1, '1', '2023-09-08 17:46:02', '0', NULL);
INSERT INTO `sys_dept` VALUES (3, 1, '0,1', '北京总公司', 1, 'daixk', '13888888888', 'aa@qq.com', 1, 1, '1', '2023-09-08 17:46:02', '0', NULL);
INSERT INTO `sys_dept` VALUES (4, 2, '0,1,2', '研发部门', 1, 'daixk', '13888888888', 'aa@qq.com', 1, 1, '1', '2023-09-08 17:46:02', '0', NULL);
INSERT INTO `sys_dept` VALUES (5, 2, '0,1,2', '市场部门', 2, 'daixk', '13888888888', 'aa@qq.com', 1, 1, '1', '2023-09-08 17:46:02', '0', NULL);

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, '13864897335', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-16 23:38:15', '系统后台');
INSERT INTO `sys_login_log` VALUES (2, '13888888888', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-16 23:43:49', '系统后台');
INSERT INTO `sys_login_log` VALUES (3, '13864897335', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-16 23:44:25', '系统后台');
INSERT INTO `sys_login_log` VALUES (4, '13888888888', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-17 00:31:07', '系统后台');
INSERT INTO `sys_login_log` VALUES (5, '13888888888', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-17 13:23:24', '系统后台');
INSERT INTO `sys_login_log` VALUES (6, '13888888888', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-17 13:23:50', '系统后台');
INSERT INTO `sys_login_log` VALUES (7, '13888888888', '127.0.0.1', '内网IP', 'PostmanRuntime-ApipostRuntime/1.1.0', '', 0, '登录成功', '2023-09-17 19:43:34', '系统后台');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父菜单ID 默认0',
  `order_num` int(4) NOT NULL DEFAULT 0 COMMENT '显示顺序 默认0',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_frame` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否为外链 1否 2是 默认1',
  `is_cache` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否缓存 1缓存 2不缓存 默认1',
  `menu_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单类型（1目录 2菜单 3按钮）',
  `visible` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单状态 1显示 2隐藏 默认1',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单状态 1正常 2停用 默认1',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#' COMMENT '菜单图标 默认\'#\'',
  `create_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建者 默认1',
  `update_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者 默认1',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '权限管理', 0, 1, '/api/v1/system/auth', '', '', 1, 1, 1, 1, 1, '/system/auth:GET', '#', '1', '1', '', '2023-09-07 23:34:20', '2023-09-07 23:34:23');
INSERT INTO `sys_menu` VALUES (2, '系统监控', 0, 1, '/api/v1/system/monitor', '', '', 1, 1, 1, 1, 1, '/api/v1/system/monitor:GET', '#', '1', '1', '', '2023-09-07 23:36:14', '2023-09-07 23:34:27');
INSERT INTO `sys_menu` VALUES (3, '菜单管理', 1, 2, '/api/v1/system/auth/menu/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/auth/menu/list:GET', '#', '1', '1', '', '2023-09-07 23:36:16', '2023-09-07 23:36:18');
INSERT INTO `sys_menu` VALUES (4, '角色管理', 1, 2, '/api/v1/system/auth/role/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/auth/role/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (5, '部门管理', 1, 2, '/api/v1/system/auth/dept/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/auth/dept/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (6, '岗位管理', 1, 2, '/api/v1/system/auth/post/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/auth/post/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (7, '用户管理', 1, 2, '/api/v1/system/auth/user/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/auth/user/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (8, '服务监控', 2, 2, '/api/v1/system/monitor/server/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/monitor/server/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (9, '登录日志', 2, 2, '/api/v1/system/monitor/login/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/monitor/login/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (10, '操作日志', 2, 2, '/api/v1/system/monitor/oper/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/monitor/oper/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (11, '在线用户', 2, 2, '/api/v1/system/monitor/online/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/monitor/online/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');
INSERT INTO `sys_menu` VALUES (12, '角色查询', 4, 1, '/api/v1/system/auth/role/query', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/role/query:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (13, '角色新增', 4, 2, '/api/v1/system/auth/role/add', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/role/add:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (14, '角色修改', 4, 3, '/api/v1/system/auth/role/edit', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/role/edit:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (15, '角色删除', 4, 4, '/api/v1/system/auth/role/delete', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/role/delete:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (16, '部门查询', 5, 1, '/api/v1/system/auth/dept/query', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/dept/query:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (17, '部门新增', 5, 2, '/api/v1/system/auth/dept/add', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/dept/add:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (18, '部门修改', 5, 3, '/api/v1/system/auth/dept/edit', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/dept/edit:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (19, '部门删除', 5, 4, '/api/v1/system/auth/dept/delete', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/dept/delete:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (20, '岗位查询', 6, 1, '/api/v1/system/auth/post/query', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/post/query:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (21, '岗位新增', 6, 2, '/api/v1/system/auth/post/add', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/post/add:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (22, '岗位修改', 6, 3, '/api/v1/system/auth/post/edit', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/post/edit:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (23, '岗位删除', 6, 4, '/api/v1/system/auth/post/delete', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/post/delete:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (24, '用户查询', 5, 1, '/api/v1/system/auth/user/query', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/user/query:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (25, '用户新增', 5, 2, '/api/v1/system/auth/user/add', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/user/add:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (26, '用户修改', 5, 3, '/api/v1/system/auth/user/edit', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/user/edit:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (27, '用户删除', 5, 4, '/api/v1/system/auth/user/delete', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/user/delete:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (28, '菜单查询', 3, 1, '/api/v1/system/auth/menu/query', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/menu/query:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (29, '菜单新增', 3, 2, '/api/v1/system/auth/menu/add', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/menu/add:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (30, '菜单修改', 3, 3, '/api/v1/system/auth/menu/edit', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/menu/edit:POST', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (31, '菜单删除', 3, 4, '/api/v1/system/auth/menu/delete', '', '', 1, 1, 3, 1, 1, '/api/v1/system/auth/menu/delete:GET', '#', '1', '1', '', '2023-09-10 01:18:52', '2023-09-10 01:18:54');
INSERT INTO `sys_menu` VALUES (32, '服务监控查询', 2, 2, '/api/v1/system/monitor/server/list', '', '', 1, 1, 2, 1, 1, '/api/v1/system/monitor/server/list:GET', '#', '1', '1', '', '2023-09-07 23:37:13', '2023-09-07 23:37:15');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位编码',
  `post_name_a` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `post_sort` int(11) NOT NULL COMMENT '显示顺序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态（1正常 2停用）',
  `del_flag` tinyint(4) NOT NULL DEFAULT 1 COMMENT '删除标志 1正常 2删除 默认1',
  `create_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, 1, 1, '1', '2023-09-09 02:16:31', '13888888888', '2023-09-09 02:16:33', '');
INSERT INTO `sys_post` VALUES (11, 'fceo', '', 2, 1, 1, '13888888888', '2023-09-17 20:27:22', '', NULL, '');
INSERT INTO `sys_post` VALUES (12, 'fceo', '', 2, 1, 1, '13888888888', '2023-09-17 20:28:39', '', NULL, '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `data_scope` tinyint(1) NOT NULL DEFAULT 2 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）默认2',
  `menu_check_strictly` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) NOT NULL DEFAULT 1 COMMENT '部门树选择项是否关联显示',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '角色状态 1正常 2停用 默认1',
  `del_flag` tinyint(1) NOT NULL DEFAULT 1 COMMENT '删除标志 1正常 2停用 默认1',
  `create_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建者 默认0',
  `update_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者 默认0',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'admin', 1, 1, 1, 1, 1, 1, '', '', '超级管理员', '2023-09-16 21:43:40', NULL);
INSERT INTO `sys_role` VALUES (2, '岗位管理员', 'admin_post', 2, 2, 1, 1, 1, 1, '', '', '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (2, 1);
INSERT INTO `sys_role_menu` VALUES (2, 6);
INSERT INTO `sys_role_menu` VALUES (2, 20);
INSERT INTO `sys_role_menu` VALUES (2, 21);
INSERT INTO `sys_role_menu` VALUES (2, 22);
INSERT INTO `sys_role_menu` VALUES (2, 23);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `dept_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '部门ID',
  `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_type` tinyint(1) NOT NULL DEFAULT 2 COMMENT '用户类型 1超级管理员 2普通用户 默认2',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(1) NOT NULL DEFAULT 1 COMMENT '用户性别 1男 2女 3未知 默认1',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '帐号状态 1正常 2停用 默认1',
  `del_flag` tinyint(1) NOT NULL DEFAULT 1 COMMENT '删除标志 1正常 2删除 默认1',
  `login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `create_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建者 默认1',
  `update_by` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新者 默认1',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `login_date` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1, 'admin', 1, '', '13888888888', 1, '', 1, 1, '', '1', '1', '', NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (18, 'test1', '5a105e8b9d40e1329780d62ea2265d8a', 1, 'test01', 2, 'aaa@qq.com', '11111111111', 1, '', 1, 2, '', '13888888888', '13888888888', '这是一个测试', NULL, '2023-09-17 00:46:21', '2023-09-17 00:46:21');

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 1);
INSERT INTO `sys_user_post` VALUES (18, 1);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (18, 2);

SET FOREIGN_KEY_CHECKS = 1;

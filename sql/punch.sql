/*
 Navicat Premium Data Transfer

 Source Server         : 本地连接
 Source Server Type    : MySQL
 Source Server Version : 50725 (5.7.25)
 Source Host           : localhost:3307
 Source Schema         : punch

 Target Server Type    : MySQL
 Target Server Version : 50725 (5.7.25)
 File Encoding         : 65001

 Date: 13/09/2022 09:19:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for plan
-- ----------------------------
DROP TABLE IF EXISTS `plan`;
CREATE TABLE `plan` (
  `id` int(11) NOT NULL COMMENT '主键',
  `plan_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '计划名称',
  `plan_desc` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '计划描述',
  `corn_time` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'corn表达式时间，用于定时触发某个特性',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '当前状态: 1- 正在进行中，2-已经结束；99-删除',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  `update_tme` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of plan
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for punch_record
-- ----------------------------
DROP TABLE IF EXISTS `punch_record`;
CREATE TABLE `punch_record` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `plan_id` int(11) NOT NULL COMMENT '计划id',
  `punch_time` bigint(20) NOT NULL COMMENT '打开时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of punch_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` int(11) NOT NULL COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '用户名',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '密码',
  `last_login_time` bigint(20) NOT NULL COMMENT '上一次登录时间',
  `qq_token` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'QQ登录所需要的token',
  `weixin_token` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '微信登录所需要的token',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user_info
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

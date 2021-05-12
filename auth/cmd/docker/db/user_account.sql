/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3306
 Source Schema         : comic

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 27/04/2021 21:00:28
*/

CREATE DATABASE `comic` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';
USE comic;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account`  (
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态是否正常(1=正常,0=异常)',
  `add_time` bigint(0) NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_user_name`(`user_name`) USING BTREE,
  INDEX `index_union`(`user_name`, `password`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 120 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_account
-- ----------------------------
INSERT INTO `user_account` VALUES ('qwe', '$2a$10$sO7VFzvcmb8BbCHO/8YXNeKh8.z6xpmT9vIAU1FKsxDjpivzEubbG', 118, 1, 1618755903);
INSERT INTO `user_account` VALUES ('chengche123', '$2a$10$JOw.1kQTIMeNaleW45qSOuhbetyKwjqUDbtXzhm46a5OiSpl8U24K', 121, 1, 1619528396);

SET FOREIGN_KEY_CHECKS = 1;

/*
 Navicat Premium Data Transfer

 Source Server         : cx
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 47.96.152.251:3306
 Source Schema         : image_acquisition

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 05/07/2021 20:28:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for patient
-- ----------------------------
DROP TABLE IF EXISTS `patient`;
CREATE TABLE `patient`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `sex` tinyint(4) NULL DEFAULT NULL,
  `id_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `eye_position` tinyint(4) NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `image` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `f1`(`user_id`) USING BTREE,
  CONSTRAINT `f1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (13, 'qwe', '$2a$10$h/JjbDdbeUKWzHBsMzpR8OYjJDX09yyf2lkIL0JhC4BB4kTcJ4wV6');
INSERT INTO `user` VALUES (14, 'gandaojiali1', '$2a$10$WwMyxoH.CG/obMpJkOuBWulaitQg.aI.JhxZEeM9P4DRX1cE/yV7e');
INSERT INTO `user` VALUES (15, 'gandaojiali2', '$2a$10$i3jmqnuybAhSsy0J/Cl7KeiS6ByOSTIu12cjwEfvP1Z7GhPnod5a.');
INSERT INTO `user` VALUES (16, 'abcd', '$2a$10$7E1LYjbeo.qrz6ddp6sAAOw3gDTdms7yOXfgrRDUEgrvEs5dOS2/q');
INSERT INTO `user` VALUES (17, 'abcde', '$2a$10$xRh.Zb6DKvp1EBXuigI1te7uLqclgRfXTyP8tNyRx8b8JuTXXGMW2');
INSERT INTO `user` VALUES (18, '114488', '$2a$10$s8MQDJMoTbRCNJdxVs/YQeMGEPvl8YUNlxCftaZYM0uEdnMCPH87O');
INSERT INTO `user` VALUES (19, '777', '$2a$10$5AVJAwAvzJQSg5POwdLQRuYzrsbIr05Eq65vvdUnSydFalsFuHAFW');
INSERT INTO `user` VALUES (20, '111', '$2a$10$a1wTSt4vrvL7/kJOsDGyL.oKV1k1B23e0XpQ.E2Tjna7kRC5KszH2');
INSERT INTO `user` VALUES (21, 'cx', '$2a$10$wsuRxFC858/apqCtDZYpneONDMdFJEdPBX0zk1/l4yera.698BOKu');
INSERT INTO `user` VALUES (22, '22', '$2a$10$3KYtmgsqGTjlHBllJCHBoelhXjWGs1nt4TCXewWO1VV6/AKkiMFW2');
INSERT INTO `user` VALUES (23, '1111', '$2a$10$2Ex8dlx.p4pKdYfAW5QP1eqnX11/a2AX5VIW.rx/8XSYBKMAdVsZi');

SET FOREIGN_KEY_CHECKS = 1;

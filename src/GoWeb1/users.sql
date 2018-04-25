/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 100121
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 100121
 File Encoding         : 65001

 Date: 26/04/2018 00:05:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `userid` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `age` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`userid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (51, 'Trần Hoài Nam', 11);
INSERT INTO `users` VALUES (52, 'Trần Hoài Nam', 11);
INSERT INTO `users` VALUES (53, 'Trần Hoài Nam', 11);
INSERT INTO `users` VALUES (54, 'Trần Hoài Nam', 11);
INSERT INTO `users` VALUES (59, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (65, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (66, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (70, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (71, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (72, 'Nguyễn Ngọc Hoàng', 26);
INSERT INTO `users` VALUES (73, 'Văn Mai Hương', 88);
INSERT INTO `users` VALUES (74, 'Văn Mai Hương', 88);

SET FOREIGN_KEY_CHECKS = 1;

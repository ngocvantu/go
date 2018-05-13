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

 Date: 13/05/2018 15:58:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `CONTENT` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `DATECREATED` date NOT NULL,
  `ISCOMPLETE` tinyint(1) NOT NULL,
  `STARTTIME` time(0) NOT NULL,
  `ENDTIME` time(0) NOT NULL,
  `status` int(6) NOT NULL DEFAULT 0,
  `timetake` time(0) NOT NULL DEFAULT '00:00:00',
  `pause` time(0) NOT NULL,
  `resume` time(0) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;

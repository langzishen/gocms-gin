/*
 Navicat Premium Data Transfer

 Source Server         : vm-ubuntu22.04
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : 192.168.244.128:3306
 Source Schema         : gocms_gin

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 01/08/2022 17:51:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access`  (
  `role_id` int UNSIGNED NOT NULL,
  `node_id` int UNSIGNED NOT NULL,
  `level` tinyint(1) NOT NULL,
  `pid` int NOT NULL,
  `module` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  INDEX `groupId`(`role_id`) USING BTREE,
  INDEX `nodeId`(`node_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of access
-- ----------------------------
INSERT INTO `access` VALUES (3, 1, 1, 0, '');
INSERT INTO `access` VALUES (3, 284, 0, 283, '');
INSERT INTO `access` VALUES (3, 285, 0, 283, '');
INSERT INTO `access` VALUES (3, 286, 0, 283, '');
INSERT INTO `access` VALUES (3, 287, 0, 283, '');
INSERT INTO `access` VALUES (3, 289, 0, 283, '');
INSERT INTO `access` VALUES (3, 290, 0, 283, '');
INSERT INTO `access` VALUES (3, 291, 0, 283, '');
INSERT INTO `access` VALUES (3, 292, 0, 283, '');
INSERT INTO `access` VALUES (3, 293, 0, 283, '');
INSERT INTO `access` VALUES (3, 294, 0, 283, '');
INSERT INTO `access` VALUES (3, 295, 0, 283, '');
INSERT INTO `access` VALUES (3, 172, 0, 4, '');
INSERT INTO `access` VALUES (3, 173, 0, 4, '');
INSERT INTO `access` VALUES (3, 174, 0, 4, '');
INSERT INTO `access` VALUES (3, 175, 0, 4, '');
INSERT INTO `access` VALUES (3, 176, 0, 4, '');
INSERT INTO `access` VALUES (3, 177, 0, 4, '');
INSERT INTO `access` VALUES (3, 178, 0, 4, '');
INSERT INTO `access` VALUES (3, 179, 0, 4, '');
INSERT INTO `access` VALUES (3, 180, 0, 4, '');
INSERT INTO `access` VALUES (3, 181, 0, 4, '');
INSERT INTO `access` VALUES (3, 182, 0, 4, '');
INSERT INTO `access` VALUES (3, 183, 0, 4, '');
INSERT INTO `access` VALUES (3, 202, 0, 28, '');
INSERT INTO `access` VALUES (3, 203, 0, 28, '');
INSERT INTO `access` VALUES (3, 204, 0, 28, '');
INSERT INTO `access` VALUES (3, 205, 0, 28, '');
INSERT INTO `access` VALUES (3, 206, 0, 28, '');
INSERT INTO `access` VALUES (3, 207, 0, 28, '');
INSERT INTO `access` VALUES (3, 208, 0, 28, '');
INSERT INTO `access` VALUES (3, 209, 0, 28, '');
INSERT INTO `access` VALUES (3, 210, 0, 28, '');
INSERT INTO `access` VALUES (3, 211, 0, 28, '');
INSERT INTO `access` VALUES (3, 212, 0, 28, '');
INSERT INTO `access` VALUES (3, 213, 0, 28, '');
INSERT INTO `access` VALUES (3, 238, 0, 112, '');
INSERT INTO `access` VALUES (3, 239, 0, 112, '');
INSERT INTO `access` VALUES (3, 240, 0, 112, '');
INSERT INTO `access` VALUES (3, 241, 0, 112, '');
INSERT INTO `access` VALUES (3, 242, 0, 112, '');
INSERT INTO `access` VALUES (3, 243, 0, 112, '');
INSERT INTO `access` VALUES (3, 244, 0, 112, '');
INSERT INTO `access` VALUES (3, 245, 0, 112, '');
INSERT INTO `access` VALUES (3, 246, 0, 112, '');
INSERT INTO `access` VALUES (3, 247, 0, 112, '');
INSERT INTO `access` VALUES (3, 248, 0, 112, '');
INSERT INTO `access` VALUES (3, 249, 0, 112, '');
INSERT INTO `access` VALUES (3, 82, 0, 1, '');
INSERT INTO `access` VALUES (3, 80, 0, 5, '');
INSERT INTO `access` VALUES (3, 81, 0, 5, '');
INSERT INTO `access` VALUES (3, 7, 0, 6, '');
INSERT INTO `access` VALUES (3, 288, 0, 283, '');
INSERT INTO `access` VALUES (4, 1, 1, 0, '');
INSERT INTO `access` VALUES (4, 112, 2, 1, '');
INSERT INTO `access` VALUES (4, 5, 2, 1, '');
INSERT INTO `access` VALUES (4, 248, 0, 112, '');
INSERT INTO `access` VALUES (4, 81, 0, 5, '');
INSERT INTO `access` VALUES (3, 4, 2, 1, '');
INSERT INTO `access` VALUES (3, 5, 2, 1, '');
INSERT INTO `access` VALUES (3, 6, 2, 1, '');
INSERT INTO `access` VALUES (3, 28, 2, 1, '');
INSERT INTO `access` VALUES (3, 112, 2, 1, '');
INSERT INTO `access` VALUES (3, 283, 2, 1, '');
INSERT INTO `access` VALUES (3, 1, 1, 0, '');

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tid` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '????????????',
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '????????????',
  `keywords` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '?????????',
  `description` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '??????',
  `img` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `content` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '??????',
  `create_time` int NOT NULL COMMENT '????????????',
  `update_time` int NOT NULL COMMENT '????????????',
  `creater_id` int UNSIGNED NOT NULL COMMENT '?????????',
  `sort` smallint NOT NULL DEFAULT 0 COMMENT '?????????',
  `apv` smallint NOT NULL DEFAULT 0 COMMENT '?????????',
  `rewrite` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'URL?????????',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '?????? 1:??????,0:??????',
  `template` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `attrtj` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `outurl` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `is_login_show` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '???????????????????????????',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `tid`(`tid`) USING BTREE,
  INDEX `title`(`title`) USING BTREE,
  FULLTEXT INDEX `content`(`content`)
) ENGINE = MyISAM AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '92', '??????????????????1', '??????????????????1', '??????????????????1', '/static/uploads/2022-6-18/3056da2c-b4c4-9358-d490-c0d032886e94.png', '<p>??????????????????1</p>', 0, 1656059958, 1, 999, 0, '', 1, '', '1,2,3', '', 0);
INSERT INTO `article` VALUES (2, '92', '??????2??????1', '??????2??????1', '??????2??????1', '/static/uploads/2022-6-18/85e6fa43-75d8-69b0-0c31-a34adc599617.png', '<p>??????2??????1</p>', 1655523855, 1656059963, 1, 999, 0, '', 1, '', '1', '', 0);
INSERT INTO `article` VALUES (3, '93', '??????3??????2', '??????3??????2', '??????3??????2', '/static/uploads/2022-6-21/41898ddc-bcb6-6fe8-2489-598594bf3031.png', '<p>??????3??????2??????3??????2??????3??????2??????3??????2??????3??????2</p>', 1655792951, 1656059968, 1, 999, 0, '', 1, '', '1', '', 1);
INSERT INTO `article` VALUES (4, '93', '??????4??????2', '??????4??????2', '??????4??????2', '/static/uploads/2022-6-21/f387048d-df9e-4421-3a48-68283386737a.png', '<p>??????4??????2??????4??????2??????4??????2??????4??????2??????4??????2</p>', 1655793202, 1655795196, 1, 999, 0, '', 1, '', '1,2,3', '', 0);

-- ----------------------------
-- Table structure for attribute
-- ----------------------------
DROP TABLE IF EXISTS `attribute`;
CREATE TABLE `attribute`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `model_ename` varchar(80) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '??????????????????',
  `attrname` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '????????????',
  `attrvalue` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '?????????',
  `sort` mediumint NOT NULL DEFAULT 999 COMMENT '?????????',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '?????? 1:??????,0:??????',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of attribute
-- ----------------------------
INSERT INTO `attribute` VALUES (1, 'Product', '????????????', '1', 1, 1);
INSERT INTO `attribute` VALUES (2, 'Product', '????????????', '2', 2, 1);
INSERT INTO `attribute` VALUES (3, 'Product', '????????????', '3', 98, 1);
INSERT INTO `attribute` VALUES (4, 'Product', '??????', '4', 97, 1);
INSERT INTO `attribute` VALUES (10, 'Product', '????????????', '1', 99, 1);
INSERT INTO `attribute` VALUES (12, 'Product', '??????', '1', 99, 1);
INSERT INTO `attribute` VALUES (11, 'Product', '?????????', '1', 99, 1);
INSERT INTO `attribute` VALUES (13, 'Product', '????????????', '1', 99, 1);
INSERT INTO `attribute` VALUES (17, 'article', '????????????', '1', 99, 1);
INSERT INTO `attribute` VALUES (18, 'article', '????????????', '2', 99, 1);
INSERT INTO `attribute` VALUES (19, 'article', '????????????', '3', 99, 1);
INSERT INTO `attribute` VALUES (20, 'article', '??????', '4', 99, 1);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `classid` int NOT NULL AUTO_INCREMENT COMMENT '??????id',
  `classpid` int NOT NULL DEFAULT 0 COMMENT '?????????id',
  `classpids` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '?????????ids',
  `classchild` tinyint(1) NOT NULL DEFAULT 0 COMMENT '???????????????',
  `classchildids` varchar(2000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????ids',
  `classarrchildids` mediumtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '??????????????????',
  `classtitle` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '????????????',
  `classsort` smallint NOT NULL DEFAULT 0 COMMENT '??????',
  `classstatus` tinyint(1) NOT NULL DEFAULT 1 COMMENT '??????',
  `classkeywords` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '?????????',
  `classdescription` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '??????',
  `classmodule` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `classrewrite` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'URL?????????',
  `classtemplate` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `newstemplate` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '????????????',
  `classimg` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '???????????????',
  `classapv` int NOT NULL DEFAULT 0 COMMENT '???????????????',
  `classdomain` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '??????????????????',
  `classouturl` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '??????????????????',
  `classmenushow` tinyint(1) NOT NULL DEFAULT 1 COMMENT '?????????????????????',
  PRIMARY KEY (`classid`) USING BTREE,
  INDEX `pid`(`classpid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 94 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (38, 0, '', 1, '38', 'a:1:{i:0;s:2:\"38\";}', 'banner', 0, 1, '', '', 'Photo', 'tupian', '', '', NULL, 0, '', '', 1);
INSERT INTO `category` VALUES (46, 0, '', 1, '46', 'a:1:{i:0;s:2:\"46\";}', '????????????', 999, 1, '', '', 'About', NULL, NULL, NULL, NULL, 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (47, 0, '', 0, '47', 'a:1:{i:0;s:2:\"47\";}', '????????????', 999, 1, '', '', 'About', NULL, NULL, NULL, '/static/uploads/2022-6-15/8f4a2781-c09e-3ad1-830c-58d2db80b157.png', 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (48, 0, '', 0, '48', 'a:1:{i:0;s:2:\"48\";}', '????????????', 999, 1, '', '', 'About', NULL, NULL, NULL, NULL, 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (49, 0, '', 0, '49', 'a:1:{i:0;s:2:\"49\";}', '????????????', 999, 1, '', '', 'About', NULL, NULL, NULL, NULL, 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (50, 0, '0,0', 1, '50,73,75', 'a:3:{i:0;s:2:\"50\";i:1;a:2:{i:0;a:4:{s:10:\"classchild\";i:1;s:7:\"classid\";i:73;s:11:\"classmodule\";s:5:\"About\";s:10:\"classtitle\";s:4:\"jhkh\";}i:1;a:1:{i:0;s:2:\"73\";}}i:2;a:2:{i:0;a:4:{s:10:\"classchild\";i:0;s:7:\"classid\";i:75;s:11:\"classmodule\";s:5:\"About\";s:10:\"classtitle\";s:4:\"nnnn\";}i:1;a:1:{i:0;s:2:\"75\";}}}', '????????????', 999, 1, '', '', 'About', NULL, NULL, NULL, NULL, 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (73, 50, '0,0,50', 1, '73', 'a:1:{i:0;s:2:\"73\";}', 'jhkh', 999, 1, '', '', 'About', NULL, NULL, NULL, '/static/uploads/2022-6-15/a8740f89-ca50-8556-67d2-40d9c5f0a510.png', 0, NULL, NULL, 1);
INSERT INTO `category` VALUES (93, 92, '', 0, '', '', '??????????????????', 999, 1, '??????????????????', '??????????????????', 'article', '', '', '', '/static/uploads/2022-6-18/707f18d8-1dfc-117f-85aa-b17d8c8e8e22.png', 0, '', '', 1);
INSERT INTO `category` VALUES (91, 73, '', 0, '', '', 'mmmmmm', 999, 1, '3131', '97987', 'Product', '', '', '', '', 0, '', '', 1);
INSERT INTO `category` VALUES (90, 46, '', 0, '', '', '131', 999, 1, '13131', '6454', 'About', '', '', '', '', 0, '', '', 1);
INSERT INTO `category` VALUES (92, 0, '', 0, '', '', '??????????????????', 999, 1, '??????????????????', '??????????????????', 'article', '', '', '', '/static/uploads/2022-6-18/ef8c84e2-8015-26a8-3328-982beb1e4b0e.png', 0, '', '', 1);
INSERT INTO `category` VALUES (89, 38, '', 0, '', '', '131', 999, 1, '6464', '9879', 'Product', '', '', '', '/static/uploads/2022-6-15/e4e55c42-0ae4-d12e-50e0-b2fc042b0867.png', 0, '', '', 1);
INSERT INTO `category` VALUES (87, 38, '', 0, '', '', '2?????????', 999, 1, '2', '2', 'Product', '', '', '', '', 0, '', '', 1);

-- ----------------------------
-- Table structure for data_access
-- ----------------------------
DROP TABLE IF EXISTS `data_access`;
CREATE TABLE `data_access`  (
  `role_id` int UNSIGNED NOT NULL COMMENT '?????????ID',
  `node_id` int UNSIGNED NOT NULL COMMENT '??????ID',
  `tid` int UNSIGNED NOT NULL COMMENT '??????ID',
  `model` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '????????????',
  `create_time` int UNSIGNED NOT NULL COMMENT '????????????'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '???????????????' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of data_access
-- ----------------------------
INSERT INTO `data_access` VALUES (3, 143, 93, 'article', 1656571030);
INSERT INTO `data_access` VALUES (3, 142, 92, 'article', 1656571663);

-- ----------------------------
-- Table structure for log_login
-- ----------------------------
DROP TABLE IF EXISTS `log_login`;
CREATE TABLE `log_login`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '????????????',
  `user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '??????ID',
  `last_login_time` int UNSIGNED NOT NULL COMMENT '?????????????????????',
  `login_app` varchar(80) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '?????????????????????',
  `from` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '?????????1???web2???android3???ios',
  `last_login_ip` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '???????????????Ip',
  `device_token` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '???????????????????????????',
  `count` int UNSIGNED NOT NULL COMMENT '???????????????',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 149 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '???????????????' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_login
-- ----------------------------
INSERT INTO `log_login` VALUES (1, 1, 1652670362, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (2, 1, 1652670510, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (3, 1, 1652672233, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (4, 1, 1652680997, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (5, 1, 1652774627, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (6, 1, 1653886419, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (7, 1, 1653886434, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (8, 1, 1653887005, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (9, 1, 1653887142, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (10, 1, 1653887437, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (11, 1, 1653887580, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (12, 1, 1653887747, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (13, 1, 1653887863, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (14, 1, 1653887951, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (15, 1, 1653888060, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (16, 1, 1653888177, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (17, 1, 1653888336, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (18, 1, 1653888352, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (19, 1, 1653890358, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (20, 1, 1653890365, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (21, 1, 1653890376, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (22, 1, 1653890715, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (23, 1, 1653890726, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (24, 1, 1653890727, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (25, 1, 1653890752, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (26, 1, 1653891435, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (27, 1, 1655370704, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (28, 1, 1655894924, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (29, 1, 1656123708, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (30, 2, 1656396074, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (31, 1, 1656396836, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (32, 2, 1656397078, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (33, 11, 1656397496, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (34, 1, 1656398993, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (35, 11, 1656399012, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (36, 1, 1656403320, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (37, 11, 1656403470, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (38, 11, 1656404416, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (39, 1, 1656404441, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (40, 11, 1656404473, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (41, 11, 1656404484, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (42, 1, 1656404645, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (43, 11, 1656404716, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (44, 11, 1656405785, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (45, 11, 1656406014, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (46, 11, 1656407194, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (47, 11, 1656407194, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (48, 11, 1656407194, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (49, 11, 1656407194, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (50, 1, 1656408613, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (51, 1, 1658415676, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (52, 1, 1658836414, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (53, 1, 1658836431, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (54, 1, 1658837034, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (55, 1, 1658837174, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (56, 1, 1658837373, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (57, 1, 1658837646, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (58, 1, 1658841340, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (59, 1, 1658855602, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (60, 1, 1658856067, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (61, 1, 1658856098, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (62, 1, 1659254789, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (63, 1, 1659255302, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (64, 1, 1659255725, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (65, 1, 1659255752, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (66, 1, 1659257231, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (67, 1, 1659257243, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (68, 1, 1659257247, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (69, 1, 1659257261, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (70, 1, 1659257273, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (71, 1, 1659257944, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (72, 1, 1659257949, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (73, 1, 1659258190, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (74, 1, 1659258287, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (75, 1, 1659259904, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (76, 1, 1659261184, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (77, 1, 1659262805, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (78, 1, 1659265586, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (79, 1, 1659265808, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (80, 1, 1659266114, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (81, 1, 1659275350, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (82, 1, 1659275639, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (83, 1, 1659275696, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (84, 1, 1659275714, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (85, 1, 1659275861, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (86, 1, 1659277167, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (87, 1, 1659277870, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (88, 1, 1659277988, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (89, 1, 1659278373, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (90, 1, 1659279074, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (91, 1, 1659279739, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (92, 1, 1659279831, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (93, 1, 1659280429, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (94, 1, 1659282443, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (95, 1, 1659282577, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (96, 1, 1659286487, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (97, 1, 1659286506, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (98, 1, 1659286601, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (99, 1, 1659286610, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (100, 1, 1659286981, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (101, 1, 1659287095, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (102, 1, 1659287618, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (103, 1, 1659287712, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (104, 1, 1659287722, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (105, 1, 1659287730, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (106, 1, 1659287769, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (107, 1, 1659287995, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (108, 1, 1659288388, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (109, 1, 1659288410, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (110, 1, 1659288420, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (111, 1, 1659288451, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (112, 1, 1659288459, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (113, 1, 1659288614, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (114, 1, 1659288818, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (115, 1, 1659288864, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (116, 1, 1659288960, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (117, 1, 1659288968, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (118, 1, 1659289022, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (119, 1, 1659289059, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (120, 1, 1659289083, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (121, 1, 1659289476, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (122, 1, 1659289484, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (123, 1, 1659289510, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (124, 1, 1659289607, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (125, 1, 1659289611, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (126, 1, 1659290215, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (127, 1, 1659290366, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (128, 1, 1659290373, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (129, 1, 1659290470, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (130, 1, 1659290471, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (131, 1, 1659290471, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (132, 1, 1659290471, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (133, 1, 1659290472, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (134, 1, 1659290478, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (135, 1, 1659290562, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (136, 1, 1659290569, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (137, 1, 1659318194, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (138, 1, 1659318471, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (139, 1, 1659318479, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (140, 1, 1659318741, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (141, 1, 1659318745, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (142, 1, 1659319121, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (143, 1, 1659319135, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (144, 1, 1659320569, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (145, 11, 1659320676, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (146, 1, 1659320705, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (147, 11, 1659324711, 'boss', 1, '0.0.0.0', '', 1);
INSERT INTO `log_login` VALUES (148, 1, 1659339176, 'boss', 1, '0.0.0.0', '', 1);

-- ----------------------------
-- Table structure for model
-- ----------------------------
DROP TABLE IF EXISTS `model`;
CREATE TABLE `model`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '?????????',
  `ename` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
  `cname` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
  `notes` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '????????????',
  `menugroup` tinyint NOT NULL DEFAULT 2 COMMENT '???????????????',
  `sort` tinyint NOT NULL DEFAULT 1 COMMENT '??????',
  `author` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
  `version` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '??????',
  `create_time` int NOT NULL DEFAULT 0,
  `update_time` int NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '??????',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ename`(`ename`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '?????????' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of model
-- ----------------------------
INSERT INTO `model` VALUES (1, 'article', '????????????', '', 2, 1, '', '', 0, 0, 1);
INSERT INTO `model` VALUES (4, 'photo', '????????????', '', 2, 1, '', '', 0, 1656473819, 1);
INSERT INTO `model` VALUES (27, 'news', '??????????????????', '', 2, 1, '', '', 0, 1656473829, 0);
INSERT INTO `model` VALUES (28, 'case', '????????????', '', 2, 1, '', '', 0, 1656473837, 0);
INSERT INTO `model` VALUES (29, 'product', '??????????????????', NULL, 2, 1, '', '', 0, 1656473846, 1);
INSERT INTO `model` VALUES (30, 'about', '??????????????????', NULL, 2, 1, '', '', 0, 1656473854, 1);

-- ----------------------------
-- Table structure for node
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `sort` smallint UNSIGNED NOT NULL DEFAULT 0,
  `pid` smallint UNSIGNED NOT NULL DEFAULT 0,
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `type` tinyint(1) NOT NULL DEFAULT 0,
  `group_id` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `icon` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '??????',
  `left_menu_action` tinyint(1) UNSIGNED ZEROFILL NOT NULL COMMENT '??????????????????????????????0?????????1???????????????0',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `level`(`level`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 314 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of node
-- ----------------------------
INSERT INTO `node` VALUES (1, 'boss', '????????????', 1, '', 0, 0, 1, 0, '0', '', 0);
INSERT INTO `node` VALUES (2, 'node', '????????????', 1, '', 100, 1, 2, 0, 'developers', '', 0);
INSERT INTO `node` VALUES (3, 'user', '???????????????', 1, '', 98, 1, 2, 0, 'groupuser', '', 0);
INSERT INTO `node` VALUES (4, 'role', '????????????', 1, '', 99, 1, 2, 0, 'groupuser', '', 0);
INSERT INTO `node` VALUES (5, 'public', '????????????', 1, '', 0, 1, 2, 0, '0', '', 0);
INSERT INTO `node` VALUES (6, 'index', '????????????', 1, '', 0, 1, 2, 0, '0', '', 0);
INSERT INTO `node` VALUES (7, 'index', '????????????', 1, '', 0, 6, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (155, 'forbid', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (154, 'delete', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (153, 'foreverdelete', '????????????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (152, 'doedit', '????????????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (16, 'category', '????????????', 1, '', 99, 1, 2, 0, 'framework', '', 0);
INSERT INTO `node` VALUES (17, 'article', '????????????', 1, '', 100, 1, 2, 0, 'content', '', 0);
INSERT INTO `node` VALUES (151, 'edit', '????????????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (150, 'doadd', '????????????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (28, 'photo', '????????????', 1, '', 98, 1, 2, 0, 'content', 'fa fa-file-picture-o', 0);
INSERT INTO `node` VALUES (29, 'Link', '????????????', 0, '', 95, 1, 2, 0, 'content', '', 0);
INSERT INTO `node` VALUES (32, 'system', '????????????', 1, '', 96, 1, 2, 0, 'system', '', 0);
INSERT INTO `node` VALUES (33, 'router', '????????????', 0, '', 90, 1, 2, 0, 'content', '', 0);
INSERT INTO `node` VALUES (34, 'static', '??????????????????', 1, '', 97, 1, 2, 0, 'system', 'fa fa-file-code-o', 0);
INSERT INTO `node` VALUES (80, 'tree', '????????????', 1, '', 0, 5, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (81, 'seltpl', '????????????', 1, '', 0, 5, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (82, 'ui', '????????????', 1, '', 10, 1, 2, 0, '0', '', 0);
INSERT INTO `node` VALUES (112, 'attribute', '????????????', 1, '', 98, 1, 2, 0, 'framework', '', 0);
INSERT INTO `node` VALUES (149, 'add', '????????????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (148, 'index', '????????????', 1, '', 999, 2, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (142, 'index', '????????????', 1, '', 0, 17, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (143, 'add', '????????????', 1, '', 0, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (144, 'doadd', '????????????', 1, '', 0, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (145, 'edit', '????????????', 1, '', 0, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (146, 'doedit', '????????????', 1, '', 0, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (147, 'foreverdelete', '????????????', 1, '', 0, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (156, 'resume', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (157, 'checkPass', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (158, 'recycle', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (159, 'setSort', '??????', 1, '', 999, 2, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (160, 'index', '???????????????', 1, '', 999, 3, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (161, 'add', '????????????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (162, 'doadd', '????????????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (163, 'edit', '????????????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (164, 'doedit', '????????????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (165, 'foreverdelete', '????????????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (166, 'delete', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (167, 'forbid', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (168, 'resume', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (169, 'checkPass', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (170, 'recycle', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (171, 'setSort', '??????', 1, '', 999, 3, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (172, 'index', '????????????', 1, '', 999, 4, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (173, 'add', '????????????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (174, 'doadd', '????????????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (175, 'edit', '????????????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (176, 'doedit', '????????????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (177, 'foreverdelete', '????????????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (178, 'delete', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (179, 'forbid', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (180, 'resume', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (181, 'checkPass', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (182, 'recycle', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (183, 'setSort', '??????', 1, '', 999, 4, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (184, 'index', '????????????', 1, '', 999, 16, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (185, 'add', '????????????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (186, 'doadd', '????????????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (187, 'edit', '????????????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (188, 'doedit', '????????????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (189, 'foreverdelete', '????????????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (190, 'delete', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (191, 'forbid', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (192, 'resume', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (193, 'checkPass', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (194, 'recycle', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (195, 'setSort', '??????', 1, '', 999, 16, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (196, 'delete', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (197, 'forbid', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (198, 'resume', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (199, 'checkPass', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (200, 'recycle', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (201, 'setSort', '??????', 1, '', 999, 17, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (202, 'index', '????????????', 1, '', 999, 28, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (203, 'add', '????????????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (204, 'doadd', '????????????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (205, 'edit', '????????????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (206, 'doedit', '????????????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (207, 'foreverdelete', '????????????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (208, 'delete', '??????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (209, 'forbid', '??????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (210, 'resume', '??????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (211, 'checkPass', '??????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (212, 'recycle', '??????', 1, '', 999, 28, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (213, 'setSort', '??????', 1, '', 999, 28, 3, 0, '0', 'fa fa-etsy', 0);
INSERT INTO `node` VALUES (214, 'index', '??????', 1, '', 999, 29, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (215, 'add', '????????????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (216, 'doadd', '????????????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (217, 'edit', '????????????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (218, 'doedit', '????????????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (219, 'foreverdelete', '????????????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (220, 'delete', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (221, 'forbid', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (222, 'resume', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (223, 'checkPass', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (224, 'recycle', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (225, 'setSort', '??????', 1, '', 999, 29, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (226, 'index', '??????', 1, '', 999, 33, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (229, 'edit', '????????????', 1, '', 999, 33, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (230, 'doedit', '????????????', 1, '', 999, 33, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (237, 'setSort', '??????', 1, '', 999, 33, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (238, 'index', '????????????', 1, '', 999, 112, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (239, 'add', '????????????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (240, 'doadd', '????????????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (241, 'edit', '????????????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (242, 'doedit', '????????????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (243, 'foreverdelete', '????????????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (244, 'delete', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (245, 'forbid', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (246, 'resume', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (247, 'checkPass', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (248, 'recycle', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (249, 'setSort', '??????', 1, '', 999, 112, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (250, 'news', '????????????', 0, '', 999, 1, 2, 0, 'content', '', 0);
INSERT INTO `node` VALUES (251, 'index', '??????', 1, '', 999, 250, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (252, 'add', '????????????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (253, 'doadd', '????????????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (254, 'edit', '????????????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (255, 'doedit', '????????????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (256, 'foreverdelete', '????????????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (257, 'delete', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (258, 'forbid', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (259, 'resume', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (260, 'checkPass', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (261, 'recycle', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (262, 'setSort', '??????', 1, '', 999, 250, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (263, 'case', '??????', 0, '', 999, 1, 2, 0, 'content', '', 0);
INSERT INTO `node` VALUES (264, 'index', '????????????', 1, '', 999, 263, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (265, 'add', '????????????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (266, 'doadd', '????????????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (267, 'edit', '????????????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (268, 'doedit', '????????????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (269, 'foreverdelete', '????????????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (270, 'delete', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (271, 'forbid', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (272, 'resume', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (273, 'checkPass', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (274, 'recycle', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (275, 'setSort', '??????', 1, '', 999, 263, 3, 0, '0', '', 0);
INSERT INTO `node` VALUES (276, 'product', '????????????', 0, '', 999, 1, 2, 0, 'content', 'table', 0);
INSERT INTO `node` VALUES (277, 'index', '????????????', 1, '', 999, 276, 3, 0, '0', 'table', 1);
INSERT INTO `node` VALUES (278, 'add', '??????', 1, '', 999, 276, 3, 0, '0', 'table', 0);
INSERT INTO `node` VALUES (279, 'doadd', '????????????', 1, '', 999, 276, 3, 0, '0', 'table', 0);
INSERT INTO `node` VALUES (280, 'edit', '??????', 1, '', 999, 276, 3, 0, '0', 'table', 0);
INSERT INTO `node` VALUES (281, 'doedit', '????????????', 1, '', 999, 276, 3, 0, '0', 'table', 0);
INSERT INTO `node` VALUES (282, 'delete', '??????', 1, '', 999, 276, 3, 0, '0', 'table', 0);
INSERT INTO `node` VALUES (283, 'about', '????????????', 0, '', 999, 1, 2, 0, 'content', 'table', 0);
INSERT INTO `node` VALUES (284, 'index', '????????????', 1, '', 999, 283, 3, 0, '0', '', 1);
INSERT INTO `node` VALUES (285, 'add', '????????????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (286, 'doadd', '????????????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (287, 'edit', '????????????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (288, 'doedit', '????????????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (289, 'foreverdelete', '????????????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (290, 'delete', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (291, 'forbid', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (292, 'resume', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (293, 'checkPass', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (294, 'recycle', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (295, 'setSort', '??????', 1, '', 999, 283, 3, 0, '0', NULL, 0);
INSERT INTO `node` VALUES (300, 'test2', '???????????????', 0, '??????', 999, 1, 2, 0, 'content', 'table', 0);
INSERT INTO `node` VALUES (301, 'index', '??????', 1, '', 999, 300, 3, 0, '0', 'table', 1);
INSERT INTO `node` VALUES (302, 'dadas', '????????????', 1, '?????????', 999, 0, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (303, 'model', '????????????', 1, '', 999, 1, 2, 0, 'framework', 'fa fa-life-buoy', 0);
INSERT INTO `node` VALUES (304, 'index', '????????????', 1, '', 999, 303, 3, 0, '', '', 1);
INSERT INTO `node` VALUES (306, 'add', '??????', 1, '', 999, 303, 3, 0, '', 'fa fa-plus', 0);
INSERT INTO `node` VALUES (307, 'edit', '??????', 1, '', 999, 303, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (308, 'del', '??????', 1, '', 999, 303, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (309, 'resume', '??????', 1, '', 999, 303, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (310, 'forbid', '??????', 1, '', 999, 303, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (311, 'recycle', '??????', 1, '', 999, 303, 3, 0, '', '', 0);
INSERT INTO `node` VALUES (312, 'index', '????????????', 1, '????????????', 999, 32, 3, 0, '', 'fa fa-sun-o', 1);
INSERT INTO `node` VALUES (313, 'index', '??????????????????', 1, '', 999, 34, 3, 0, '', 'fa fa-ellipsis-v', 1);

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '????????????',
  `tid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '??????ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '??????',
  `img` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '????????????',
  `create_time` int UNSIGNED NOT NULL COMMENT '????????????',
  `update_time` int UNSIGNED NOT NULL COMMENT '????????????',
  `creater_id` int UNSIGNED NOT NULL COMMENT '?????????',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '??????',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of photo
-- ----------------------------
INSERT INTO `photo` VALUES (1, '38', 'banner1', '/static/uploads/2022-6-24/756b2deb-cc8e-8501-b83b-510d8428faf1.png', 1656052619, 1656052619, 1, 1);
INSERT INTO `photo` VALUES (2, '38', 'banner2', '/static/uploads/2022-6-24/a71bdf04-6add-8359-e842-3ff415a5f4cd.png', 1656052670, 1656052670, 1, 1);
INSERT INTO `photo` VALUES (3, '38', 'banner3', '/static/uploads/2022-6-24/a211569f-ffa0-3a82-2f54-01aa306bc004.png', 1656052687, 1656052687, 1, 1);
INSERT INTO `photo` VALUES (4, '38', 'banner4', '/static/uploads/2022-6-24/dae8ffec-97f3-10e0-4a6e-814acbe8bbac.png', 1656052701, 1656052701, 1, 1);
INSERT INTO `photo` VALUES (5, '38', 'banner5', '/static/uploads/2022-6-24/48b40f83-7330-7b7c-c66b-63af0fbfb841.png', 1656052723, 1656052723, 1, -1);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `pid` smallint NOT NULL DEFAULT 0,
  `status` tinyint NOT NULL DEFAULT 0,
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `ename` varchar(5) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `create_time` int UNSIGNED NOT NULL,
  `update_time` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `parentId`(`pid`) USING BTREE,
  INDEX `ename`(`ename`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (2, '?????????', 0, 1, '??????', '', 0, 1426173947);
INSERT INTO `role` VALUES (3, '?????????', 0, 1, '???????????????', '', 0, 1426386658);
INSERT INTO `role` VALUES (4, '??????1', 2, 1, '??????', '', 0, 1654155540);
INSERT INTO `role` VALUES (5, '123', 2, 1, '', '', 0, 1654593929);

-- ----------------------------
-- Table structure for role_user
-- ----------------------------
DROP TABLE IF EXISTS `role_user`;
CREATE TABLE `role_user`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0,
  `user_id` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0',
  INDEX `group_id`(`role_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Records of role_user
-- ----------------------------
INSERT INTO `role_user` VALUES (3, '2');
INSERT INTO `role_user` VALUES (5, '5');
INSERT INTO `role_user` VALUES (5, '2');
INSERT INTO `role_user` VALUES (3, '11');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `object_id` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '??????????????????',
  `account` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '?????????',
  `nickname` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `logo` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '??????',
  `age` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '??????',
  `sex` tinyint UNSIGNED NOT NULL COMMENT '?????????1?????????2??????',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `phone` varchar(13) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `email` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `create_time` int UNSIGNED NOT NULL,
  `update_time` int UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '???????????????1???1?????????0?????????????????????-1?????????',
  `type_id` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '9????????????????????????',
  PRIMARY KEY (`id`, `object_id`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `objectId`(`object_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '55224b780064f91200000001', 'admin', '???????????????', NULL, 0, 2, '21232f297a57a5a743894a0e4a801fc3', '13662110436', '123270073@qq.com', 1381212542, 1656127756, 1, 9);
INSERT INTO `user` VALUES (2, '55224b780064f91200000108', 'langzishen', '?????????', NULL, 0, 1, '21232f297a57a5a743894a0e4a801fc3', NULL, '1606737725@qq.com', 1426385437, 1426401231, 1, 9);
INSERT INTO `user` VALUES (11, 'e5jpusqe', 'test2', 'test2', '', 0, 0, '21232f297a57a5a743894a0e4a801fc3', '13662110436', '123270073@qq.com', 1656396165, 1656396165, 1, 1);

SET FOREIGN_KEY_CHECKS = 1;

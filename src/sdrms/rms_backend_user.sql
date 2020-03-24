/*
Navicat MySQL Data Transfer

Source Server         : 123.56.4.34
Source Server Version : 50728
Source Host           : 123.56.4.34:3306
Source Database       : azkaban

Target Server Type    : MYSQL
Target Server Version : 50728
File Encoding         : 65001

Date: 2020-03-16 17:07:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rms_backend_user
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user`;
CREATE TABLE `rms_backend_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_backend_user
-- ----------------------------
INSERT INTO `rms_backend_user` VALUES ('1', '超级管理员', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '18612348765', 'lhtzbj18@126.com', '/static/upload/1.jpg');
INSERT INTO `rms_backend_user` VALUES ('6', '张良', 'zhangliang', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '');
INSERT INTO `rms_backend_user` VALUES ('7', '刘悲', 'liubei', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '');

-- ----------------------------
-- Table structure for rms_backend_user_rms_roles
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user_rms_roles`;
CREATE TABLE `rms_backend_user_rms_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rms_backend_user_id` int(11) NOT NULL,
  `rms_role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_backend_user_rms_roles
-- ----------------------------

-- ----------------------------
-- Table structure for rms_course
-- ----------------------------
DROP TABLE IF EXISTS `rms_course`;
CREATE TABLE `rms_course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `short_name` varchar(8) DEFAULT '',
  `price` double DEFAULT '0',
  `real_price` double DEFAULT '0',
  `img` varchar(256) DEFAULT '',
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `seq` int(11) DEFAULT '0',
  `creator_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_course
-- ----------------------------
INSERT INTO `rms_course` VALUES ('36', '高一数学上册', '', '0', '0', '', null, null, '0', '1');
INSERT INTO `rms_course` VALUES ('37', '高一数学下册', '', '0', '0', '', null, null, '0', '1');
INSERT INTO `rms_course` VALUES ('38', '高二语文上册', '', '0', '0', '', null, null, '0', '1');
INSERT INTO `rms_course` VALUES ('39', '高二语文下册', '', '0', '0', '', null, null, '0', '1');

-- ----------------------------
-- Table structure for rms_course_limited_info
-- ----------------------------
DROP TABLE IF EXISTS `rms_course_limited_info`;
CREATE TABLE `rms_course_limited_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL COMMENT '课程id',
  `user_id` int(11) NOT NULL COMMENT '教师id',
  `number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '课程限时打卡编号',
  `create_time` datetime DEFAULT NULL,
  `start_time` datetime DEFAULT NULL COMMENT '考勤开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '考勤结束时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of rms_course_limited_info
-- ----------------------------
INSERT INTO `rms_course_limited_info` VALUES ('18', '37', '6', '103254', '2020-03-16 15:13:38', '2020-03-16 15:13:38', '2020-03-16 15:14:38');

-- ----------------------------
-- Table structure for rms_resource
-- ----------------------------
DROP TABLE IF EXISTS `rms_resource`;
CREATE TABLE `rms_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `parent_id` int(11) DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=204 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_resource
-- ----------------------------
INSERT INTO `rms_resource` VALUES ('7', '1', '权限管理', '8', '100', 'fa fa-balance-scale', '');
INSERT INTO `rms_resource` VALUES ('8', '0', '系统菜单', null, '200', '', '');
INSERT INTO `rms_resource` VALUES ('9', '1', '资源管理', '7', '100', '', 'ResourceController.Index');
INSERT INTO `rms_resource` VALUES ('12', '1', '角色管理', '7', '100', '', 'RoleController.Index');
INSERT INTO `rms_resource` VALUES ('13', '1', '教师管理', '21', '100', '', 'BackendUserController.Index');
INSERT INTO `rms_resource` VALUES ('14', '1', '系统管理', '8', '90', 'fa fa-gears', '');
INSERT INTO `rms_resource` VALUES ('21', '0', '业务菜单', null, '170', '', '');
INSERT INTO `rms_resource` VALUES ('22', '1', '课程管理', '21', '100', 'fa fa-book', 'CourseController.Index');
INSERT INTO `rms_resource` VALUES ('23', '1', '日志管理(空)', '14', '100', '', '');
INSERT INTO `rms_resource` VALUES ('25', '2', '编辑', '9', '100', 'fa fa-pencil', 'ResourceController.Edit');
INSERT INTO `rms_resource` VALUES ('26', '2', '编辑', '13', '100', 'fa fa-pencil', 'BackendUserController.Edit');
INSERT INTO `rms_resource` VALUES ('27', '2', '删除', '9', '100', 'fa fa-trash', 'ResourceController.Delete');
INSERT INTO `rms_resource` VALUES ('29', '2', '删除', '13', '100', 'fa fa-trash', 'BackendUserController.Delete');
INSERT INTO `rms_resource` VALUES ('30', '2', '编辑', '12', '100', 'fa fa-pencil', 'RoleController.Edit');
INSERT INTO `rms_resource` VALUES ('31', '2', '删除', '12', '100', 'fa fa-trash', 'RoleController.Delete');
INSERT INTO `rms_resource` VALUES ('32', '2', '分配资源', '12', '100', 'fa fa-th', 'RoleController.Allocate');
INSERT INTO `rms_resource` VALUES ('35', '1', ' 首页', null, '100', 'fa fa-dashboard', 'HomeController.Index');
INSERT INTO `rms_resource` VALUES ('36', '2', '编辑', '22', '100', '', 'CourseController.Edit');
INSERT INTO `rms_resource` VALUES ('37', '2', '删除', '22', '100', '', 'CourseController.Delete');
INSERT INTO `rms_resource` VALUES ('201', '1', '学生管理', '21', '100', '', 'StudentInfoController.Index');
INSERT INTO `rms_resource` VALUES ('202', '1', '考勤管理', '21', '100', '', 'StuKqInfoController.Index');
INSERT INTO `rms_resource` VALUES ('203', '1', '限时考勤', '21', '100', '', 'UserCourseController.Index');

-- ----------------------------
-- Table structure for rms_role
-- ----------------------------
DROP TABLE IF EXISTS `rms_role`;
CREATE TABLE `rms_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role
-- ----------------------------
INSERT INTO `rms_role` VALUES ('22', '超级管理员', '20');
INSERT INTO `rms_role` VALUES ('24', '角色管理员', '10');
INSERT INTO `rms_role` VALUES ('25', '课程资源管理员', '5');

-- ----------------------------
-- Table structure for rms_role_backenduser_rel
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_backenduser_rel`;
CREATE TABLE `rms_role_backenduser_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role_backenduser_rel
-- ----------------------------
INSERT INTO `rms_role_backenduser_rel` VALUES ('61', '22', '1', '2017-12-18 07:35:58');
INSERT INTO `rms_role_backenduser_rel` VALUES ('72', '25', '6', '2020-03-16 14:56:32');
INSERT INTO `rms_role_backenduser_rel` VALUES ('73', '25', '7', '2020-03-16 14:57:21');

-- ----------------------------
-- Table structure for rms_role_resource_rel
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_resource_rel`;
CREATE TABLE `rms_role_resource_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=540 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role_resource_rel
-- ----------------------------
INSERT INTO `rms_role_resource_rel` VALUES ('448', '24', '8', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('449', '24', '14', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('450', '24', '23', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('477', '25', '21', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('478', '25', '36', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('479', '25', '37', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('480', '25', '201', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('481', '25', '202', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('482', '25', '203', '2020-03-15 13:17:20');
INSERT INTO `rms_role_resource_rel` VALUES ('522', '22', '21', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('523', '22', '13', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('524', '22', '26', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('525', '22', '29', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('526', '22', '22', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('527', '22', '36', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('528', '22', '37', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('529', '22', '202', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('530', '22', '14', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('531', '22', '23', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('532', '22', '7', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('533', '22', '9', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('534', '22', '25', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('535', '22', '27', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('536', '22', '12', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('537', '22', '30', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('538', '22', '31', '2020-03-15 14:08:49');
INSERT INTO `rms_role_resource_rel` VALUES ('539', '22', '32', '2020-03-15 14:08:49');

-- ----------------------------
-- Table structure for rms_stu_kq_info
-- ----------------------------
DROP TABLE IF EXISTS `rms_stu_kq_info`;
CREATE TABLE `rms_stu_kq_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stu_id` int(11) NOT NULL COMMENT '学生id',
  `course_id` int(11) NOT NULL COMMENT '课程id',
  `user_id` int(11) NOT NULL COMMENT '教师id',
  `course_limited_id` int(11) NOT NULL COMMENT '课程限时考勤id',
  `create_time` datetime DEFAULT NULL COMMENT '学生打卡时间',
  `status` int(2) DEFAULT NULL COMMENT '考勤状态：0打卡成功，1缺勤',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of rms_stu_kq_info
-- ----------------------------
INSERT INTO `rms_stu_kq_info` VALUES ('1', '4', '3', '1', '10', '2020-03-14 17:22:19', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('2', '5', '3', '1', '10', '2020-03-14 17:22:19', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('3', '4', '1', '1', '9', '2020-03-14 17:45:59', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('4', '5', '3', '1', '11', '2020-03-14 17:45:59', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('5', '6', '3', '1', '11', '2020-03-14 17:45:59', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('6', '6', '1', '1', '9', '2020-03-14 17:45:59', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('7', '4', '3', '1', '11', '2020-03-14 18:10:41', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('8', '5', '5', '6', '12', '2020-03-15 13:21:25', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('9', '6', '5', '6', '12', '2020-03-15 13:21:25', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('10', '5', '5', '6', '13', '2020-03-15 14:39:38', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('11', '6', '5', '6', '13', '2020-03-15 14:39:38', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('13', '6', '5', '6', '14', '2020-03-16 11:26:23', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('16', '5', '4', '6', '15', '2020-03-16 12:56:42', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('17', '6', '4', '6', '15', '2020-03-16 12:56:42', '1');
INSERT INTO `rms_stu_kq_info` VALUES ('18', '5', '37', '6', '17', '2020-03-16 15:01:47', '0');
INSERT INTO `rms_stu_kq_info` VALUES ('19', '6', '37', '6', '18', '2020-03-16 15:13:53', '0');
INSERT INTO `rms_stu_kq_info` VALUES ('20', '5', '37', '6', '18', '2020-03-16 15:14:28', '0');
INSERT INTO `rms_stu_kq_info` VALUES ('21', '7', '37', '6', '18', '2020-03-16 15:15:49', '1');

-- ----------------------------
-- Table structure for rms_student_info
-- ----------------------------
DROP TABLE IF EXISTS `rms_student_info`;
CREATE TABLE `rms_student_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sno` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '学号',
  `real_name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of rms_student_info
-- ----------------------------
INSERT INTO `rms_student_info` VALUES ('5', '2020031611', '赵云', 'e10adc3949ba59abbe56e057f20f883e', '2020-03-16 15:10:22');
INSERT INTO `rms_student_info` VALUES ('6', '2020031612', '张山峰', 'e10adc3949ba59abbe56e057f20f883e', '2020-03-16 15:10:22');
INSERT INTO `rms_student_info` VALUES ('7', '2020031613', '黄忠', 'e10adc3949ba59abbe56e057f20f883e', '2020-03-16 15:10:22');

-- ----------------------------
-- Table structure for rms_user_course_ref
-- ----------------------------
DROP TABLE IF EXISTS `rms_user_course_ref`;
CREATE TABLE `rms_user_course_ref` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '教师id',
  `course_id` int(11) NOT NULL COMMENT '课程id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of rms_user_course_ref
-- ----------------------------
INSERT INTO `rms_user_course_ref` VALUES ('1', '1', '1');
INSERT INTO `rms_user_course_ref` VALUES ('2', '1', '3');
INSERT INTO `rms_user_course_ref` VALUES ('3', '1', '6');
INSERT INTO `rms_user_course_ref` VALUES ('4', '1', '7');
INSERT INTO `rms_user_course_ref` VALUES ('5', '1', '8');
INSERT INTO `rms_user_course_ref` VALUES ('6', '1', '9');
INSERT INTO `rms_user_course_ref` VALUES ('7', '1', '10');
INSERT INTO `rms_user_course_ref` VALUES ('8', '1', '11');
INSERT INTO `rms_user_course_ref` VALUES ('9', '1', '12');
INSERT INTO `rms_user_course_ref` VALUES ('10', '1', '13');
INSERT INTO `rms_user_course_ref` VALUES ('11', '1', '14');
INSERT INTO `rms_user_course_ref` VALUES ('12', '1', '15');
INSERT INTO `rms_user_course_ref` VALUES ('13', '1', '16');
INSERT INTO `rms_user_course_ref` VALUES ('14', '1', '17');
INSERT INTO `rms_user_course_ref` VALUES ('15', '1', '18');
INSERT INTO `rms_user_course_ref` VALUES ('16', '1', '19');
INSERT INTO `rms_user_course_ref` VALUES ('17', '1', '20');
INSERT INTO `rms_user_course_ref` VALUES ('18', '1', '21');
INSERT INTO `rms_user_course_ref` VALUES ('19', '1', '22');
INSERT INTO `rms_user_course_ref` VALUES ('20', '1', '23');
INSERT INTO `rms_user_course_ref` VALUES ('21', '1', '24');
INSERT INTO `rms_user_course_ref` VALUES ('37', '6', '36');
INSERT INTO `rms_user_course_ref` VALUES ('38', '6', '37');
INSERT INTO `rms_user_course_ref` VALUES ('39', '7', '38');

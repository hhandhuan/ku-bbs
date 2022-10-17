/*
 Navicat Premium Data Transfer

 Source Server         : localhost-db
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : ku-bbs

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 18/09/2022 21:27:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for checkins
-- ----------------------------
DROP TABLE IF EXISTS `checkins`;
CREATE TABLE `checkins` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户 ID',
  `cumulative_days` bigint unsigned DEFAULT '0' COMMENT '累计签到(天)',
  `continuity_days` bigint unsigned DEFAULT '0' COMMENT '连续签到(天)',
  `last_time` timestamp NULL DEFAULT NULL COMMENT '最后签到时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '评论人 ID',
  `reply_id` bigint unsigned DEFAULT '0' COMMENT '回复人 ID',
  `topic_id` bigint unsigned DEFAULT NULL COMMENT '话题 ID',
  `target_id` bigint unsigned DEFAULT '0' COMMENT '回复目标 ID',
  `content` text COMMENT '回复内容',
  `md_content` text COMMENT 'MD 内容',
  `like_count` bigint unsigned DEFAULT '0' COMMENT '喜欢统计',
  `floor` bigint unsigned DEFAULT '0' COMMENT '评论楼层',
  `reply_floor` bigint DEFAULT NULL COMMENT '被评论楼层',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `topic_id_index` (`topic_id`) USING BTREE,
  KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户 ID',
  `target_id` bigint unsigned DEFAULT NULL COMMENT '被关注用户 ID',
  `state` tinyint DEFAULT NULL COMMENT '状态:1-关注/0-取消',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `target_id_index` (`target_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for integral_logs
-- ----------------------------
DROP TABLE IF EXISTS `integral_logs`;
CREATE TABLE `integral_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '用户 ID',
  `rewards` bigint DEFAULT NULL COMMENT '奖励积分',
  `mode` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '获取方式',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '用户 ID',
  `target_user_id` bigint unsigned DEFAULT '0' COMMENT '目标用户ID',
  `source_id` bigint unsigned DEFAULT '0' COMMENT '资源 ID',
  `source_type` char(50) DEFAULT NULL COMMENT '资源类型:topic/comment',
  `state` tinyint unsigned DEFAULT '0' COMMENT '状态: 0-否/1-是',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_source_id_index` (`user_id`,`source_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for nodes
-- ----------------------------
DROP TABLE IF EXISTS `nodes`;
CREATE TABLE `nodes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `name` char(50) DEFAULT NULL COMMENT '节点名称',
  `alias` char(50) DEFAULT NULL COMMENT '节点别名',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '节点介绍',
  `count` bigint unsigned DEFAULT '0' COMMENT '资源统计',
  `pid` bigint unsigned DEFAULT '0' COMMENT '节点父级',
  `sort` tinyint unsigned DEFAULT '0' COMMENT '排序值',
  `state` tinyint unsigned DEFAULT '1' COMMENT '节点状态:0-关闭/1-开启',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `alias_unique_index` (`alias`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for reminds
-- ----------------------------
DROP TABLE IF EXISTS `reminds`;
CREATE TABLE `reminds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sender` bigint unsigned DEFAULT NULL COMMENT '发送人ID',
  `receiver` bigint unsigned DEFAULT NULL COMMENT '接受者ID',
  `source_id` bigint unsigned DEFAULT NULL COMMENT '资源 ID',
  `source_type` char(30) DEFAULT NULL COMMENT '资源类型',
  `source_content` varchar(255) DEFAULT NULL COMMENT '资源内容',
  `source_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '提醒发生地址',
  `action` char(30) DEFAULT NULL COMMENT '动作类型',
  `readed_at` timestamp NULL DEFAULT NULL COMMENT '阅读时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `receiver_index` (`receiver`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for reports
-- ----------------------------
DROP TABLE IF EXISTS `reports`;
CREATE TABLE `reports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '举报人',
  `remark` varchar(255) DEFAULT NULL COMMENT '举报备注',
  `target_id` bigint unsigned DEFAULT '0' COMMENT '被举报人ID',
  `source_id` bigint unsigned DEFAULT '0' COMMENT '目标ID',
  `source_type` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '目标类型',
  `source_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '目标链接',
  `handler_id` bigint unsigned DEFAULT '0' COMMENT '处理人ID',
  `state` tinyint unsigned DEFAULT '0' COMMENT '状态:0-待处理/1-已处理',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for system_notices
-- ----------------------------
DROP TABLE IF EXISTS `system_notices`;
CREATE TABLE `system_notices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) DEFAULT NULL COMMENT '消息标题',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '发布人 ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '消息内容',
  `md_content` text COMMENT 'markdown 内容',
  `target_id` varchar(255) DEFAULT '0' COMMENT '接受者 ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for system_user_notices
-- ----------------------------
DROP TABLE IF EXISTS `system_user_notices`;
CREATE TABLE `system_user_notices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '用户 ID',
  `notice_id` bigint unsigned DEFAULT '0' COMMENT '通知 ID',
  `readed_at` timestamp NULL DEFAULT NULL COMMENT '阅读时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `node_id` bigint DEFAULT NULL COMMENT '分类 ID',
  `user_id` bigint DEFAULT NULL COMMENT '用户 ID',
  `reply_id` bigint unsigned DEFAULT '0' COMMENT '最后回复者ID',
  `title` varchar(255) DEFAULT NULL COMMENT '话题标题',
  `tags` varchar(255) DEFAULT NULL COMMENT '话题标签',
  `comment_count` bigint unsigned DEFAULT '0' COMMENT '评论统计',
  `view_count` bigint unsigned DEFAULT '0' COMMENT '浏览统计',
  `like_count` bigint unsigned DEFAULT '0' COMMENT '喜欢统计',
  `state` tinyint unsigned DEFAULT '0' COMMENT '话题状态: 0-暂存/1-发布',
  `comment_state` tinyint unsigned DEFAULT '1' COMMENT '讨论状态:0-关闭/1-开启',
  `type` tinyint unsigned DEFAULT '0' COMMENT '话题类型:0-默认/1-精华/2-置顶',
  `content` text COMMENT '话题内容',
  `md_content` text COMMENT 'markdown内容',
  `last_reply_at` timestamp NULL DEFAULT NULL COMMENT '回复时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` char(50) DEFAULT NULL COMMENT '用户名',
  `gender` tinyint unsigned DEFAULT '0' COMMENT '性别:1-男/2-女/0-未知',
  `city` char(50) DEFAULT NULL COMMENT '城市',
  `email` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `site` varchar(255) DEFAULT NULL COMMENT '个人网站',
  `job` char(50) DEFAULT NULL COMMENT '职业',
  `desc` varchar(255) DEFAULT NULL COMMENT '简介',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `is_admin` tinyint unsigned DEFAULT '0' COMMENT '是否管理员:1-是/0-否',
  `integral` bigint unsigned DEFAULT '0' COMMENT '个人积分',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态:1-正常/0-禁用',
  `last_login_ip` char(20) DEFAULT NULL COMMENT '最后登录IP',
  `last_login_at` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

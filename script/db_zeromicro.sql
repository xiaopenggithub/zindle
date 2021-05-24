-- MySQL dump 10.13  Distrib 5.7.32, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: db_zeromicro
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `book_orders`
--

DROP TABLE IF EXISTS `book_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `book_orders` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int DEFAULT '0' COMMENT '用户ID',
  `book_id` int DEFAULT '0' COMMENT '书ID',
  `return_date` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '还书日期',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间(借书日期)',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='图书订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book_orders`
--

LOCK TABLES `book_orders` WRITE;
/*!40000 ALTER TABLE `book_orders` DISABLE KEYS */;
INSERT INTO `book_orders` VALUES (1,1,1,'2021-05-23 21:35:19','2021-05-23 21:35:08','2021-05-23 21:35:19','2006-01-02 15:04:05'),(2,1,5,'2006-01-02 15:04:05','2021-05-23 21:47:37','2021-05-23 21:47:37','2006-01-02 15:04:05'),(3,1,6,'2006-01-02 15:04:05','2021-05-23 21:49:31','2021-05-23 21:49:31','2006-01-02 15:04:05');
/*!40000 ALTER TABLE `book_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `description` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '简介',
  `quantity` int DEFAULT '0' COMMENT '数量',
  `cover` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '封面',
  `sort` int DEFAULT '0' COMMENT '排序',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '创建者',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '更新者',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='图书信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'西游记','《西游记》是“一部具有丰富内容和光辉思想的神话小说”，它是明代小说家吴承恩以积极浪漫主义手',11,'uploads/cbfacb9df0c7caf9a2b8a8ffbd72d1a0_20210523214513.jpeg',11,'','','2021-05-23 21:30:58','2021-05-23 21:45:13','2006-01-02 15:04:05'),(2,'红楼梦','《红楼梦》塑造了众多的人物形象，他们各自具有自己独特而鲜明的个性特征，成为不朽的艺术典型，在中国文学史和世界文学史上永远放射着奇光异彩。',1,'uploads/95d936f124af05ac385691acf0165549_20210523214450.jpeg',4,'','','2021-05-23 21:38:04','2021-05-23 21:44:50','2006-01-02 15:04:05'),(3,'三国演义','《三国演义》，又名《三国志演义》《三国志通俗演义》，这部小说描写了从汉灵帝黄巾起义写起，到西晋三国统一为止，九十余年的重大历史事件及历史人物的活动',0,'uploads/5fa4e6dbffb12be2c7f72886ae74bab4_20210523214429.jpeg',3,'','','2021-05-23 21:39:41','2021-05-23 21:44:29','2006-01-02 15:04:05'),(4,'水浒传','《水浒传》是一部长篇英雄传奇，是中国古代长篇小说的代表作之一，是以宋江起义故事为线索创作出来的。宋江起义发生在北宋徽宗时期，',9,'uploads/bab9a317948c81793e98469117af3c3c_20210523214527.jpeg',21,'','','2021-05-23 21:40:59','2021-05-23 21:45:27','2006-01-02 15:04:05'),(5,'九阳神功','九阳神功，是《九阳真经》中多种学术理论体系中的武学体系，出自金庸小说《倚天屠龙记》。而在金庸的《神雕侠侣》结尾处，由少林派觉远',0,'uploads/8907b57b406a790f508ee1ebc853625c_20210523214359.jpeg',0,'','','2021-05-23 21:44:00','2021-05-23 21:47:37','2006-01-02 15:04:05'),(6,'乾坤大挪移','乾坤大挪移是武侠小说《倚天屠龙记》中张无忌修习的武功，在中原明教总坛昆仑山光明顶的禁地中习得，功法源自波斯明教，乃镇教之宝。',1,'uploads/c527e761732892ae9997a26a0fbf7236_20210523214702.jpeg',1,'','','2021-05-23 21:47:02','2021-05-23 21:49:31','2006-01-02 15:04:05');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `readers`
--

DROP TABLE IF EXISTS `readers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `readers` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户账号',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密码',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '手机',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户邮箱',
  `status` int DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '最后登录时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'Empty string' COMMENT '备注',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `readers`
--

LOCK TABLES `readers` WRITE;
/*!40000 ALTER TABLE `readers` DISABLE KEYS */;
INSERT INTO `readers` VALUES (1,'liming','e10adc3949ba59abbe56e057f20f883e','','','361018729@qq.com',0,'192.168.1.2','2021-05-23 21:53:13','','2021-05-22 11:42:16','2021-05-23 21:53:13','2006-01-02 15:04:05');
/*!40000 ALTER TABLE `readers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_apis`
--

DROP TABLE IF EXISTS `system_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05',
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '请求方法',
  PRIMARY KEY (`id`),
  KEY `idx_system_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_apis`
--

LOCK TABLES `system_apis` WRITE;
/*!40000 ALTER TABLE `system_apis` DISABLE KEYS */;
INSERT INTO `system_apis` VALUES (1,'2021-04-23 20:56:12','2021-05-04 01:25:11','2021-05-04 09:25:11','/admin/updateadminpassword','修改密码','admin','PUT'),(2,'2021-04-23 20:56:18','2021-05-04 01:25:17','2021-05-04 09:25:17','/admin/logout','退出登录','admin','POST'),(3,'2021-04-23 22:43:51','2021-05-04 01:25:05','2021-05-04 09:25:05','/admin/info','-获取信息','admin','GET'),(4,'2021-04-24 04:05:11','2021-04-24 04:19:43','2006-01-02 15:04:05','/admin/delete','删除','admin','DELETE'),(5,'2021-04-24 04:07:24','2021-04-24 04:19:22','2006-01-02 15:04:05','/admin/list','列表','admin','POST'),(6,'2021-04-24 04:20:13','2021-04-24 04:20:13','2006-01-02 15:04:05','/admin/deleteBatch','删除批量','admin','DELETE'),(7,'2021-04-24 04:20:40','2021-04-24 04:20:40','2006-01-02 15:04:05','/admin/find','查询','admin','POST'),(8,'2021-04-24 04:20:59','2021-04-24 04:20:59','2006-01-02 15:04:05','/admin/adminAdd','添加','admin','POST'),(9,'2021-04-24 04:21:16','2021-04-24 04:21:16','2006-01-02 15:04:05','/admin/adminUpdate','修改','admin','PUT'),(10,'2021-04-24 04:22:13','2021-04-24 04:22:13','2006-01-02 15:04:05','/systemMenu/list','列表','systemMenu','POST'),(11,'2021-04-24 04:22:31','2021-04-24 04:22:31','2006-01-02 15:04:05','/systemMenu/delete','删除','systemMenu','DELETE'),(12,'2021-04-24 04:22:53','2021-04-24 04:22:53','2006-01-02 15:04:05','/systemMenu/deleteBatch','删除批量','systemMenu','DELETE'),(13,'2021-04-24 04:23:21','2021-04-24 04:23:21','2006-01-02 15:04:05','/systemMenu/find','查询','systemMenu','POST'),(14,'2021-04-24 04:23:35','2021-04-24 04:23:35','2006-01-02 15:04:05','/systemMenu/add','添加','systemMenu','POST'),(15,'2021-04-24 04:23:54','2021-04-24 04:23:54','2006-01-02 15:04:05','/systemMenu/update','修改','systemMenu','PUT'),(16,'2021-04-24 04:28:06','2021-04-24 04:28:06','2006-01-02 15:04:05','/systemApi/list','列表','systemApi','POST'),(17,'2021-04-24 04:28:06','2021-04-24 04:28:06','2006-01-02 15:04:05','/systemApi/delete','删除','systemApi','DELETE'),(18,'2021-04-24 04:28:06','2021-04-24 04:31:36','2006-01-02 15:04:05','/systemApi/deleteBatch','删除批量','systemApi','DELETE'),(19,'2021-04-24 04:28:06','2021-04-24 04:28:06','2006-01-02 15:04:05','/systemApi/find','查询','systemApi','POST'),(20,'2021-04-24 04:28:06','2021-04-24 04:28:06','2006-01-02 15:04:05','/systemApi/add','添加','systemApi','POST'),(21,'2021-04-24 04:28:06','2021-04-24 04:28:06','2006-01-02 15:04:05','/systemApi/update','修改','systemApi','PUT'),(22,'2021-04-24 04:30:40','2021-04-24 04:30:40','2006-01-02 15:04:05','/systemRole/list','列表','systemRole','POST'),(23,'2021-04-24 04:30:40','2021-04-24 04:30:40','2006-01-02 15:04:05','/systemRole/delete','删除','systemRole','DELETE'),(24,'2021-04-24 04:30:40','2021-04-24 04:31:37','2006-01-02 15:04:05','/systemRole/deleteBatch','删除批量','systemRole','DELETE'),(25,'2021-04-24 04:30:40','2021-04-24 04:30:40','2006-01-02 15:04:05','/systemRole/find','查询','systemRole','POST'),(26,'2021-04-24 04:30:40','2021-04-24 04:30:40','2006-01-02 15:04:05','/systemRole/add','添加','systemRole','POST'),(27,'2021-04-24 04:30:40','2021-04-24 04:30:40','2006-01-02 15:04:05','/systemRole/update','修改','systemRole','PUT'),(28,'2021-04-24 04:31:04','2021-04-24 04:31:04','2006-01-02 15:04:05','/systemDepartment/list','列表','systemDepartment','POST'),(29,'2021-04-24 04:31:04','2021-04-24 04:31:04','2006-01-02 15:04:05','/systemDepartment/delete','删除','systemDepartment','DELETE'),(30,'2021-04-24 04:31:04','2021-04-24 04:31:39','2006-01-02 15:04:05','/systemDepartment/deleteBatch','删除批量','systemDepartment','DELETE'),(31,'2021-04-24 04:31:04','2021-04-24 04:31:04','2006-01-02 15:04:05','/systemDepartment/find','查询','systemDepartment','POST'),(32,'2021-04-24 04:31:04','2021-04-24 04:31:04','2006-01-02 15:04:05','/systemDepartment/add','添加','systemDepartment','POST'),(33,'2021-04-24 04:31:04','2021-04-24 04:31:04','2006-01-02 15:04:05','/systemDepartment/update','修改','systemDepartment','PUT'),(34,'2021-04-27 16:23:30','2021-04-27 16:23:30','2006-01-02 15:04:05','/systemDepartment/parentList','获取一级部门','systemDepartment','POST'),(35,'2021-04-27 16:24:59','2021-05-02 14:17:53','2006-01-02 15:04:05','/systemRoleApi/byRoleId','-获取角色iD\'sApi','systemRoleApi','POST'),(36,'2021-04-27 16:29:25','2021-04-27 16:29:25','2006-01-02 15:04:05','/systemMenu/treeList','获取菜单树[权限分配]','systemMenu','POST'),(37,'2021-04-27 16:30:27','2021-04-27 16:30:27','2006-01-02 15:04:05','/systemRoleApi/add','角色API添加','systemRoleApi','POST'),(38,'2021-04-27 16:32:45','2021-05-02 14:17:59','2006-01-02 15:04:05','/systemRoleMenu/byRoleId','-获取Role\'sMenu','systemRoleMenu','POST'),(39,'2021-04-27 16:33:47','2021-04-27 16:33:47','2006-01-02 15:04:05','/systemRoleMenu/add','Role\'sMenu添加','systemRoleMenu','POST'),(40,'2021-04-27 16:35:38','2021-04-27 16:35:38','2006-01-02 15:04:05','/systemMenu/parentList','获取一级菜单','systemMenu','POST'),(41,'2021-04-27 16:36:35','2021-04-27 16:36:35','2006-01-02 15:04:05','/systemRole/parentList','获取一级角色','systemRole','POST'),(42,'2021-05-02 13:53:11','2021-05-02 13:54:00','2021-05-02 21:53:59','bb','bb','bb','POST'),(43,'2021-05-02 13:53:22','2021-05-02 13:54:15','2021-05-02 21:54:15','bb1','bb1','bb','PUT'),(44,'2021-05-02 13:53:30','2021-05-02 13:54:15','2021-05-02 21:54:15','bb2','bb2','bb','DELETE'),(45,'2021-05-02 13:53:39','2021-05-02 13:54:15','2021-05-02 21:54:15','bb4123123','bb4','bb','GET'),(52,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/add','添加','book','POST'),(53,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/update','修改','book','PUT'),(54,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/find','查询','book','POST'),(55,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/delete','删除','book','DELETE'),(56,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/deleteBatch','删除批量','book','DELETE'),(57,'2021-05-18 17:40:32','2021-05-18 17:40:32','2006-01-02 15:04:05','/book/list','列表','book','POST'),(59,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/add','添加','bookOrder','POST'),(60,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/update','修改','bookOrder','PUT'),(61,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/find','查询','bookOrder','POST'),(62,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/delete','删除','bookOrder','DELETE'),(63,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/deleteBatch','删除批量','bookOrder','DELETE'),(64,'2021-05-19 15:17:10','2021-05-19 15:17:10','2006-01-02 15:04:05','/bookOrder/list','列表','bookOrder','POST'),(71,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/add','添加','reader','POST'),(72,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/update','修改','reader','PUT'),(73,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/find','查询','reader','POST'),(74,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/delete','删除','reader','DELETE'),(75,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/deleteBatch','删除批量','reader','DELETE'),(76,'2021-05-19 16:10:04','2021-05-19 16:10:04','2006-01-02 15:04:05','/reader/list','列表','reader','POST'),(77,'2021-05-22 01:23:55','2021-05-22 01:23:55','2006-01-02 15:04:05','/verifyCode/add','添加','verifyCode','POST'),(78,'2021-05-22 01:23:56','2021-05-22 01:23:56','2006-01-02 15:04:05','/verifyCode/update','修改','verifyCode','PUT'),(79,'2021-05-22 01:23:56','2021-05-22 01:23:56','2006-01-02 15:04:05','/verifyCode/find','查询','verifyCode','POST'),(80,'2021-05-22 01:23:56','2021-05-22 01:23:56','2006-01-02 15:04:05','/verifyCode/delete','删除','verifyCode','DELETE'),(81,'2021-05-22 01:23:57','2021-05-22 01:23:57','2006-01-02 15:04:05','/verifyCode/deleteBatch','删除批量','verifyCode','DELETE'),(82,'2021-05-22 01:23:57','2021-05-22 01:23:57','2006-01-02 15:04:05','/verifyCode/list','列表','verifyCode','POST');
/*!40000 ALTER TABLE `system_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_departments`
--

DROP TABLE IF EXISTS `system_departments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '祖级列表',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '部门名称',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '邮箱',
  `status` int DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '创建者',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '更新者',
  `sort` int DEFAULT '0' COMMENT '排序',
  `parent_id` int DEFAULT '0' COMMENT '父级ID',
  PRIMARY KEY (`id`),
  KEY `idx_system_departments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_departments`
--

LOCK TABLES `system_departments` WRITE;
/*!40000 ALTER TABLE `system_departments` DISABLE KEYS */;
INSERT INTO `system_departments` VALUES (4,'2021-04-24 00:52:57','2021-04-24 14:35:33','2006-01-02 15:04:05','1','成都分部','1','1','1',1,'1','1',123,9),(5,'2021-04-24 01:06:25','2021-04-24 14:59:01','2006-01-02 15:04:05','12','武汉分部','135212343873','12','11',12,'361018729@qq.com','12',12,9),(6,'2021-04-24 01:12:12','2021-04-25 09:06:57','2006-01-02 15:04:05','123,123,12,321,3,213,12,3','深圳分部','张明','1352345677','12345729@qq.com',0,'李节节','李明',123,0),(7,'2021-04-24 01:47:29','2021-04-24 14:59:09','2006-01-02 15:04:05','12','广州分部','145134893873','12','12',12,'361018729@qq.com','12',12,0),(8,'2021-04-24 02:17:40','2021-04-24 14:58:37','2006-01-02 15:04:05','1123,1111','北京分部','liming11','135248235','361012312318729@qq.com',0,'zhangshan123','wangwu',1,9),(9,'2021-04-24 02:26:22','2021-04-24 14:58:23','2006-01-02 15:04:05','123','总部','23','13524891234','123@qq.com',1,'1','2',0,0),(10,'2021-04-24 06:03:02','2021-04-24 14:58:53','2006-01-02 15:04:05','1,1,1,1,1,1','上海分部','李明','13512345123','34324@a.com',0,'111','122',0,9),(11,'2021-05-02 13:38:57','2021-05-02 13:40:12','2021-05-02 21:40:12','','aaa','bbb','123123123','123123123',0,'','',0,0),(12,'2021-05-02 13:39:10','2021-05-02 13:40:12','2021-05-02 21:40:12','','aaaaaa23123231','123123123','123123','123123123',0,'','',0,0),(13,'2021-05-02 13:39:19','2021-05-02 13:39:33','2021-05-02 21:39:33','','a12312312','123123','123123','123123',0,'','',0,0);
/*!40000 ALTER TABLE `system_departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_menus`
--

DROP TABLE IF EXISTS `system_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_menus` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加日期',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除日期',
  `parent_id` int DEFAULT '0' COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` int DEFAULT '0' COMMENT '1是0否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端vue文件路径',
  `sort` int DEFAULT '0' COMMENT '排序标记',
  `title` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单名称',
  `icon` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '图标',
  PRIMARY KEY (`id`),
  KEY `idx_system_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_menus`
--

LOCK TABLES `system_menus` WRITE;
/*!40000 ALTER TABLE `system_menus` DISABLE KEYS */;
INSERT INTO `system_menus` VALUES (11,'2021-04-23 13:31:02','2021-05-02 15:31:54','2006-01-02 15:04:05',15,'adminList','adminList',0,'view/superAdmin/systemUser/index',1,'管理员管理','user-solid'),(12,'2021-04-24 06:03:33','2021-05-02 15:34:42','2006-01-02 15:04:05',15,'systemApiList','systemApiList',0,'view/superAdmin/systemApi/index',3,'Api管理','s-tools'),(13,'2021-04-24 11:39:33','2021-05-02 15:31:43','2006-01-02 15:04:05',15,'systemMenuList','systemMenuList',0,'view/superAdmin/systemMenu/index',2,'菜单管理','menu'),(14,'2021-04-24 11:44:54','2021-05-02 15:34:07','2006-01-02 15:04:05',15,'systemRoleList','systemRoleList',0,'view/superAdmin/systemRole/index',4,'角色管理','s-tools'),(15,'2021-04-24 11:50:20','2021-05-24 01:09:31','2006-01-02 15:04:05',0,'admin','admin',0,'view/superAdmin/index',0,'超级管理员','suitcase'),(16,'2021-04-24 12:57:27','2021-05-02 15:33:58','2006-01-02 15:04:05',15,'systemDepartmentList','systemDepartmentList',0,'view/superAdmin/systemDepartment/index',5,'部门管理','s-tools'),(17,'2021-04-24 13:24:29','2021-05-02 11:42:33','2006-01-02 15:04:05',0,'dashboard','dashboard',0,'view/dashboard/index',-1,'Dashboard','s-home'),(18,'2021-05-04 20:34:16','2021-05-24 01:10:48','2006-01-02 15:04:05',0,'k8s','k8s',0,'view/superAdmin/index',2,'K8S管理','s-grid'),(19,'2021-05-04 20:36:52','2021-05-05 10:28:31','2006-01-02 15:04:05',18,'deployment','deployment',0,'view/k8s/deployment',0,'Deployment','star-on'),(20,'2021-05-05 10:18:20','2021-05-05 11:08:35','2006-01-02 15:04:05',18,'namespace','namespace',0,'view/k8s/namespace',1,'Namespace','star-on'),(21,'2021-05-05 10:18:20','2021-05-05 11:09:04','2006-01-02 15:04:05',18,'service','service',0,'view/k8s/service',2,'Service','star-on'),(22,'2021-05-05 10:18:20','2021-05-05 11:09:27','2006-01-02 15:04:05',18,'ingress','ingress',0,'view/k8s/ingress',3,'Ingress','star-on'),(23,'2021-05-05 10:18:20','2021-05-07 01:54:36','2006-01-02 15:04:05',18,'pods','pods',0,'view/k8s/pods',3,'Pods','star-on'),(24,'2021-05-18 16:27:05','2021-05-24 01:11:21','2006-01-02 15:04:05',0,'bookstore','bookstore',0,'view/superAdmin/index',4,'书店管理','bell'),(25,'2021-05-18 16:28:42','2021-05-18 16:28:42','2006-01-02 15:04:05',24,'books','books',0,'view/bookstore/book/list',1,'图书管理','menu'),(26,'2021-05-19 15:27:05','2021-05-19 15:27:05','2006-01-02 15:04:05',24,'bookOrder','bookOrder',0,'view/bookstore/book/ordersList',1,'图书借阅记录','menu'),(27,'2021-05-19 15:57:14','2021-05-19 16:17:16','2006-01-02 15:04:05',24,'readers','readers',0,'view/bookstore/book/readerList',3,'用户管理','menu'),(28,'2021-05-22 01:31:03','2021-05-22 01:45:58','2006-01-02 15:04:05',24,'verifyCodeList','verifyCodeList',0,'view/bookstore/book/verifyCodeList',4,'验证码','menu');
/*!40000 ALTER TABLE `system_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_apis`
--

DROP TABLE IF EXISTS `system_role_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role_apis` (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色API关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_apis`
--

LOCK TABLES `system_role_apis` WRITE;
/*!40000 ALTER TABLE `system_role_apis` DISABLE KEYS */;
INSERT INTO `system_role_apis` VALUES ('p','4','/admin/adminUpdate','PUT','','',''),('p','4','/admin/adminAdd','POST','','',''),('p','4','/admin/find','POST','','',''),('p','1','/systemRole/list','POST','','',''),('p','6','/admin/adminUpdate','PUT','','',''),('p','6','/admin/adminAdd','POST','','',''),('p','6','/admin/find','POST','','',''),('p','6','/admin/deleteBatch','DELETE','','',''),('p','6','/admin/list','POST','','',''),('p','6','/admin/delete','DELETE','','',''),('p','6','/admin/info','GET','','',''),('p','6','/admin/logout','POST','','',''),('p','6','/admin/updateadminpassword','PUT','','',''),('p','6','/systemApi/add','POST','','',''),('p','6','/systemApi/find','POST','','',''),('p','6','/systemApi/deleteBatch','DELETE','','',''),('p','6','/systemApi/delete','DELETE','','',''),('p','8','/systemApi/update','PUT','','',''),('p','8','/systemApi/add','POST','','',''),('p','8','/systemApi/find','POST','','',''),('p','8','/systemApi/deleteBatch','DELETE','','',''),('p','8','/systemApi/delete','DELETE','','',''),('p','8','/systemApi/list','POST','','',''),('p','5','/admin/info','GET','','',''),('p','5','/systemDepartment/parentList','POST','','',''),('p','5','/systemDepartment/update','PUT','','',''),('p','5','/systemDepartment/add','POST','','',''),('p','5','/systemDepartment/find','POST','','',''),('p','5','/systemDepartment/deleteBatch','DELETE','','',''),('p','5','/systemDepartment/delete','DELETE','','',''),('p','5','/systemDepartment/list','POST','','',''),('p','5','/systemRole/parentList','POST','','',''),('p','5','/systemRole/update','PUT','','',''),('p','5','/systemRole/add','POST','','',''),('p','5','/systemRole/find','POST','','',''),('p','5','/systemRole/deleteBatch','DELETE','','',''),('p','5','/systemRole/delete','DELETE','','',''),('p','5','/systemRole/list','POST','','',''),('p','5','/systemRoleApi/byRoleId','POST','','',''),('p','5','/systemRoleMenu/byRoleId','POST','','',''),('p','7','/admin/adminUpdate','PUT','','',''),('p','7','/admin/adminAdd','POST','','',''),('p','7','/admin/find','POST','','',''),('p','7','/admin/deleteBatch','DELETE','','',''),('p','7','/admin/list','POST','','',''),('p','7','/admin/delete','DELETE','','',''),('p','7','/book/list','POST','','',''),('p','7','/book/deleteBatch','DELETE','','',''),('p','7','/book/delete','DELETE','','',''),('p','7','/book/find','POST','','',''),('p','7','/book/update','PUT','','',''),('p','7','/book/add','POST','','',''),('p','7','/bookOrder/list','POST','','',''),('p','7','/bookOrder/deleteBatch','DELETE','','',''),('p','7','/bookOrder/delete','DELETE','','',''),('p','7','/bookOrder/find','POST','','',''),('p','7','/bookOrder/update','PUT','','',''),('p','7','/bookOrder/add','POST','','',''),('p','7','/reader/list','POST','','',''),('p','7','/reader/deleteBatch','DELETE','','',''),('p','7','/reader/delete','DELETE','','',''),('p','7','/reader/find','POST','','',''),('p','7','/reader/update','PUT','','',''),('p','7','/reader/add','POST','','',''),('p','7','/systemApi/update','PUT','','',''),('p','7','/systemApi/add','POST','','',''),('p','7','/systemApi/find','POST','','',''),('p','7','/systemApi/deleteBatch','DELETE','','',''),('p','7','/systemApi/delete','DELETE','','',''),('p','7','/systemApi/list','POST','','',''),('p','7','/systemDepartment/parentList','POST','','',''),('p','7','/systemDepartment/update','PUT','','',''),('p','7','/systemDepartment/add','POST','','',''),('p','7','/systemDepartment/find','POST','','',''),('p','7','/systemDepartment/deleteBatch','DELETE','','',''),('p','7','/systemDepartment/delete','DELETE','','',''),('p','7','/systemDepartment/list','POST','','',''),('p','7','/systemMenu/parentList','POST','','',''),('p','7','/systemMenu/treeList','POST','','',''),('p','7','/systemMenu/update','PUT','','',''),('p','7','/systemMenu/add','POST','','',''),('p','7','/systemMenu/find','POST','','',''),('p','7','/systemMenu/deleteBatch','DELETE','','',''),('p','7','/systemMenu/delete','DELETE','','',''),('p','7','/systemMenu/list','POST','','',''),('p','7','/systemRole/parentList','POST','','',''),('p','7','/systemRole/update','PUT','','',''),('p','7','/systemRole/add','POST','','',''),('p','7','/systemRole/find','POST','','',''),('p','7','/systemRole/deleteBatch','DELETE','','',''),('p','7','/systemRole/delete','DELETE','','',''),('p','7','/systemRole/list','POST','','',''),('p','7','/systemRoleApi/add','POST','','',''),('p','7','/systemRoleApi/byRoleId','POST','','',''),('p','7','/systemRoleMenu/add','POST','','',''),('p','7','/systemRoleMenu/byRoleId','POST','','',''),('p','7','/verifyCode/list','POST','','',''),('p','7','/verifyCode/deleteBatch','DELETE','','',''),('p','7','/verifyCode/delete','DELETE','','',''),('p','7','/verifyCode/find','POST','','',''),('p','7','/verifyCode/update','PUT','','',''),('p','7','/verifyCode/add','POST','','','');
/*!40000 ALTER TABLE `system_role_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_menus`
--

DROP TABLE IF EXISTS `system_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role_menus` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT '0' COMMENT '角色ID',
  `menu_id` int DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=222 DEFAULT CHARSET=utf8 COMMENT='角色菜单关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_menus`
--

LOCK TABLES `system_role_menus` WRITE;
/*!40000 ALTER TABLE `system_role_menus` DISABLE KEYS */;
INSERT INTO `system_role_menus` VALUES (54,4,17),(69,6,17),(70,6,15),(71,6,11),(72,6,13),(84,8,15),(85,8,13),(86,8,12),(94,5,17),(95,5,15),(96,5,14),(97,5,16),(204,7,17),(205,7,15),(206,7,11),(207,7,13),(208,7,12),(209,7,14),(210,7,16),(211,7,18),(212,7,19),(213,7,20),(214,7,21),(215,7,23),(216,7,22),(217,7,24),(218,7,26),(219,7,25),(220,7,28),(221,7,27);
/*!40000 ALTER TABLE `system_role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_roles`
--

DROP TABLE IF EXISTS `system_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名',
  `sort` int DEFAULT '0' COMMENT '排序',
  `parent_id` int DEFAULT '0' COMMENT '父级ID',
  PRIMARY KEY (`id`),
  KEY `idx_system_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_roles`
--

LOCK TABLES `system_roles` WRITE;
/*!40000 ALTER TABLE `system_roles` DISABLE KEYS */;
INSERT INTO `system_roles` VALUES (4,'2021-04-23 23:56:39','2021-04-24 23:02:31','2006-01-02 15:04:05','销售部管理员',6,6),(5,'2021-04-23 23:57:33','2021-04-25 09:07:58','2006-01-02 15:04:05','技术部管理员',6,0),(6,'2021-04-23 23:58:44','2021-04-24 23:02:24','2006-01-02 15:04:05','总部管理员',0,0),(7,'2021-04-24 06:02:20','2021-04-24 13:11:51','2006-01-02 15:04:05','超级管理员',0,0),(8,'2021-05-02 14:05:23','2021-05-02 14:06:13','2006-01-02 15:04:05','aa',0,6),(9,'2021-05-02 14:05:30','2021-05-02 14:06:14','2006-01-02 15:04:05','aa1',0,0),(10,'2021-05-02 14:05:34','2021-05-02 14:06:15','2006-01-02 15:04:05','aa234324324',0,0);
/*!40000 ALTER TABLE `system_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_users`
--

DROP TABLE IF EXISTS `system_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_users` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` int DEFAULT '0' COMMENT '部门ID',
  `role_id` int DEFAULT '0' COMMENT '角色ID',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
  `user_type` int DEFAULT '0' COMMENT '用户类型（0系统用户）',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '手机号码',
  `sex` int DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密码',
  `status` int DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` int DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '最后登录时间',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '创建者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'Empty string' COMMENT '备注',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '更新者',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_users`
--

LOCK TABLES `system_users` WRITE;
/*!40000 ALTER TABLE `system_users` DISABLE KEYS */;
INSERT INTO `system_users` VALUES (1,9,7,'admin','张四',1,'1234729@qq.com','12345677',1,'bf0812cf724eef1e27fb4c3946d8b05f_20210519015947.jpeg','e10adc3949ba59abbe56e057f20f883e',0,0,'127.0.0.1','2021-05-24 01:51:47','','remarks1','','2021-04-24 06:34:04','2021-05-24 01:51:46','2006-01-02 15:04:05'),(2,10,6,'root','张三',0,'','',0,'','96e79218965eb72c92a549dd5a330112',0,0,'','2006-01-02 15:04:05','','root','','2021-04-24 06:41:34','2021-05-01 15:44:46','2006-01-02 15:04:05'),(3,8,5,'admin3','李明',0,'','13512341234',0,'','e10adc3949ba59abbe56e057f20f883e',0,0,'127.0.0.1:61484','2021-04-27 09:24:28','','','','2021-04-24 10:58:13','2021-05-01 15:44:37','2006-01-02 15:04:05'),(4,10,5,'liming','',0,'','13512341234',0,'04dd9c5a7fb2e5a554cced10089ed6de_20210504092028.png','96e79218965eb72c92a549dd5a330112',0,0,'127.0.0.1:64314','2021-05-16 01:59:10','','23','','2021-04-27 16:34:49','2021-05-15 17:59:10','2006-01-02 15:04:05'),(5,10,4,'aaa','',0,'','',0,'','e10adc3949ba59abbe56e057f20f883e',0,0,'','2006-01-02 15:04:05','','aaa','','2021-05-02 13:30:10','2021-05-02 13:30:25','2021-05-02 21:30:25'),(6,10,7,'admin1','',0,'','',0,'','e10adc3949ba59abbe56e057f20f883e',0,0,'','2006-01-02 15:04:05','','','','2021-05-02 13:31:49','2021-05-02 13:32:23','2021-05-02 21:32:23'),(7,0,0,'aaaadfaf','',0,'','',0,'','2cecbfb14b16ad5b3a953748733c254b',0,0,'','2006-01-02 15:04:05','','123123','','2021-05-02 13:32:05','2021-05-02 13:32:23','2021-05-02 21:32:23');
/*!40000 ALTER TABLE `system_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `verify_codes`
--

DROP TABLE IF EXISTS `verify_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `verify_codes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT '2006-01-02 15:04:05',
  `account` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '号码(手机或邮箱)',
  `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '验证码',
  `type` int DEFAULT '0' COMMENT '类型0手机1邮箱',
  `status` int DEFAULT '0' COMMENT '状态0未验证1已验证2验证错误',
  PRIMARY KEY (`id`),
  KEY `idx_system_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='验证码表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `verify_codes`
--

LOCK TABLES `verify_codes` WRITE;
/*!40000 ALTER TABLE `verify_codes` DISABLE KEYS */;
INSERT INTO `verify_codes` VALUES (1,'2021-05-22 11:42:32','2021-05-22 11:42:16','2006-01-02 15:04:05','361018729@qq.com','3777',1,1),(2,'2021-05-22 12:26:00','2021-05-22 12:28:05','2006-01-02 15:04:05','361018729@qq.com','8522',1,1),(3,'2021-05-22 12:26:18','2021-05-22 12:26:20','2006-01-02 15:04:05','361018729@qq.com','2227',1,1);
/*!40000 ALTER TABLE `verify_codes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-05-24 14:05:45

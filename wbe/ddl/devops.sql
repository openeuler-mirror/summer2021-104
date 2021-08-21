-- MySQL dump 10.13  Distrib 5.7.34, for osx10.16 (x86_64)
--
-- Host: 127.0.0.1    Database: devops
-- ------------------------------------------------------
-- Server version	5.7.34

--
-- Table structure for table `host_management`
--

DROP TABLE IF EXISTS `host_management`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `host_management` (
  `host_management_id` int(11) NOT NULL AUTO_INCREMENT,
  `host_management_ip` varchar(255) DEFAULT NULL,
  `host_management_name` varchar(255) NOT NULL,
  `host_management_user` varchar(255) NOT NULL,
  `host_management_password` varchar(255) DEFAULT NULL,
  `host_management_key` varchar(255) DEFAULT NULL,
  `host_management_connect_method` varchar(255) DEFAULT NULL,
  `host_management_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `host_management_modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`host_management_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `host_management`
--

LOCK TABLES `host_management` WRITE;
/*!40000 ALTER TABLE `host_management` DISABLE KEYS */;
INSERT INTO `host_management` VALUES (1,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 13:35:32','2021-08-10 05:35:32'),(2,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 13:59:55','2021-08-10 05:59:55'),(3,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:39','2021-08-10 06:01:39'),(4,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:40','2021-08-10 06:01:40'),(5,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:41','2021-08-10 06:01:41'),(6,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:42','2021-08-10 06:01:42'),(7,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:42','2021-08-10 06:01:42'),(8,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 14:01:43','2021-08-10 06:01:43'),(12,'127.0.0.1','localhost','root','123456','NULL','KEY','2021-08-10 18:02:18','2021-08-10 10:02:18'),(13,'','','','','','','2021-08-10 18:04:54','2021-08-10 10:04:54');
/*!40000 ALTER TABLE `host_management` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_username` varchar(255) DEFAULT NULL,
  `user_password` varchar(255) NOT NULL,
  `user_realname` varchar(255) NOT NULL,
  `user_department` varchar(255) DEFAULT NULL,
  `user_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'ops','6729f0ce06d29e26760ac56aa24c162910bf9f653f6e57c1d51b16ec71ec27f5ce2b81b1868809fa5c78cb4522094983f465cccc2b006f1f5a068db0f6a3653e','运维工程师',NULL,'2021-07-20 13:20:26','2021-08-10 07:28:07'),(2,'admin','6729f0ce06d29e26760ac56aa24c162910bf9f653f6e57c1d51b16ec71ec27f5ce2b81b1868809fa5c78cb4522094983f465cccc2b006f1f5a068db0f6a3653e','系统管理员',NULL,'2021-07-20 13:20:26','2021-08-10 07:28:07');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

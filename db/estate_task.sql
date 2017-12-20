-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: estate
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task` (
  `Task_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '工作ID',
  `User_ID` bigint(11) NOT NULL COMMENT '員工ID',
  `NickName` varchar(45) NOT NULL COMMENT '員工編號',
  `CreateTime` timestamp NULL DEFAULT NULL COMMENT '建立日期',
  `UpdateTime` timestamp NULL DEFAULT NULL COMMENT '更新日期',
  `TaskName` varchar(45) NOT NULL COMMENT '工作名稱',
  `TaskDescribe` varchar(500) NOT NULL COMMENT '工作描述',
  `Memo` varchar(100) NOT NULL COMMENT '工作備忘錄',
  PRIMARY KEY (`Task_ID`),
  UNIQUE KEY `Task_ID_UNIQUE` (`Task_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
INSERT INTO `task` VALUES (31,1,'cat1112t','2017-11-27 16:39:31','2017-11-27 16:39:31','22','',''),(32,1,'cat1112t','2017-11-27 16:39:36','2017-11-27 16:39:36','33','',''),(33,1,'cat1112t','2017-11-27 16:39:40','2017-11-27 16:40:27','44','44',''),(34,1,'cat1112t','2017-11-27 16:39:43','2017-11-27 17:48:33','55','55',''),(35,1,'cat1112t','2017-11-27 16:39:46','2017-11-27 16:39:46','66','','');
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-18 22:38:33

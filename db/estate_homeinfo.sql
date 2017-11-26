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
-- Table structure for table `homeinfo`
--

DROP TABLE IF EXISTS `homeinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homeinfo` (
  `HomeID` int(11) NOT NULL AUTO_INCREMENT COMMENT '房屋ID',
  `User_ID` bigint(11) NOT NULL COMMENT '員工編號(哪個員工新增的)\n',
  `NickName` varchar(45) NOT NULL COMMENT '員工暱稱',
  `CreateTime` datetime NOT NULL COMMENT '房屋新增時間',
  `UpdateTime` datetime NOT NULL COMMENT '房屋更新時間',
  `HomeName` varchar(60) NOT NULL COMMENT '房屋名稱',
  `HomeAddress` varchar(80) NOT NULL COMMENT '房屋地址',
  `HomeAge` int(11) NOT NULL COMMENT '屋齡',
  `HomeFootage` float NOT NULL COMMENT '房屋坪數',
  `HomePrice` int(11) NOT NULL COMMENT '房屋價格',
  `Vip_rank` int(11) NOT NULL COMMENT '房屋等級 0:雅房 1:套房 2:兩房一廳 3:三房兩廳 4:工廠 5:辦公室 6:透天厝 7:豪宅',
  `Memo` varchar(45) NOT NULL,
  PRIMARY KEY (`HomeID`),
  UNIQUE KEY `HomeID_UNIQUE` (`HomeID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='紀錄每個房屋資料';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `homeinfo`
--

LOCK TABLES `homeinfo` WRITE;
/*!40000 ALTER TABLE `homeinfo` DISABLE KEYS */;
INSERT INTO `homeinfo` VALUES (1,1,'cat1111','2017-11-10 21:25:21','2017-11-23 08:27:57','鼎藏1','我住在地球',1,2,200,7,'測試的memo'),(3,1,'cat1111','2017-11-27 00:43:48','2017-11-27 00:44:29','cat-home','tw',50,100,200,6,'ccccc');
/*!40000 ALTER TABLE `homeinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-11-27  0:49:06

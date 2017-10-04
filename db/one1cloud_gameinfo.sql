-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: one1cloud
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
-- Table structure for table `gameinfo`
--

DROP TABLE IF EXISTS `gameinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gameinfo` (
  `PlatformID` int(11) NOT NULL COMMENT '紀錄有哪些 介接的第三方廠商編號   0: unkonw 1: 老虎城   2:太陽城',
  `GameID` int(11) NOT NULL COMMENT '遊戲編號  e.g.   先暫定 1001是捕魚  編碼規則在討論\n編碼規則   GameMode   + 流水號(3碼)  e.g 魚機 1 001',
  `GameName` varchar(45) NOT NULL COMMENT '遊戲的中文名稱',
  `GameEnName` varchar(3) NOT NULL COMMENT '遊戲的英文名稱  配桌系統桌號使用',
  `GameMode` tinyint(2) NOT NULL COMMENT '遊戲模式 0:魚機 1:SLOT 2:撲克 3:麻將',
  `TableDestoryMode` tinyint(2) NOT NULL COMMENT '刪除桌模式  0: unknow 1:散桌後刪除此桌資訊  2: 散桌後保留此桌資訊,等待玩家重新入桌',
  `TablePlayerMax` tinyint(2) NOT NULL COMMENT '桌內人數上限',
  PRIMARY KEY (`GameID`),
  KEY `GameID` (`GameID`),
  KEY `PlatformID` (`PlatformID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='紀錄有哪些遊戲 和 遊戲行為';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gameinfo`
--

LOCK TABLES `gameinfo` WRITE;
/*!40000 ALTER TABLE `gameinfo` DISABLE KEYS */;
INSERT INTO `gameinfo` VALUES (1,1001,'海王魚機','FH',1,1,4),(1,1002,'海王魚機2代','FH',1,2,4),(1,1003,'海王魚機','FH',1,1,4),(1,2001,'HugaSlot','HG',2,1,1),(1,2002,'HugaSlot2代','HG2',2,1,1);
/*!40000 ALTER TABLE `gameinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-10-04 11:51:52

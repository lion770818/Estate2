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
-- Table structure for table `gamelog_fish`
--

DROP TABLE IF EXISTS `gamelog_fish`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gamelog_fish` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `PlatformID` int(2) NOT NULL COMMENT '第三方平台編號',
  `LobbyID` int(3) NOT NULL COMMENT '大廳規則編號',
  `GameID` int(4) NOT NULL COMMENT '遊戲編號',
  `TableID` varchar(45) NOT NULL COMMENT '桌子編號',
  `Seat_ID` int(11) NOT NULL COMMENT '座位編號',
  `GameMode` tinyint(2) NOT NULL COMMENT '遊戲模式   遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將',
  `CreateTime` datetime NOT NULL COMMENT '建立時間',
  `User_ID` bigint(11) NOT NULL COMMENT '玩家帳號編號',
  `Account` varchar(45) NOT NULL COMMENT '玩家帳號',
  `NickName` varchar(45) NOT NULL COMMENT '玩家暱稱',
  `Round` bigint(11) NOT NULL COMMENT '在遊戲內的第幾局',
  `Before_Balance_ci` double NOT NULL COMMENT '玩家分數_投 (之前)',
  `Before_Balance_win` double NOT NULL COMMENT '玩家分數_贏 (之前)',
  `Balance_ci` double NOT NULL COMMENT '玩家分數_投 (spin之後)',
  `Balance_win` double NOT NULL COMMENT '玩家分數_贏 (spin之後)',
  `BetLevel` double NOT NULL COMMENT '玩家壓注',
  `Bet_Win` double NOT NULL COMMENT '玩家贏分',
  `Process_Status` varchar(45) NOT NULL COMMENT '玩家狀態\n\n0:unknow 1:shoot 2:hit 3:feature_shoot 4:feature_hit ',
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gamelog_fish`
--

LOCK TABLES `gamelog_fish` WRITE;
/*!40000 ALTER TABLE `gamelog_fish` DISABLE KEYS */;
INSERT INTO `gamelog_fish` VALUES (75,1,1,1001,'FH11001-0000060',0,1,'2017-09-14 14:21:20',1,'cat111','嘿嘿喵1號',1,3000,0,2900,0,100,0,'1'),(76,1,1,1001,'FH11001-0000060',0,1,'2017-09-14 14:21:27',1,'cat111','嘿嘿喵1號',2,2900,0,2800,0,100,0,'1'),(77,1,1,1001,'FH11001-0000060',0,1,'2017-09-14 14:21:30',1,'cat111','嘿嘿喵1號',3,2800,0,2700,0,100,0,'1');
/*!40000 ALTER TABLE `gamelog_fish` ENABLE KEYS */;
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

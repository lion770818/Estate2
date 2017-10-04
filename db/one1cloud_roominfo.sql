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
-- Table structure for table `roominfo`
--

DROP TABLE IF EXISTS `roominfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roominfo` (
  `PlatformID` int(11) NOT NULL COMMENT '第三方平台編號',
  `LobbyID` int(11) NOT NULL COMMENT '大廳規則編號',
  `GameID` int(11) NOT NULL COMMENT '遊戲編號',
  `RoomID` int(11) NOT NULL COMMENT '房間編號',
  `CreateTime` datetime NOT NULL COMMENT '房間建立時間',
  `UpdateTime` datetime NOT NULL COMMENT '房間更新時間',
  `GameMode` tinyint(2) NOT NULL COMMENT '開桌模式',
  `interval_bet` int(11) NOT NULL COMMENT '區間押注紀錄\n',
  `interval_bet_pt` int(11) NOT NULL COMMENT '區間押注指標\n',
  `avg_bet` int(11) NOT NULL COMMENT '平均押注',
  `progress_water` bigint(15) NOT NULL COMMENT '獎項水池[Y]',
  `progress_odds` int(11) NOT NULL COMMENT '獎項倍數[Y]',
  `progress_support` int(11) NOT NULL COMMENT '獎項貢獻[Y]',
  `wait_item_id` int(11) NOT NULL COMMENT '待中獎項[Y]',
  `wait_item_seat` int(11) NOT NULL COMMENT '待中座位[Y]',
  `wait_item_odds` int(11) NOT NULL COMMENT '待中倍數[Y]',
  `wait_item_value` int(11) NOT NULL COMMENT '待中分數[Y]',
  `win_item_id` varchar(45) NOT NULL COMMENT '座位獎項名稱',
  `win_item_bet` int(11) NOT NULL COMMENT '座位獎項押注',
  `win_item_value` int(11) NOT NULL COMMENT '座位獎項分數',
  `win_item_win` int(11) NOT NULL COMMENT '座位獎項贏分',
  `RoomInfocol` varchar(45) NOT NULL,
  PRIMARY KEY (`RoomID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roominfo`
--

LOCK TABLES `roominfo` WRITE;
/*!40000 ALTER TABLE `roominfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `roominfo` ENABLE KEYS */;
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

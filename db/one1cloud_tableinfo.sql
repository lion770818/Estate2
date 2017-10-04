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
-- Table structure for table `tableinfo`
--

DROP TABLE IF EXISTS `tableinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tableinfo` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `PlatformID` int(11) NOT NULL COMMENT '第三方平台編號',
  `LobbyID` int(11) NOT NULL COMMENT '大廳規則編號',
  `GameID` int(11) NOT NULL COMMENT '遊戲編號',
  `TableID` varchar(20) NOT NULL COMMENT '桌子編號',
  `TableArrayIdx` int(11) NOT NULL COMMENT 'TableInfoList 的索引陣列 方便快速搜尋資料',
  `TablePlayerMax` tinyint(2) NOT NULL COMMENT '桌內人數上限',
  `TablePlayerNow` tinyint(2) NOT NULL COMMENT '桌內目前人數',
  `GameMode` tinyint(2) NOT NULL COMMENT '遊戲模式   遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將',
  `TableDestoryMode` tinyint(2) NOT NULL COMMENT ' 0: unknow 1:散桌後刪除此桌資訊  2: 散桌後保留此桌資訊,等待玩家重新入桌',
  `CreateTime` datetime NOT NULL COMMENT '一般時間格式即可, 紀錄開桌時間',
  `UpdateTime` datetime NOT NULL COMMENT '一般時間格式即可, 開桌, 玩家入桌, 玩家離桌, 都更新此時間',
  `LobbyMatchID` tinyint(4) DEFAULT NULL COMMENT '大廳配桌編號, 編號相同的廳館, 才可以配桌在一起\n不知道怎樣設定就先跟 PlatformID 一樣\n等想跨平台配在一起  在調整\nif( LobbyMatchID 相同  && GameID 相同 && 人數未滿 )\n     加入桌()',
  PRIMARY KEY (`id`,`TableArrayIdx`),
  KEY `PlatformID` (`PlatformID`),
  KEY `LobbyID` (`LobbyID`),
  KEY `GameID` (`GameID`),
  KEY `TableID` (`TableID`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8 COMMENT='桌子 結構 (只記錄桌號和散桌行為)';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tableinfo`
--

LOCK TABLES `tableinfo` WRITE;
/*!40000 ALTER TABLE `tableinfo` DISABLE KEYS */;
INSERT INTO `tableinfo` VALUES (62,1,1,1001,'FH11001-0000062',0,4,1,1,1,'2017-09-14 14:35:47','2017-09-14 23:16:20',1),(63,1,8,2001,'HG12001-0000063',1,1,1,2,1,'2017-09-14 14:40:52','2017-09-14 23:16:20',100);
/*!40000 ALTER TABLE `tableinfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-10-04 11:51:51

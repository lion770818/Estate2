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
-- Table structure for table `seatinfo`
--

DROP TABLE IF EXISTS `seatinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seatinfo` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '流水號的id  自動遞增, 作主KEY使用',
  `PlatformID` int(11) NOT NULL COMMENT '第三方平台編號',
  `LobbyID` int(11) NOT NULL COMMENT '大廳規則編號',
  `GameID` int(11) NOT NULL COMMENT '遊戲編號',
  `TableID` varchar(20) NOT NULL COMMENT '桌子編號',
  `Seat_ID` int(11) NOT NULL COMMENT '座位編號',
  `GameMode` tinyint(2) NOT NULL COMMENT '遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將 ',
  `CreateTime` datetime NOT NULL COMMENT '一般時間格式即可, 紀錄開桌時間',
  `UpdateTime` datetime NOT NULL COMMENT '一般時間格式即可, 開桌, 玩家入桌, 玩家離桌, 都更新此時間',
  `User_ID` bigint(11) NOT NULL COMMENT 'Serial Number 一個流水號, 當新增一個帳號自動加1, 帳號唯一號碼',
  `Account` varchar(45) NOT NULL COMMENT '魚機帳號',
  `NickName` varchar(45) NOT NULL COMMENT '魚機玩家暱稱',
  `Balance_ci` double NOT NULL COMMENT '單位 人民幣/浮點數  ( 從平台帶過來的錢',
  `Balance_win` double NOT NULL COMMENT '玩家贏的錢   win 先扣  隨遊戲不斷變動',
  `BetLevel` double NOT NULL COMMENT '單一押注金額',
  `Interval_bet` double NOT NULL,
  `Interval_bet_pt` double NOT NULL,
  `Avg_bet` double NOT NULL,
  `Progress_water` double NOT NULL COMMENT '獎項數規劃中',
  `Progress_odds` double NOT NULL,
  `Progress_support` double NOT NULL,
  `Wait_item_id` varchar(45) NOT NULL,
  `Wait_item_seat` varchar(45) NOT NULL,
  `Wait_item_odds` double NOT NULL,
  `Wait_item_value` double NOT NULL,
  `Win_item_id` varchar(4) NOT NULL,
  `Win_item_bet` double NOT NULL,
  `Win_item_value` double NOT NULL,
  `Win_item_win` double NOT NULL,
  PRIMARY KEY (`id`),
  KEY `TableID` (`TableID`),
  KEY `User_ID` (`User_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8 COMMENT='座位資料表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `seatinfo`
--

LOCK TABLES `seatinfo` WRITE;
/*!40000 ALTER TABLE `seatinfo` DISABLE KEYS */;
INSERT INTO `seatinfo` VALUES (68,1,1,1001,'FH11001-0000062',0,1,'2017-09-14 14:35:47','2017-09-14 14:35:47',1,'cat111','嘿嘿喵1號',3000,0,0,0,0,0,0,0,0,'','',0,0,'',0,0,0),(70,1,8,2001,'HG12001-0000063',0,2,'2017-09-14 14:40:52','2017-09-14 14:41:38',2,'cat222','嘿嘿喵2號',3000,400,0,0,0,0,0,0,0,'','',0,0,'',0,0,0);
/*!40000 ALTER TABLE `seatinfo` ENABLE KEYS */;
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

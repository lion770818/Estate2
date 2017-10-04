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
-- Table structure for table `memberinfo`
--

DROP TABLE IF EXISTS `memberinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `memberinfo` (
  `User_ID` bigint(11) NOT NULL COMMENT '玩家帳號編號\nSerial Number 一個流水號, 當新增一個帳號自動加1, 帳號唯一號碼',
  `PlatformID` int(11) unsigned NOT NULL COMMENT '第三方平台編號',
  `IP` varchar(40) NOT NULL COMMENT '玩家IP Address',
  `MacAddress` varchar(50) DEFAULT NULL COMMENT '玩家網卡 Address',
  `CreateTime` datetime NOT NULL COMMENT '帳號建立時間',
  `LoginTime` datetime NOT NULL COMMENT '帳號登入時間',
  `UpdateTime` datetime NOT NULL COMMENT '帳號更新時間',
  `Account` varchar(45) NOT NULL COMMENT '魚機帳號',
  `Password` varchar(45) NOT NULL COMMENT '魚機密碼',
  `NickName` varchar(45) NOT NULL COMMENT '魚機玩家暱稱',
  `DeviceID` int(11) DEFAULT NULL COMMENT '玩家裝置編號',
  `PhoneNumber` varchar(10) DEFAULT NULL COMMENT '玩家電話號碼',
  `Balance` double NOT NULL COMMENT '玩家的錢\n單位 人民幣/浮點數  ( 從平台帶過來的錢)',
  `Balance_ci` double NOT NULL COMMENT '玩家分數_投 (玩家帶進房間的錢)  \n單位 人民幣/浮點數 ',
  `Balance_wi` double NOT NULL COMMENT '玩家分數_贏\n\n玩家贏的錢',
  `Total_bet` int(11) NOT NULL COMMENT '總押注',
  `Total_win` bigint(15) NOT NULL COMMENT '總贏分',
  `Interval_bet` int(11) NOT NULL COMMENT '區間押注',
  `Interval_win` int(11) NOT NULL COMMENT '區間贏分',
  `Pi` int(11) NOT NULL COMMENT '玩家喜好',
  `Portable_water` int(11) NOT NULL COMMENT '攜帶水位',
  `Portable_support` int(11) NOT NULL COMMENT '攜帶貢獻',
  `Status` tinyint(4) NOT NULL COMMENT '玩家狀態  0: 登出 1:登入 2:斷線中 3:斷線連回中',
  `Vip_rank` tinyint(4) NOT NULL COMMENT 'vip位階',
  PRIMARY KEY (`User_ID`,`Account`),
  KEY `PlatformID` (`PlatformID`),
  KEY `Account` (`Account`),
  KEY `Password` (`Password`),
  KEY `User_ID` (`User_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='紀錄每個會員基本資料';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `memberinfo`
--

LOCK TABLES `memberinfo` WRITE;
/*!40000 ALTER TABLE `memberinfo` DISABLE KEYS */;
INSERT INTO `memberinfo` VALUES (1,1,'192.168.1.1','11:22:33:44:55','2017-07-21 14:28:39','2017-09-16 08:05:14','2017-09-16 08:05:14','cat111','1234','嘿嘿喵1號',0,'0955645161',826100,10,0,0,0,0,0,0,0,0,1,0),(2,1,'192.168.1.2','11:22:33:44:55','2017-07-21 14:28:39','2017-09-14 14:37:09','2017-09-14 14:40:52','cat222','1234','嘿嘿喵2號',0,'1234567890',1939450,20,0,0,0,0,0,0,0,0,1,0),(3,1,'192.168.1.2','11:22:33:44:55','2017-07-21 14:28:39','2017-09-13 18:25:25','2017-09-13 19:02:42','cat333','1234','嘿嘿喵3號',0,'1234567890',2979000,30,0,0,0,0,0,0,0,0,1,0),(4,1,'192.168.1.2','11:22:33:44:55','2017-07-21 14:28:39','2017-09-08 17:32:54','2017-09-08 17:34:13','cat444','1234','嘿嘿喵4號',0,'1234567890',4000000,40,0,0,0,0,0,0,0,0,1,0),(5,1,'192.168.1.2','11:22:33:44:55','2017-07-21 14:28:39','2017-09-08 17:33:03','2017-09-08 17:34:24','cat555','1234','嘿嘿喵5號',0,'1234567890',5000000,50,0,0,0,0,0,0,0,0,1,0),(6,1,'192.168.1.2','11:22:33:44:55','2017-07-21 14:28:39','2017-09-08 17:29:54','2017-09-08 17:34:32','cat666','1234','嘿嘿喵6號',0,'1234567890',6000000,60,0,0,0,0,0,0,0,0,1,0);
/*!40000 ALTER TABLE `memberinfo` ENABLE KEYS */;
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

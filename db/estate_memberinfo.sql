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
-- Table structure for table `memberinfo`
--

DROP TABLE IF EXISTS `memberinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `memberinfo` (
  `User_ID` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '員工帳號編號\nSerial Number 一個流水號, 當新增一個帳號自動加1, 帳號唯一號碼',
  `PlatformID` int(11) NOT NULL COMMENT '第三方平台編號',
  `DeviceID` int(11) NOT NULL COMMENT '員工裝置編號',
  `IP` varchar(40) NOT NULL COMMENT '員工IP Address',
  `MacAddress` varchar(50) DEFAULT NULL COMMENT '員工網卡 Address',
  `CreateTime` datetime NOT NULL COMMENT '帳號建立時間',
  `LoginTime` datetime NOT NULL COMMENT '帳號登入時間',
  `UpdateTime` datetime NOT NULL COMMENT '帳號更新時間',
  `Account` varchar(45) NOT NULL COMMENT '員工帳號',
  `Password` varchar(45) NOT NULL COMMENT '員工密碼',
  `NickName` varchar(45) NOT NULL COMMENT '員工暱稱',
  `IdentityNumber` char(11) NOT NULL COMMENT '員工身分證字號',
  `Address` varchar(45) NOT NULL COMMENT '員工地址',
  `PhoneNumber` varchar(10) NOT NULL COMMENT '員工電話號碼',
  `Balance` int(11) NOT NULL COMMENT '員工的錢\n單位 人民幣/浮點數  ( 從平台帶過來的錢)',
  `Bonus` int(11) NOT NULL COMMENT '員工紅利',
  `Salary` int(11) NOT NULL COMMENT '員工薪水',
  `Status` tinyint(3) NOT NULL COMMENT '員工狀態  0: 登出 1:登入 2:斷線中 3:斷線連回中',
  `Vip_rank` tinyint(3) NOT NULL COMMENT '員工位階\n0:實習生 1:一般員工 2:組長 3:主任 4:經理 5:人資 6:會計 7:最高管理者',
  PRIMARY KEY (`User_ID`,`Account`),
  UNIQUE KEY `User_ID_UNIQUE` (`User_ID`),
  UNIQUE KEY `Account_UNIQUE` (`Account`),
  UNIQUE KEY `IdentityNumber_UNIQUE` (`IdentityNumber`),
  KEY `PlatformID` (`PlatformID`),
  KEY `Account` (`Account`),
  KEY `Password` (`Password`),
  KEY `User_ID` (`User_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 COMMENT='紀錄每個會員基本資料';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `memberinfo`
--

LOCK TABLES `memberinfo` WRITE;
/*!40000 ALTER TABLE `memberinfo` DISABLE KEYS */;
INSERT INTO `memberinfo` VALUES (1,0,0,'','','2017-06-23 01:00:53','2017-11-27 00:44:59','2017-11-27 00:44:59','cat111','1234','cat1111','F123456789','地球','0955645161',0,0,200,1,6),(39,0,0,'','','2017-10-04 17:26:46','2017-10-04 17:26:46','2017-10-04 17:26:46','dog111','1234','dog111','dog111','tw','111',0,0,3000,0,3),(45,0,0,'','','2017-10-12 00:55:37','2017-10-12 00:55:37','2017-10-23 22:10:19','eagle111','1234','eagle111','F124000001','tw','0966123456',0,0,999999,0,0);
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

-- Dump completed on 2017-11-27  0:49:06

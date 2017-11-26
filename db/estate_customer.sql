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
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `CustomerID` int(11) NOT NULL AUTO_INCREMENT COMMENT '顧客ID',
  `User_ID` bigint(11) NOT NULL COMMENT '會員編號',
  `NickName` varchar(45) NOT NULL COMMENT '會員名稱',
  `CreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '建立日期',
  `UpdateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
  `CustomerName` varchar(45) NOT NULL COMMENT '顧客名稱',
  `CustomerAge` int(11) DEFAULT NULL COMMENT '顧客年紀',
  `CustomerGender` varchar(2) DEFAULT NULL COMMENT '顧客性別',
  `CustomerIdentityNumber` varchar(11) NOT NULL COMMENT '顧客身份證字號',
  `CustomerPhoneNumber` varchar(11) NOT NULL COMMENT '顧客手機號碼',
  `CustomerAddress` varchar(60) DEFAULT NULL COMMENT '顧客地址',
  `CustomerHomeID` int(11) DEFAULT NULL COMMENT '顧客所擁有的房屋物件ID',
  `CustomerHomeAge` int(11) DEFAULT NULL COMMENT '房屋屋齡',
  `CustomerHomeFootage` float DEFAULT NULL COMMENT '房屋坪數',
  `CustomerHomePrice` int(11) DEFAULT NULL COMMENT '房屋價格',
  `Vip_rank` tinyint(3) DEFAULT NULL COMMENT '顧客等級 0:一般 1:熟客 2:VIP',
  PRIMARY KEY (`CustomerID`),
  UNIQUE KEY `CustomerIdentityNumber_UNIQUE` (`CustomerIdentityNumber`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (6,1,'cat1111','2017-10-31 16:54:40','2017-10-31 16:54:40','測試顧客',0,'F','','','',0,0,0,0,2),(13,1,'cat1111','2017-11-02 14:57:42','2017-11-02 14:57:42','cat',18,'F','f00001','0988888','tw',1,20,0,200,2),(14,1,'cat1111','2017-11-02 15:03:42','2017-11-02 15:03:42','dog1',20,'F','f002','0955666123','jp',1,30,30,30,1);
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
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

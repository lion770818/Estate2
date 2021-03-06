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
-- Table structure for table `gamelog`
--

DROP TABLE IF EXISTS `gamelog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gamelog` (
  `idx` int(11) NOT NULL AUTO_INCREMENT,
  `CreateTime` datetime DEFAULT NULL,
  `Platform` int(1) NOT NULL,
  `UID` int(11) DEFAULT NULL,
  `Name` varchar(45) NOT NULL,
  `Account` varchar(32) NOT NULL,
  `TotalBet` int(11) DEFAULT NULL,
  `Bet` int(11) DEFAULT NULL,
  `Before_Game_Money` int(11) DEFAULT NULL,
  `After_Game_Money` int(11) DEFAULT NULL,
  PRIMARY KEY (`idx`,`Name`,`Platform`,`Account`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gamelog`
--

LOCK TABLES `gamelog` WRITE;
/*!40000 ALTER TABLE `gamelog` DISABLE KEYS */;
INSERT INTO `gamelog` VALUES (1,'2017-06-23 17:24:07',1,1,'劉智明1','cat111',10,10,499900,499890),(2,'2017-06-23 17:24:25',1,1,'劉智明1','cat111',10,10,499890,499880),(3,'2017-06-23 17:25:22',1,1,'劉智明1','cat111',10,10,499880,499870),(4,'2017-06-23 17:27:25',1,1,'劉智明1','cat111',10,10,499870,499860),(5,'2017-06-23 17:27:28',1,1,'劉智明1','cat111',10,10,499860,499850),(6,'2017-06-23 17:45:33',1,1,'劉智明1','cat111',10,10,499850,499840),(7,'2017-06-23 17:51:08',1,1,'劉智明1','cat111',10,10,499840,499830),(8,'2017-06-23 17:51:12',1,1,'劉智明1','cat111',10,10,499830,499820),(9,'2017-06-23 17:51:13',1,1,'劉智明1','cat111',10,10,499820,499810),(10,'2017-06-23 19:37:01',1,1,'劉智明1','cat111',10,10,499810,499800),(11,'2017-06-23 19:37:07',1,1,'劉智明1','cat111',10,10,499800,499790),(12,'2017-06-23 19:37:09',1,1,'劉智明1','cat111',10,10,499790,499780),(13,'2017-06-23 19:37:10',1,1,'劉智明1','cat111',10,10,499780,499770);
/*!40000 ALTER TABLE `gamelog` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-18 22:38:34

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
-- Table structure for table `lobbyinfo`
--

DROP TABLE IF EXISTS `lobbyinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lobbyinfo` (
  `PlatformID` int(11) NOT NULL COMMENT '第三方平台編號',
  `LobbyID` int(11) NOT NULL AUTO_INCREMENT COMMENT '大廳規則編號',
  `GameID` int(11) NOT NULL COMMENT '遊戲編號',
  `LobbyName` varchar(45) NOT NULL COMMENT '大聽中文名稱',
  `LobbyMatchID` tinyint(4) NOT NULL COMMENT '大廳配桌編號, 編號相同的廳館, 才可以配桌在一起\n不知道怎樣設定就先跟 PlatformID 一樣\n等想跨平台配在一起  在調整\nif( LobbyMatchID 相同  && GameID 相同 && 人數未滿 )\n     加入桌()',
  `total_water1` double NOT NULL COMMENT '總水池1',
  `total_water2` double NOT NULL COMMENT '總水池2',
  `BetLevel` double NOT NULL COMMENT '單一押注金額',
  PRIMARY KEY (`LobbyID`),
  KEY `LobbyID` (`LobbyID`),
  KEY `PlatformID` (`PlatformID`),
  KEY `GameID` (`GameID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='紀錄每個區廳的下注 倍率';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lobbyinfo`
--

LOCK TABLES `lobbyinfo` WRITE;
/*!40000 ALTER TABLE `lobbyinfo` DISABLE KEYS */;
INSERT INTO `lobbyinfo` VALUES (1,1,1001,'捕魚機一般館',1,10,20,10),(1,2,1002,'捕魚機中級館',1,20,30,20.5),(1,3,1003,'捕魚機高級館',1,50,60,30.55555),(1,4,201,'德州撲克',1,5,10,10),(2,5,1001,'捕魚機一般館',2,10,20,10),(2,6,1002,'捕魚機中級館',2,20,30,20),(2,7,1003,'捕魚機高級館',2,50,60,40),(1,8,2001,'huga一般館',100,100,100,20),(1,9,2001,'huga一般館',100,150,150,40),(1,10,2001,'huga一般館',100,200,200,60);
/*!40000 ALTER TABLE `lobbyinfo` ENABLE KEYS */;
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

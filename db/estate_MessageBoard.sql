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
-- Table structure for table `MessageBoard`
--

DROP TABLE IF EXISTS `MessageBoard`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MessageBoard` (
  `MessageBoardID` int(11) NOT NULL AUTO_INCREMENT COMMENT '留言板ID',
  `MessageBoardName` varchar(45) NOT NULL,
  `User_ID` bigint(11) NOT NULL COMMENT '員工編號(哪個員工新增的)\n',
  `NickName` varchar(45) NOT NULL COMMENT '員工暱稱',
  `HomeName` varchar(45) NOT NULL COMMENT '房屋名稱',
  `CreateTime` datetime NOT NULL COMMENT '留言板新增時間',
  `UpdateTime` datetime NOT NULL COMMENT '留言板更新時間',
  `PhoneNumber` varchar(10) NOT NULL COMMENT '員工電話號碼',
  `Gender` varchar(2) NOT NULL COMMENT '性別',
  `Rent` int(11) NOT NULL COMMENT '租金',
  `IsPet` int(11) NOT NULL COMMENT '是否養寵物',
  `IsSmoke` int(11) NOT NULL COMMENT '是否抽煙',
  `IsHouseholdRegister` int(11) NOT NULL COMMENT '是否入戶籍',
  `IsTax` int(11) NOT NULL COMMENT '是否報稅',
  `Isfaith` int(11) NOT NULL COMMENT '是否信仰',
  `Memo` varchar(45) NOT NULL COMMENT '備忘錄',
  PRIMARY KEY (`MessageBoardID`),
  UNIQUE KEY `MessageBoardID_UNIQUE` (`MessageBoardID`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MessageBoard`
--

LOCK TABLES `MessageBoard` WRITE;
/*!40000 ALTER TABLE `MessageBoard` DISABLE KEYS */;
INSERT INTO `MessageBoard` VALUES (14,'測試上傳',1,'cat1112t','','2017-12-04 08:25:56','2017-12-04 08:25:56','0955645161','',3,1,0,1,0,0,'我想住\n急需'),(15,'',1,'cat1112t','','2017-12-04 08:32:42','2017-12-04 08:32:42','','',1,0,0,0,0,0,'內容'),(16,'',1,'cat1112t','','2017-12-04 08:35:30','2017-12-04 08:35:30','','',1,0,0,0,0,0,'內容'),(17,'cat111',1,'cat1112t','','2017-12-04 08:37:51','2017-12-04 08:37:51','1234','',2,0,1,0,1,0,'內容cccc'),(18,'cat111',1,'cat1112t','','2017-12-04 08:37:57','2017-12-04 08:37:57','1234','',2,0,1,0,1,0,'內容cccc'),(19,'cat111',1,'cat1112t','','2017-12-04 08:37:57','2017-12-04 08:37:57','1234','',2,0,1,0,1,0,'內容cccc'),(20,'cat222',1,'cat1112t','','2017-12-05 08:30:54','2017-12-05 08:30:54','09888123','',1,0,0,0,0,0,'內容'),(21,'測試新增3',1,'cat1112t','','2017-12-05 08:40:22','2017-12-05 08:40:22','09123','',3,1,0,0,1,1,'哈哈哈哈哈'),(22,'測試上傳3',1,'cat1112t','','2017-12-05 08:51:28','2017-12-05 08:51:28','09000001','',3,0,1,0,1,0,'j嗚嗚嗚嗚嗚嗚'),(23,'測試上傳4',1,'cat1112t','','2017-12-05 08:53:47','2017-12-05 08:53:47','111','',4,1,1,1,1,1,'吧吧吧吧ㄅ'),(24,'測試上傳5',1,'cat1112t','','2017-12-05 09:00:25','2017-12-05 09:00:25','cccc','',3,0,0,0,1,0,'111'),(25,'新的留言板',1,'嘿嘿喵','','2017-12-05 12:58:40','2017-12-05 12:58:40','0912345678','F',0,0,0,0,0,0,''),(26,'新的留言板',1,'嘿嘿喵','','2017-12-05 13:13:19','2017-12-05 13:13:19','0912345678','F',0,0,0,0,0,0,''),(27,'新的留言板',1,'嘿嘿喵','','2017-12-06 01:30:16','2017-12-06 01:30:16','0912345678','F',0,0,0,0,0,0,''),(28,'del111',1,'cat1112t','','2017-12-06 08:17:30','2017-12-06 08:17:30','1111','',4,0,0,0,1,1,'sssss'),(29,'新的留言板',1,'嘿嘿喵','','2017-12-08 00:57:09','2017-12-08 00:57:09','0912345678','F',0,0,0,0,0,0,''),(30,'新的留言板',1,'嘿嘿喵','','2017-12-11 23:12:29','2017-12-11 23:12:29','0912345678','F',0,0,0,0,0,0,'');
/*!40000 ALTER TABLE `MessageBoard` ENABLE KEYS */;
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

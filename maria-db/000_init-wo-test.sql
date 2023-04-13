



--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `bird`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bird` (
  `idx` int(11) NOT NULL,
  `name`     varchar(100) NOT NULL,
  `description` varchar(1024) NOT NULL,
  PRIMARY KEY (`idx`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;



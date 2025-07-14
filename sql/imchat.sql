-- MySQL dump 10.13  Distrib 8.0.42, for Linux (x86_64)
--
-- Host: 117.72.47.192    Database: imchat
-- ------------------------------------------------------
-- Server version	8.0.42-0ubuntu0.22.04.1

--
-- Table structure for table `im_users`
--

DROP TABLE IF EXISTS `im_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `im_users` (
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `gender` int DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `im_online_status` int DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_frozen` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_im_users_user_name` (`user_name`),
  KEY `idx_im_users_user_code` (`user_code`),
  KEY `idx_im_users_phone` (`phone`),
  KEY `idx_im_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `im_users`
--

LOCK TABLES `im_users` WRITE;
/*!40000 ALTER TABLE `im_users` DISABLE KEYS */;
INSERT INTO `im_users` VALUES ('17518934737600000001','33033761','ZhangSan','$2a$10$Xa55vLFhA76J20RoGLDUa.c/5vfAUDXNm2fZV3UVO5zXW0jb2BH2O','15070241126','https://avatars.githubusercontent.com/u/95828444',1,'hourin.dev@gmail.com',0,'2025-07-07 21:04:34','2025-07-07 21:04:34',1),('17518979022830000003','5846810130','aa111111111a111','$2a$10$7xP/LaTCWUo8c6mZui5dQ.A1CEC93uKfF01c5uzVMdhLJwAQ0u6La','','',0,'',0,'2025-07-07 22:18:22','2025-07-07 22:18:22',1),('17520735848490000004','3425234016','zhangzhangsan','$2a$10$zdaxOGFWms8YQ1LcETMAL.QRyo1EJO5Q0kUbX4TgEAtlwsD/1CjlS','03328379575','https://avatars.githubusercontent.com/u/95828444',8,'uclgm3.ngt45@gmail.com',0,'2025-07-09 23:06:25','2025-07-09 23:06:25',1),('17520762603740000006','934221828','zhangzhangsan','$2a$10$UoowdJti8Uhp7T94huUcge9JwYxSwBcDAuHsll3/Mvgj1fkOP1EM6','03328379575','https://avatars.githubusercontent.com/u/95828444',8,'uclgm3.ngt45@gmail.com',0,'2025-07-09 23:51:00','2025-07-09 23:51:00',1);
/*!40000 ALTER TABLE `im_users` ENABLE KEYS */;
UNLOCK TABLES;
-- Dump completed on 2025-07-10 20:53:59

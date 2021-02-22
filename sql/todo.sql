
SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `todo`;
CREATE TABLE `todo` (
  `sno` int(11) NOT NULL AUTO_INCREMENT,
  `id` varchar(32) NOT NULL,
  `todoID` varchar(32) NOT NULL,
  `title` varchar(64) NOT NULL,
  `description` text NOT NULL,
  `todopriority` enum('HIGH','MEDIUM','LOW') NOT NULL,
  `todotype` enum('TASK') NOT NULL,
  `enddate` datetime NOT NULL,
  `iscompleted` enum('true','false') NOT NULL DEFAULT 'false',
  PRIMARY KEY (`sno`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


SET NAMES utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `sno` int(11) NOT NULL AUTO_INCREMENT,
  `id` varchar(32) NOT NULL,
  `firstname` varchar(32) NOT NULL,
  `lastname` varchar(32) NOT NULL,
  `email` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  PRIMARY KEY (`sno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



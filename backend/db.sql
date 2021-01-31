create database if not exists `blog`;

use `blog`;

CREATE TABLE `blog_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL,
  `password` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `introduction` varchar(255) DEFAULT NULL,
  `nickname` varchar(32) DEFAULT NULL,
  `about` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

 
INSERT INTO blog_user
(`username`, `password`, `avatar`, `introduction`, `nickname`, `about`) VALUES
('admin', 'e10adc3949ba59abbe56e057f20f883e', 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2884107401,3797902000&fm=26&gp=0.jpg', 'a gopher', 'gopher', '## hello world');


CREATE TABLE `blog_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_name` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(16) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(32) NOT NULL,
  `content` text NOT NULL,
  `html` text NOT NULL,
  `category_id` int(10) unsigned NOT NULL,
  `created_time` varchar(32) NOT NULL,
  `updated_time` varchar(32) DEFAULT '',
  `status` varchar(16) NOT NULL DEFAULT 'published',
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`) USING BTREE,
  KEY `category` (`category_id`),
  CONSTRAINT `category` FOREIGN KEY (`category_id`) REFERENCES `blog_category` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `blog_tag_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned NOT NULL,
  `article_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tag_id` (`tag_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `article_id` FOREIGN KEY (`article_id`) REFERENCES `blog_article` (`id`),
  CONSTRAINT `tag_id` FOREIGN KEY (`tag_id`) REFERENCES `blog_tag` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `blog_soup` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` tinytext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO blog_soup (`content`) VALUES ('你全心全力做到的最好\r\n可能还不如别人的随便搞搞');

CREATE TABLE `blog_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL,
  `is_author` tinyint(1) NOT NULL DEFAULT '0',
  `parent_id` int(10) unsigned DEFAULT NULL,
  `root_id` int(10) unsigned DEFAULT NULL,
  `article_id` int(10) unsigned NOT NULL,
  `content` varchar(255) NOT NULL,
  `created_time` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article` (`article_id`),
  KEY `parent` (`parent_id`),
  KEY `root` (`root_id`),
  CONSTRAINT `root` FOREIGN KEY (`root_id`) REFERENCES `blog_comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `article` FOREIGN KEY (`article_id`) REFERENCES `blog_article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `parent` FOREIGN KEY (`parent_id`) REFERENCES `blog_comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `blog_visitor` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `uri` varchar(255) NOT NULL COMMENT '访问路径',
    `referer` varchar(255) DEFAULT '',
    `ua` varchar(255) NOT NULL COMMENT 'user_agent',
    `ip` varchar(64) NOT NULL,
    `visit_date` varchar(64) NOT NULL,
    `visit_time` varchar(255) NOT NULL COMMENT '访问时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
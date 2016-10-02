CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `content` text COLLATE utf8_unicode_ci,
  `created_at` datetime DEFAULT NOW(),
  `updated_at` datetime DEFAULT NOW(),
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_articles_on_title` (`title`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `articles` (
  `id`,
  `title`,
  `content`
)
VALUES
(1,  'hello1',  'world'),
(2,  'hello2',  'world'),
(3,  'hello3',  'world'),
(4,  'hello4',  'world'),
(5,  'hello5',  'world'),
(6,  'hello6',  'world'),
(7,  'hello7',  'world'),
(8,  'hello8',  'world'),
(9,  'hello9',  'world'),
(10, 'hello10', 'world');

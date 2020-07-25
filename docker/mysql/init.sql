DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
  `user_password` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT 'パスワード',
  `user_name` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'ユーザー名',
  `user_email` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Email',
  `user_introduce` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '自己紹介文',
  `delete_flag` tinyint(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ 0:未削除　1:削除済み',
  `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `posted_photos`;
CREATE TABLE `posted_photos` (
  `photo_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '写真ID',
  `user_id` int(11) NOT NULL COMMENT 'ユーザーID',
  `photo_url` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT '写真URL',
  `photo_category` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'カテゴリ',
  `photo_comment` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'コメント',
  `delete_flag` tinyint(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ 0:未削除　1:削除済み',
  `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  PRIMARY KEY (`photo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `users` (`user_id`, `user_password`, `user_name`,`user_email`,`user_introduce`) VALUES
(1, 'foo', 'foo','foo@example.com', 'fooです！写真の腕には自信あります！'),
(2, 'bar', 'bar','bar@example.com', 'barと申します。よろしくお願いいたします。'),
(3, 'hoge', 'hoge','foo@example.com', 'hogeです。よろ');
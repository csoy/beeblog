PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `category` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `title` varchar(255) NOT NULL DEFAULT '' ,
    `created` datetime NOT NULL,
    `topic_time` datetime NOT NULL,
    `topic_count` integer NOT NULL DEFAULT 0 ,
    `topic_last_user_id` integer NOT NULL DEFAULT 0 
);
INSERT INTO "category" VALUES(1,'math',20161207,151211,23,12);
INSERT INTO "category" VALUES(2,'linux','2016-12-07 09:17:57.8092888+00:00','2016-12-07 09:17:57.8092888+00:00',0,0);
INSERT INTO "category" VALUES(3,'gggg','2016-12-08 02:10:46.3032499+00:00','2016-12-08 02:10:46.3032499+00:00',0,0);
CREATE TABLE `topic` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `uid` integer NOT NULL DEFAULT 0 ,
    `title` varchar(255) NOT NULL DEFAULT '' ,
    `content` varchar(255) NOT NULL DEFAULT '' ,
    `attachment` varchar(255) NOT NULL DEFAULT '' ,
    `created` datetime NOT NULL,
    `updated` datetime NOT NULL,
    `views` integer NOT NULL DEFAULT 0 ,
    `author` varchar(255) NOT NULL DEFAULT '' ,
    `reply_time` datetime NOT NULL,
    `reply_count` integer NOT NULL DEFAULT 0 ,
    `reply_last_user_id` integer NOT NULL DEFAULT 0 
);
INSERT INTO "topic" VALUES(2,0,'测试2','22222222222222222222222222222222222','','2016-12-07 07:00:45.7364412+00:00','2016-12-07 07:00:45.7364412+00:00',11,'','2016-12-07 07:00:45.7364412+00:00',0,0);
INSERT INTO "topic" VALUES(3,0,'测试3','333333333333333333333333333333333333333','','2016-12-07 07:05:22.5252726+00:00','2016-12-07 07:05:22.5252726+00:00',0,'','2016-12-07 07:05:22.5252726+00:00',0,0);
INSERT INTO "topic" VALUES(5,0,'linux','yum源制作','','2016-12-07 09:18:27.7530015+00:00','2016-12-07 09:18:27.7530015+00:00',1,'','2016-12-07 09:18:27.7530015+00:00',0,0);
DELETE FROM sqlite_sequence;
INSERT INTO "sqlite_sequence" VALUES('topic',5);
INSERT INTO "sqlite_sequence" VALUES('category',3);
CREATE INDEX `category_created` ON `category` (`created`);
CREATE INDEX `category_topic_time` ON `category` (`topic_time`);
CREATE INDEX `topic_created` ON `topic` (`created`);
CREATE INDEX `topic_updated` ON `topic` (`updated`);
CREATE INDEX `topic_views` ON `topic` (`views`);
CREATE INDEX `topic_reply_time` ON `topic` (`reply_time`);
COMMIT;

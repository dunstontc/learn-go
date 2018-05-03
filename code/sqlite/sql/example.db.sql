BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `User` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`name`	TEXT NOT NULL,
	`email`	TEXT NOT NULL UNIQUE,
	`password`	TEXT NOT NULL,
	`hash`	TEXT
);
INSERT INTO `User` (id,name,email,password,hash) VALUES (1,'James Bond','007@gmail.com','pussygalore',NULL);
INSERT INTO `User` (id,name,email,password,hash) VALUES (2,'Clay Dunston','me@clay.cool','password',NULL);
CREATE TABLE IF NOT EXISTS `Todo` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`user_id`	INTEGER NOT NULL,
	`complete`	INTEGER NOT NULL,
	`priority`	TEXT,
	`created`	NUMERIC NOT NULL,
	`title`	TEXT NOT NULL,
	`content`	TEXT,
	FOREIGN KEY(`user_id`) REFERENCES `User`(`id`)
);
COMMIT;

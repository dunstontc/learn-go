-- CREATE TABLE "Todo" ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `user_id` INTEGER NOT NULL, `complete` INTEGER NOT NULL, `priority` TEXT, `created` NUMERIC NOT NULL, `title` TEXT NOT NULL, `content` TEXT, FOREIGN KEY(`user_id`) REFERENCES `User`(`id`) )
CREATE TABLE `Todo` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`priority`	TEXT,
	`title`	TEXT NOT NULL,
	`content`	TEXT,
	`created`	NUMERIC NOT NULL,
	`user_id`	INTEGER,
	FOREIGN KEY(`user_id`) REFERENCES `User`(`id`)
);

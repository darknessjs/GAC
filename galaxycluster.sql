CREATE TABLE IF NOT EXISTS `sc2_user`(
   `id` INT UNSIGNED,
   `realm` INT NOT NULL,
   `name` VARCHAR(40) NOT NULL,
   `displayName` VARCHAR(40) NOT NULL,
   `clanName` VARCHAR(40) NOT NULL,
   `clanTag` VARCHAR(40) NOT NULL,
   `profilePath` VARCHAR(40) NOT NULL,
   PRIMARY KEY ( `id` )
) DEFAULT CHARSET=utf8;
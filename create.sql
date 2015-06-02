create table `user` 
    -- --------------------------------------------------
    --  Table Structure for `firstapi/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `user` (
        `id` varchar(255) NOT NULL PRIMARY KEY,
        `username` varchar(20) NOT NULL,
        `password` varchar(100) NOT NULL,
        `gender` varchar(2) NOT NULL,
        `age` integer NOT NULL,
        `address` varchar(100) NOT NULL,
        `email` varchar(30) NOT NULL,
        `created_time` datetime NOT NULL,
        `weight` integer NOT NULL
    ) ENGINE=InnoDB;
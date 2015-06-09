
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
        `created_time` datetime NOT NULL ,
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `weight` integer NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;


  

    -- --------------------------------------------------
    --  Table Structure for `firstapi/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `blogcategory` (
        `id` bigint NOT NULL PRIMARY KEY auto_increment,
        `user_id` varchar(255) NOT NULL,
        `title` varchar(50) NOT NULL,
        `descri` varchar(255) NOT NULL,
        `created_time` datetime NOT NULL,
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        foreign key (user_id) references user(id) on delete cascade on update cascade
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;




    -- --------------------------------------------------
    --  Table Structure for `firstapi/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `blog` (
        `blogid` bigint NOT NULL PRIMARY KEY auto_increment,
        `blog_category_id` bigint NOT NULL,
        `user_id` varchar(255) NOT NULL,
        `blog_title` varchar(50) NOT NULL,
        `content` varchar(9000) NOT NULL,
        `imageurl` varchar(100) NOT NULL,
        `tags` varchar(30) NOT NULL,
        `created_time` datetime NOT NULL  ,
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `public` integer NOT NULL,
        foreign key (blog_category_id) references blogcategory(id) on delete cascade on update cascade,
        foreign key (user_id) references user(id) on delete cascade on update cascade
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
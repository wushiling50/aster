create table `aster`.`user`
(
    `id`                bigint              not null comment 'ID',
    `name`              varchar(255)        not null comment 'name',
    `score`             decimal             not null comment 'score',
    `nation`            varchar(255)        not null comment 'nation',
    `confidence_level`  varchar(255)        not null comment 'level',      
    constraint `id`
        primary key (`id`)
)engine=InnoDB default charset=utf8mb4;

create table `aster`.`domain`
(
    `id`                bigint              not null comment 'ID',
    `user_id`           bigint              not null comment 'UserID',
    `tag`               varchar(255)        not null comment 'tag',  
    constraint `id`
        primary key (`id`),
    constraint `user_domain`
        foreign key (`user_id`)
            references `aster`.`user` (`id`)
            on delete cascade
)engine=InnoDB default charset=utf8mb4;
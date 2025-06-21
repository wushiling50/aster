create table `aster`.`languages` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `language`                   json            not null       comment 'Programming Languages Used (JSON Format)',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`nation` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `nation`                     varchar(255)    not null       comment 'Nation',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`score` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `contribution_id`            bigint          not null       comment 'Unique Contribution ID',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `score`                      double          not null       comment 'score',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`summary` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `summary`                    text            not null       comment 'User Summary',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;
create table `aster`.`languages` (
    `data_id`                    bigint          not null unique      comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null unique      comment 'Unique GitHub User ID',
    `language`                   json            not null       comment 'Programming Languages Used (JSON Format)',
    `created_at`            timestamp       not null default  current_timestamp,
    `updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`developer_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`nation` (
    `data_id`                    bigint          not null  unique     comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null  unique     comment 'Unique GitHub User ID',
    `nation`                     varchar(255)    not null       comment 'Nation',
    `confidence`                 double          not null       comment 'Confidence',
    `created_at`            timestamp       not null default  current_timestamp,
    `updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`developer_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`score` (
    `data_id`                    bigint          not null  unique     comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null  unique     comment 'Unique GitHub User ID',
    `score`                      double          not null       comment 'score',
    `created_at`            timestamp       not null default  current_timestamp,
    `updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`developer_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`summary` (
    `data_id`                    bigint          not null  unique     comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null  unique     comment 'Unique GitHub User ID',
    `summary`                    text            not null       comment 'User Summary',
    `created_at`            timestamp       not null default  current_timestamp,
    `updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`developer_id`)
) engine=InnoDB default charset=utf8mb4;
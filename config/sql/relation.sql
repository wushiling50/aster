create table `aster`.`create_repo` (
    `data_id`                    bigint          not null unique      comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null            comment 'Unique GitHub User ID',
    `repo_id`                    bigint          not null unique     comment 'Unique GitHub Repository ID',
    `created_at`                 timestamp       not null default  current_timestamp,
    `updated_at`                 timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `deleted_at`                 timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`created_repo_updated_at` (
    `data_id`                    bigint          not null  unique     comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null  unique     comment 'Unique GitHub User ID',
    `created_at`                 timestamp       not null default  current_timestamp,
    `updated_at`                 timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`developer_id`)
) engine=InnoDB default charset=utf8mb4;
create table `aster`.`contribution` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `contribution_id`            bigint          not null       comment 'Unique Contribution ID',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `repo_id`                    bigint          not null       comment 'Unique GitHub Repository ID',
    `category`                   varchar(20)     not null       comment 'Category',
    `content`                    text            not null       comment 'Contribution Content',
    `created_at`                 timestamp       not null       comment 'Original creation timestamp',
    `updated_at`                 timestamp       not null       comment 'Last update timestamp',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`issue_pr_of_user_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`comment_of_user_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`review_of_user_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;
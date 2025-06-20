create table `aster`.`developer` (
    `data_id`               bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `id`                    bigint          not null       comment 'Unique GitHub User ID',
    `name`                  varchar(255)    not null       comment 'Name',
    `login`                 varchar(255)    not null       comment 'Login',
    `avatar_url`            varchar(255)    not null       comment 'Profile Avatar URL',
    `company`               varchar(255)    not null       comment 'Company',
    `location`              varchar(255)    not null       comment 'Location',
    `bio`                   varchar(255)    not null       comment 'User Biography',
    `blog`                  varchar(255)    not null       comment 'Personal Blog URL',
    `email`                 varchar(255)    not null       comment 'Contact Email',
    `twitter_username`      varchar(255)    not null       comment 'Twitter Username',
    `repos`                 bigint          not null       comment 'Repository Count',
    `following`             bigint          not null       comment 'Following Count',
    `followers`             bigint          not null       comment 'Followers Count',
    `stars`                 bigint          not null       comment 'Total Stars Received',
    `gists`                 bigint          not null       comment 'Gists count',
    `created_at`            timestamp       not null       comment 'GitHub account creation time',
    `updated_at`            timestamp       not null       comment 'GitHub account update time',
    `data_created_at`       timestamp       not null default  current_timestamp,
    `data_updated_at`       timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`       timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_developer_id` (`id`),
    index `idx_login` (`login`)
) engine=InnoDB default charset=utf8mb4;

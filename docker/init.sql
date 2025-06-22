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

create table `aster`.`repo` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `id`                         bigint          not null       comment 'Unique GitHub Repository ID',
    `name`                       varchar(255)    not null       comment 'Repository Name',
    `star_count`                 bigint          not null       comment 'Number Of Stars',
    `fork_count`                 bigint          not null       comment 'Number Of Forks',
    `issue_count`                bigint          not null       comment 'Number Of Open Issues',
    `commit_count`               bigint          not null       comment 'Number Of Commits',
    `pr_count`                   bigint          not null       comment 'Total Pull Requests Count',
    `language`                   json            not null       comment 'Programming Languages Used (JSON Format)',
    `description`                varchar(255)    not null       comment 'Repository Description',
    `last_fetch_fork_at`         bigint          not null       comment 'Unix Timestamp Of Last Fork Data Fetch',
    `last_fetch_contribution_at` bigint          not null       comment 'Unix Timestamp Of Last Contribution Data Fetch',
    `merged_pr_count`            bigint          not null       comment 'Count Of Merged Pull Requests',
    `open_pr_count`              bigint          not null       comment 'Count Of Open Pull Requests',
    `comment_count`              bigint          not null       comment 'Total Comments Count',
    `review_count`               bigint          not null       comment 'Code Reviews Count',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`),
    index `idx_repo_id` (`id`)
) engine=InnoDB default charset=utf8mb4;

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

create table `aster`.`create_repo` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `repo_id`                    bigint          not null       comment 'Unique GitHub Repository ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`follow` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `follower_id`                bigint          not null       comment 'Follower ID',
    `following_id`               bigint          not null       comment 'Following ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`fork` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `original_repo_id`           bigint          not null       comment 'Original Repository ID',
    `fork_repo_id`               bigint          not null       comment 'Fork Repository ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    `data_deleted_at`            timestamp       null     default null,
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`star` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `repo_id`                    bigint          not null       comment 'Unique GitHub Repository ID',
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

create table `aster`.`created_repo_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`following_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`follower_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`fork_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;

create table `aster`.`starred_repo_updated_at` (
    `data_id`                    bigint          not null       comment 'Generated Primary Key, Must Not Be Changed',
    `developer_id`               bigint          not null       comment 'Unique GitHub User ID',
    `data_created_at`            timestamp       not null default  current_timestamp,
    `data_updated_at`            timestamp       not null default  current_timestamp on update current_timestamp comment 'update data time',
    constraint `pk_data_id`
        primary key (`data_id`)
) engine=InnoDB default charset=utf8mb4;
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

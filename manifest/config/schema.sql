create table if not exists rss_feed_channel (
    id varchar(64) not null primary key,
    title varchar(128) null,
    channel_desc mediumtext null,
    image_url varchar(128) null,
    link varchar(128) null,
    rss_link varchar(128) null
);
-- create table if not exists rss_feed_item (
--     id varchar(64) not null primary key,
--     channel_id varchar(64) not null,
--     title mediumtext null,
--     content mediumtext null,
--     link varchar(128) null,
--     date date null,
--     author varchar(128) null,
--     input_date datetime null,
--     thumbnail varchar(128) null,
--     description mediumtext null
-- );
-- create index rfi_idx_channel_id on rss_feed_item (channel_id);

-- MySQL fulltext search table
create table if not exists rss_feed_item (
    id varchar(64) not null primary key,
    channel_id varchar(64) not null,
    title mediumtext null,
    content mediumtext null,
    link varchar(128) null,
    date date null,
    author varchar(128) null,
    input_date datetime null,
    thumbnail varchar(256) null,
    description mediumtext null,
    FULLTEXT KEY rss_feed_item_fts_key (title, content, author, description) WITH PARSER ngram
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
create index rfi_idx_channel_id on rss_feed_item (channel_id);
-- create table if not exists rss_feed_tag
-- (
--     name       varchar(64) not null,
--     channel_id varchar(64) not null,
--     title      mediumtext  null,
--     date       date        null
-- );
-- create index rss_feed_tag_idx_name
--     on rss_feed_tag (name);
create table if not exists user_info (
    uid varchar(32) not null primary key,
    password varchar(32) not null,
    email varchar(32) null,
    mobile varchar(32) null,
    username varchar(32) null,
    create_date datetime null,
    update_date datetime null
);
create table if not exists user_mark_feed_item (
    user_id varchar(32) null,
    channel_item_id varchar(64) null,
    input_time datetime null,
    status int default 1 null,
    constraint umfi_uid_itemid_pk unique (user_id, channel_item_id)
);
create index umfi_uid_index on user_mark_feed_item (user_id);
create table if not exists user_sub_channel (
    user_id varchar(32) null,
    channel_id varchar(64) null,
    input_time datetime null,
    status int default 1 null,
    constraint usc_uid_cid_pk unique (user_id, channel_id)
);
create index usc_idx_uid on user_sub_channel (user_id);

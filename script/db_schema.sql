create table if not exists rss_feed_channel
(
    id varchar(64) not null
        primary key,
    title varchar(128) null,
    channel_desc mediumtext null,
    image_url varchar(128) null,
    link varchar(128) null,
    rsshub_link varchar(128) null
);

create table if not exists rss_feed_item
(
    id varchar(64) not null
        primary key,
    channel_id varchar(64) not null,
    title mediumtext null,
    channel_desc mediumtext null,
    link varchar(128) null,
    date date null,
    author varchar(128) null,
    input_date datetime null,
    thumbnail varchar(128) null
);

create index rfi_idx_channel_id
    on rss_feed_item (channel_id);

create table if not exists rss_feed_tag
(
    name varchar(64) not null,
    channel_id varchar(64) not null,
    title mediumtext null,
    date date null
);

create index rss_feed_tag_idx_name
    on rss_feed_tag (name);

SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
FLUSH PRIVILEGES;
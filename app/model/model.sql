create table rss_feed_channel
(
    id           varchar(64) not null,
    title        varchar(128) null,
    channel_desc mediumtext null,
    image_url    varchar(128) null,
    link         varchar(128) null,
    constraint rss_feed_channel_pk
        primary key (id)
);

create table rss_feed_item
(
    id           int auto_increment
        primary key,
    channel_id   varchar(64)  not null,
    title        mediumtext   null,
    channel_desc mediumtext   null,
    link         varchar(128) null,
    date         date         null,
    author       varchar(128) null,
    input_date   datetime     null
);


create table rss_feed_tag
(
    name       varchar(64) not null,
    channel_id varchar(64) not null,
    title      mediumtext null,
    date       date null
);

create
index rss_feed_tag_idx_name
    on rss_feed_tag (name);

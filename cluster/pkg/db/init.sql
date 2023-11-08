create table if not exists audio
(
    data         mediumblob  not null,
    speaker      varchar(50) null,
    text_id      int         not null,
    sec          float       null,
    id           bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    written_date timestamp   null
);

create table if not exists job
(
    id           bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    date         date         null,
    max_index    int          null,
    speaker      varchar(50)  null,
    playing_time float        null,
    title        varchar(255) null
);

create table if not exists job_text
(
    job_id  int null,
    text_id int null,
    no      int null
);

create table if not exists text
(
    value   varchar(255) null,
    speaker varchar(50)  null,
    id      bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    constraint value
        unique (value, speaker)
);


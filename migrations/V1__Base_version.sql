create table dogs
(
    owner varchar(250) PRIMARY KEY,
    name   varchar(250),
    gender varchar(250)
);

create table users
(
    login    varchar(250) PRIMARY KEY,
    password varchar(250)
);

create table tokens
(
    login varchar(250) PRIMARY KEY,
    value varchar(250)
)

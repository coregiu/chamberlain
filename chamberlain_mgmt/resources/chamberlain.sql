create database chamberlain;

create table users
(
    Username varchar(16) primary key,
    Password varchar(32),
    Role     varchar(16)
);
create index idx_pass on users(password);

create table inputs
(
    InputTime   bigint primary key,
    Year        int,
    Month       int,
    Type        varchar(16),
    Base        float,
    AllInput    float,
    Tax         float,
    Actual      float,
    Description varchar(512)
);
create index idx_input_month on year, month);


insert into users value ('test', '123456', 'admin');
insert into users value ('test1', '123456', 'user');
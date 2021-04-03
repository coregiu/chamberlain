create table inputs(
InputTime bigint primary key,
Year int,
Month int,
Type varchar(16),
Base float,
AllInput float,
Tax float,
Actual float,
Description varchar(512)
);
create index idx_input_month on Inputs(year, month);

create table users(
Username varchar(16) primary key,
Password varchar(32),
Role varchar(16)
);
create index idx_pass on Users(password);

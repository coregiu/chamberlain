CREATE DATABASE chamberlain;

CREATE TABLE USERS
(
    USERNAME varchar(16) primary key,
    PASSWORD varchar(32),
    ROLE     varchar(16)
);
CREATE INDEX IDX_PASS ON USERS(PASSWORD);

CREATE TABLE INPUTS
(
    INPUT_TIME  bigint primary key,
    YEAR        int,
    MONTH       int,
    TYPE        varchar(16),
    BASE        float,
    ALL_INPUT   float,
    TAX         float,
    ACTUAL      float,
    DESCRIPTION varchar(512)
);
CREATE INDEX IDX_INPUT_MONTH ON YEAR, MONTH);

CREATE TABLE LOGS
(
    LOG_ID       bigint primary key,
    USERNAME     varchar(16),
    OPERATION    varchar(32),
    OP_TIME      timestamp,
    DESCRIPTION  varchar(512)
);

INSERT INTO USERS VALUE ('test', '123456', 'admin');
INSERT INTO USERS VALUE ('test1', '123456', 'user');
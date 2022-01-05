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
CREATE INDEX IDX_INPUT_MONTH ON INPUTS(YEAR, MONTH);

CREATE TABLE LOGS
(
    LOG_ID       bigint primary key,
    USERNAME     varchar(16),
    OPERATION    varchar(64),
    OP_TIME      timestamp,
    OP_RESULT    varchar(16),
    DESCRIPTION  varchar(512)
);
CREATE INDEX IDX_LOGS_OP_TIME ON LOGS(OP_TIME);

CREATE TABLE NOTEBOOK
(
    NOTE_ID      varchar(32) primary key,
    USERNAME     varchar(16),
    CONTENT      text,
    NOTE_TIME    timestamp,
    FINISH_TIME  timestamp,
    STATUS       varchar(16)
);
CREATE INDEX IDX_LOGS_OP_TIME ON LOGS(OP_TIME);

INSERT INTO USERS VALUE ('test', 'test', 'admin');
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
    LOG_ID      bigint primary key,
    USERNAME    varchar(16),
    OPERATION   varchar(64),
    OP_CLIENT   varchar(32),
    OP_TIME     timestamp,
    OP_RESULT   varchar(16),
    DESCRIPTION varchar(512)
);
CREATE INDEX IDX_LOGS_OP_TIME ON LOGS(OP_TIME);

CREATE TABLE NOTEBOOKS
(
    NOTE_ID          varchar(36) primary key,
    USERNAME         varchar(16),
    CONTENT          text,
    LEVEL            char(1),
    NOTE_TIME        timestamp,
    FINISH_TIME      timestamp,
    OWNER            varchar(32),
    REAL_FINISH_TIME timestamp,
    STATUS           varchar(8)
);
CREATE INDEX IDX_NOTEBOOK_FINISH_TIME ON NOTEBOOKS(FINISH_TIME);

CREATE TABLE NOTE_SUMMARIES
(
    BOOK_ID        varchar(36) primary key,
    PARENT_BOOK_ID varchar(36),
    BOOK_NAME      varchar(64),
    USERNAME       varchar(16),
    CONTENT        MEDIUMTEXT,
    BOOK_TIME      timestamp
);
CREATE INDEX IDX_NOTE_SUMMARIES_TIME ON NOTE_SUMMARIES(PARENT_BOOK_ID, BOOK_TIME);

INSERT INTO USERS VALUES ('test', 'test', 'admin');

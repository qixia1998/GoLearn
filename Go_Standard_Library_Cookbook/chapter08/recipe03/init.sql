CREATE DATABASE posttest;
\c posttest

DROP TABLE IF EXISTS post;
CREATE TABLE post (
    ID serial,
    TITLE varchar(40),
    CONNECT varchar(255),
    CONSTRAINT pk_post PRIMARY KEY(ID)
);
SELECT * FROM post;
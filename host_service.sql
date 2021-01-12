CREATE DATABASE host_service;

USE host_service;

CREATE TABLE hosts(
    id       VARCHAR(255),
    ip       TEXT,
    name     TEXT,
    metadata VARCHAR(255),
    status   VARCHAR(255)
);

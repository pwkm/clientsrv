CREATE DATABASE bigprogram;

DROP TABLE client;

CREATE TABLE client (
id UUID PRIMARY KEY ,
name TEXT
);

DROP TABLE login;

CREATE TABLE login (
id UUID PRIMARY KEY,
login TEXT,
password TEXT
);

DROP TABLE profile;

CREATE TABLE profile (
id UUID PRIMARY KEY,
street TEXT,
number TEXT,
postcode NUMERIC,
community TEXT,
email TEXT,
birthday TIMESTAMP
);

DROP TABLE product;

CREATE TABLE product (
id UUID PRIMARY KEY,
name TEXT,
description TEXT,
annualcontribution Numeric
);
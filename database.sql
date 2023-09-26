/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE users (
	user_id           serial PRIMARY KEY,
	full_name         VARCHAR (60) NOT NULL,
  phone_number      BIGINT UNIQUE NOT NULL,
  passwd            VARCHAR (64) NOT NULL,
  salt              VARCHAR (16) NOT NULL,
  created_on        TIMESTAMP NOT NULL,
  successful_login  INTEGER DEFAULT 0 NOT NULL
);

CREATE UNIQUE INDEX idx_phone_number ON users (phone_number);

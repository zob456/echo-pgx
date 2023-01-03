CREATE SCHEMA IF NOT EXISTS "Library";
SET SCHEMA 'Library';

CREATE TABLE IF NOT EXISTS "Author"
(
    "ID"         uuid PRIMARY KEY NOT NULL,
    "First_Name" text             NOT NULL,
    "Last_Name"  text             NOT NULL
);

CREATE TABLE IF NOT EXISTS "Book"
(
    "ID"       uuid PRIMARY KEY NOT NULL,
    "Title"    text,
    "AuthorID" uuid REFERENCES "Library"."Author" ("ID")
);
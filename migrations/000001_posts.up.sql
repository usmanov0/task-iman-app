CREATE TABLE "posts"(
    "id" INTEGER PRIMARY KEY,
    "user_id" INTEGER NOT NULL ,
    "title" VARCHAR,
    "body"  VARCHAR,
    "page"  INTEGER
);
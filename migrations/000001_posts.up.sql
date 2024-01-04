CREATE TABLE "posts"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" INT NOT NULL,
    "title" VARCHAR,
    "body"  VARCHAR
);
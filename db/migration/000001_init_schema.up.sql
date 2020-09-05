CREATE TABLE "shorturl" (
  "id" SERIAL PRIMARY KEY,
  "origin" text,
  "short" text,
  "hits" bigint DEFAULT (0),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" text,
  "password" text,
  "email" text,
  "phone" text,
  "is_active" int DEFAULT (0),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users_action" (
  "id" SERIAL PRIMARY KEY,
  "action_name" text,
  "action_result" text,
  "is_active" int DEFAULT (0),
  "created_at" timestamptz DEFAULT (now())
);
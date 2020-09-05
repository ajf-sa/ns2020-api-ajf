CREATE TABLE "shorturl" (
  "id" SERIAL PRIMARY KEY,
  "origin" text,
  "short" text,
  "created_at" timestamptz DEFAULT (now())
);
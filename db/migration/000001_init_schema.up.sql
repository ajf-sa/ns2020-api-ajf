CREATE TABLE "todos" (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT,
  "done" BOOLEAN DEFAULT (FALSE),
  "created_at" timestamptz DEFAULT (now())
);

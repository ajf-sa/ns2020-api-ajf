CREATE TABLE todos (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
  created_at timestamptz DEFAULT (now())
);
  
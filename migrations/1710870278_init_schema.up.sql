CREATE TABLE "tables" (
  "id" varchar PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);
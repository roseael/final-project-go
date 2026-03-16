CREATE TABLE recipes (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer NOT NULL,
  "title" varchar NOT NULL,
  "instructions" text NOT NULL,
  "prep_time_minutes" integer,
  "created_at" timestamptz DEFAULT (now()),
  CONSTRAINT "fk_user" 
    FOREIGN KEY ("user_id") 
    REFERENCES "users" ("id") 
    ON DELETE CASCADE
);
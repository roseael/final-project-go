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

--Initial inserts
INSERT INTO recipes (user_id, title, instructions) VALUES 
(1, 'Belizean Rice and Beans', 'Cook beans with garlic, add coconut milk and rice, cook over medium heat until rice is fluffy'),
(2, 'Tres Leches Cake', 'Mix the three milks, pour over the sponge cake, let soak and chill for 4 hours.');

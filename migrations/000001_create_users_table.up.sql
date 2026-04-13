CREATE TABLE users (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

--Initial inserts
INSERT INTO users (username, email, password_hash) VALUES
('Rose', 'rose1205@gmail.com', 'rose1230'),
('Aelin', 'aelin2530@gmail.com', 'ael0525'),
('Violet', 'violet0330@gmail.com', 'violence0325');

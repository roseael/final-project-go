CREATE TABLE ingredients (
  "id" SERIAL PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

--Initial inserts
-- Adding a bunch of ingredients at once
INSERT INTO ingredients (name) VALUES 
    ('Flour'),
    ('Sugar'),
    ('Baking Powder'),
    ('Coconut Milk'),
    ('Vanilla Extract'),
    ('Condensed Milk'),
    ('Evaporated Milk'),
    ('Butter'),
    ('Eggs'),
    ('Beans'),
    ('Rice'),
    ('Garlic'),
    ('Onion'),
    ('Chicken'),
    ('Beef'),
    ('Pork');
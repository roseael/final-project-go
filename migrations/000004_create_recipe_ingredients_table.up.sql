CREATE TABLE recipe_ingredients (
  "id" SERIAL PRIMARY KEY,
  "recipe_id" integer NOT NULL,
  "ingredient_id" integer NOT NULL,
  "quantity" varchar, -- e.g., '1/2 cup' or '200g'
  CONSTRAINT "fk_recipe" 
    FOREIGN KEY ("recipe_id") 
    REFERENCES "recipes" ("id") 
    ON DELETE CASCADE,
  CONSTRAINT "fk_ingredient" 
    FOREIGN KEY ("ingredient_id") 
    REFERENCES "ingredients" ("id") 
    ON DELETE CASCADE
);

-- Let's say Tres Leches is Recipe #1 
-- And Flour is Ingredient #1, Sugar is #2, Coconut Milk is #4

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity) VALUES 
    (2, 1, '2 cups'),
    (2, 2, '1 cup'),
    (2, 4, '1 can'),
    (2, 6, '1 can'),
    (2, 7, '1 can');
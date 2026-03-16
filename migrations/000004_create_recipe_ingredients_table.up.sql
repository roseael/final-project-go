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
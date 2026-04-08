package models

import (
	"database/sql"
)

type Recipe struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	Title        string `json:"title"`
	Instructions string `json:"instructions"`
}

type RecipeModel struct {
	DB *sql.DB
}

// Logic for creating a recipe record
func (m RecipeModel) Insert(recipe *Recipe) error {
	query := `INSERT INTO recipes (user_id, title, instructions) 
              VALUES ($1, $2, $3) RETURNING id`
	
	return m.DB.QueryRow(query, recipe.UserID, recipe.Title, recipe.Instructions).Scan(&recipe.ID)
}
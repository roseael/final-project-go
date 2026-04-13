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

func (m *RecipeModel) List() ([]*Recipe, error) {
    // 1. The SQL Sentence
    query := `SELECT id, user_id, title, instructions FROM recipes ORDER BY id`

    // 2. Ask the database
    rows, err := m.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close() // Clean up when we're done

    var recipes []*Recipe

    // 3. Loop through the rows and "scan" them into Go objects
    for rows.Next() {
        var r Recipe
        err := rows.Scan(&r.ID, &r.UserID, &r.Title, &r.Instructions)
        if err != nil {
            return nil, err
        }
        recipes = append(recipes, &r)
    }

    return recipes, nil
}
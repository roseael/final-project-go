package models

import (
	"context"
	"database/sql"
	"time"
)

type Ingredient struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// The Model wrapper (This connects the DB
type IngredientModel struct {
	DB *sql.DB
}

func (m IngredientModel) Insert(name string) (int64, error) {
	query := `INSERT INTO ingredients (name) VALUES ($1) RETURNING id`
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := m.DB.QueryRowContext(ctx, query, name).Scan(&id)
	return id, err
}
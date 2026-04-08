package models

import (
	"context"
	"database/sql"
	"time"
)

// User struct defines the blueprint for a user record.
type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // The hyphen means this never shows up in JSON
}

// UserModel wraps the database connection.
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user to the database.
func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// We pass the pointer to the user struct so we can update the ID after the insert
	return m.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)
}

// Update modifies an existing user's details.
func (m UserModel) Update(user *User) (int64, error) {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, user.Username, user.Email, user.ID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Delete removes a user from the database.
func (m UserModel) Delete(id int64) (int64, error) {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
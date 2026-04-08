package handlers

import (
	"final-project-go/models"
	"net/http"
	"strconv"
)

func (app *Application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	// Validation
	errors := make(map[string]string)
	app.Check(errors, input.Username != "", "username", "must be provided")
	app.Check(errors, input.Email != "", "email", "must be provided")
	if len(errors) > 0 {
		app.failedValidation(w, errors)
		return
	}

	user := &models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: input.Password, // In a real app, hash this!
	}

	if err := app.Users.Insert(user); err != nil {
		app.serverError(w, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
}

func (app *Application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	app.readJSON(w, r, &input)

	user := &models.User{
		ID:       id,
		Username: input.Username,
		Email:    input.Email,
	}

	rowsAffected, err := app.Users.Update(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if rowsAffected == 0 {
		app.notFound(w)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "user updated"}, nil)
}

func (app *Application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	rowsAffected, err := app.Users.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if rowsAffected == 0 {
		app.notFound(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
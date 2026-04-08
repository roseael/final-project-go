package handlers

import (
	"final-project-go/models"
	"net/http"
	"strconv"
)

func (app *Application) CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID       int64  `json:"user_id"`
		Title        string `json:"title"`
		Instructions string `json:"instructions"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	recipe := &models.Recipe{
		UserID:       input.UserID,
		Title:        input.Title,
		Instructions: input.Instructions,
	}

	if err := app.Recipes.Insert(recipe); err != nil {
		app.serverError(w, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"recipe": recipe}, nil)
}

func (app *Application) UpdateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	var input struct {
		Title        string `json:"title"`
		Instructions string `json:"instructions"`
	}

	app.readJSON(w, r, &input)

	// We can reuse the RecipeModel's Update if you add one, 
    // or just use app.DB.ExecContext here for simplicity
	query := `UPDATE recipes SET title = $1, instructions = $2 WHERE id = $3`
	result, err := app.DB.Exec(query, input.Title, input.Instructions, id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		app.notFound(w)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "recipe updated"}, nil)
}

func (app *Application) DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	result, err := app.DB.Exec("DELETE FROM recipes WHERE id = $1", id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		app.notFound(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
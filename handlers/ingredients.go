package handlers

import (
	"net/http"
)

func (app *Application) AddIngredientHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequest(w, err.Error())
		return
	}

	id, err := app.Ingredients.Insert(input.Name)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"id": id, "name": input.Name}, nil)
}

func (app *Application) ListIngredientsHandler(w http.ResponseWriter, r *http.Request) {
    // Standard SELECT query to list all
	query := `SELECT id, name FROM ingredients ORDER BY name ASC`
	rows, err := app.DB.Query(query)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer rows.Close()

	ingredients := []envelope{}
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		ingredients = append(ingredients, envelope{"id": id, "name": name})
	}

	app.writeJSON(w, http.StatusOK, envelope{"ingredients": ingredients}, nil)
}
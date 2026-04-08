package main

import (
	"log"
	"net/http"
	"time"

	"final-project-go/database"
	"final-project-go/handlers"
	"final-project-go/models"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// 1. Database Connection String (DSN)
	dsn := "postgres://recipe_user:password@localhost:5432/recipe_db?sslmode=disable"

	// 2. Open the Database Pool
	db, err := database.OpenDB(dsn)
	if err != nil {
		log.Fatalf("Critical Error: Could not connect to database: %v", err)
	}
	defer db.Close()

	// 3. Initialize the Application Struct
	app := &handlers.Application{
		DB: db,
		Ingredients: models.IngredientModel{DB: db},
		Recipes:     models.RecipeModel{DB: db},
		Users:       models.UserModel{DB: db},
	}

	// 4. Set up the ServeMux (Router)
	mux := http.NewServeMux()

	// --- User Routes ---
	mux.HandleFunc("POST /register", app.RegisterUserHandler)
	mux.HandleFunc("PUT /users/{id}", app.UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", app.DeleteUserHandler)

	// --- Ingredient Routes ---
	mux.HandleFunc("GET /ingredients", app.ListIngredientsHandler)
	mux.HandleFunc("POST /ingredients", app.AddIngredientHandler)

	// --- Recipe Routes ---
	mux.HandleFunc("POST /recipes", app.CreateRecipeHandler)
	mux.HandleFunc("PUT /recipes/{id}", app.UpdateRecipeHandler)
	mux.HandleFunc("DELETE /recipes/{id}", app.DeleteRecipeHandler)

	// --- Health/System Check ---
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "available"}`))
	})

	// 5. Start the HTTP Server
	// We use port 4000 to match your teacher's example
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)
	log.Println("Ready to accept requests...")

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
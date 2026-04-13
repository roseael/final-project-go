package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"database/sql"
	"final-project-go/models"
)

type Application struct {
	DB          *sql.DB
	Ingredients models.IngredientModel
	Recipes     models.RecipeModel
	Users       models.UserModel
}

type envelope map[string]any

func (app *Application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, values := range headers {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Use http.MaxBytesReader to limit the size of the request body to 1MB.
	// This is a "Systems Programming" best practice to prevent DoS attacks.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Initialize the json.Decoder and configure it to disallow unknown fields.
	// This means if the user sends "user_name" instead of "username", it throws an error.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the request body into the destination (dst).
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		// Syntax errors (e.g., missing a comma or a brace)
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Wrong data types (e.g., sending a string when the struct expects an int)
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// Empty body
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Unknown keys (because of DisallowUnknownFields above)
		case strings.Contains(err.Error(), "unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// Request body was too large
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		// A panic-level error (programmer error, like passing a non-pointer to Decode)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	// Call Decode again using a pointer to an empty anonymous struct. 
	// This ensures the request body only contains a single JSON object.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %v", err)
	app.writeJSON(w, http.StatusInternalServerError, envelope{"error": "the server encountered a problem"}, nil)
}

func (app *Application) badRequest(w http.ResponseWriter, msg string) {
	app.writeJSON(w, http.StatusBadRequest, envelope{"error": msg}, nil)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.writeJSON(w, http.StatusNotFound, envelope{"error": "the requested resource could not be found"}, nil)
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (app *Application) Check(errors map[string]string, ok bool, key, message string) {
	if !ok {
		// Only add the error if one doesn't already exist for this key
		if _, exists := errors[key]; !exists {
			errors[key] = message
		}
	}
}

// failedValidation sends a 422 Unprocessable Entity response containing the errors map.
func (app *Application) failedValidation(w http.ResponseWriter, errors map[string]string) {
	app.writeJSON(w, http.StatusUnprocessableEntity, envelope{"errors": errors}, nil)
}
package handlers

import (
	"encoding/json"
	"github.com/mattdotmatt/bigstar/models"
	"github.com/mattdotmatt/bigstar/repositories"
	"gopkg.in/validator.v2"
	"net/http"
)

/*
	Get all the characters in the database
*/
func GetCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := characters.AllCharacters()

		if characters == nil || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

/*
	Save a payload of characters to the database. These replace the existing items in the store
*/
func SaveCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var input []models.Character

		err := decoder.Decode(&input)

		// Validate input
		for _, character := range input {
			if err := validator.Validate(character); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		}

		if err = characters.SaveCharacters(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

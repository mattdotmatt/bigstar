package handlers

import (
	"encoding/json"
	"github.com/mattdotmatt/bigstar/models"
	"github.com/mattdotmatt/bigstar/repositories"
	"net/http"
)

func GetCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := characters.AllCharacters()

		if characters == nil || err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(c)
	}
}

func SaveCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var input []models.Character

		err := decoder.Decode(&input)

		err = characters.SaveCharacters(input)

		if characters == nil || err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}

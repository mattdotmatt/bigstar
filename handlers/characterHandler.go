package handlers

import (
	"encoding/json"
	"github.com/mattdotmatt/bigstar/data"
	"net/http"
)

func GetCharacters(characters data.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := characters.AllCharacters()

		if characters == nil || err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		json.NewEncoder(w).Encode(c)
	}
}

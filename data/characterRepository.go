package data

import "github.com/mattdotmatt/bigstar/models"

type CharacterRepository interface {
	AllCharacters() ([]*models.Character, error)
}

type characterRepository struct {
	*JsonDB `inject:""`
}

func NewCharacterRepository() CharacterRepository {
	return &characterRepository{}
}

func (db *characterRepository) AllCharacters() ([]*models.Character, error) {

	var characters []*models.Character

	character := models.Character{FirstName: "Matt", LastName: "Young"}

	characters = append(characters, &character)

	return characters, nil
}

package repositories

import (
	"github.com/mattdotmatt/bigstar/data"
	"github.com/mattdotmatt/bigstar/models"
)

type CharacterRepository interface {
	AllCharacters() ([]models.Character, error)
	SaveCharacters(characters []models.Character) error
}

type characterRepository struct {
	*data.JsonDB `inject:""`
}

func NewCharacterRepository() CharacterRepository {
	return &characterRepository{}
}

func (db *characterRepository) AllCharacters() ([]models.Character, error) {
	return db.GetAllCharacters()
}

func (db *characterRepository) SaveCharacters(characters []models.Character) error {
	return db.SaveAllCharacters(characters)
}

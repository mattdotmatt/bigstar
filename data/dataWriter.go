package data

import (
	"github.com/mattdotmatt/bigstar/models"
)

type SaveCharacters struct {
	toSave   []models.Character
	exitChan chan error
}

func NewSaveCharacters(characters []models.Character) *SaveCharacters {
	return &SaveCharacters{
		toSave:   characters,
		exitChan: make(chan error, 1),
	}
}
func (save SaveCharacters) ExitChan() chan error {
	return save.exitChan
}

func (save SaveCharacters) Run(characters []models.Character) []models.Character {
	return save.toSave
}

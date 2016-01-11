package data

import (
	"github.com/mattdotmatt/bigstar/models"
)

type ReadCharacters struct {
	characters chan []models.Character
	exitChan   chan error
}

func NewReadCharacters() *ReadCharacters {
	return &ReadCharacters{
		characters: make(chan []models.Character, 1),
		exitChan:   make(chan error, 1),
	}
}
func (read ReadCharacters) ExitChan() chan error {
	return read.exitChan
}

func (read ReadCharacters) Run(characters []models.Character) []models.Character {

	read.characters <- characters

	return nil
}

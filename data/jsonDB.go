package data

import (
	"github.com/mattdotmatt/bigstar/models"
	"io/ioutil"
	"log"
)

type JsonDB struct {
	Requests chan DataRequest
}

func NewJsonDB(fileLocation string) *JsonDB {

	checkDatabaseExists(fileLocation)

	// create channel to communicate over
	request := make(chan DataRequest)

	// start watching request channel for work
	go processRequests(request, fileLocation)

	return &JsonDB{Requests: request}
}

func (db *JsonDB) GetAllCharacters() ([]models.Character, error) {

	request := NewReadCharacters()

	db.Requests <- request

	if err := <-request.ExitChan(); err != nil {
		return nil, err
	}

	return <-request.characters, nil
}

func (db *JsonDB) SaveAllCharacters(characters []models.Character) error {

	request := NewSaveCharacters(characters)

	db.Requests <- request

	if err := <-request.ExitChan(); err != nil {
		return err
	}

	return nil
}

func checkDatabaseExists(fileLocation string) {
	if _, err := ioutil.ReadFile(fileLocation); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(fileLocation, []byte(str), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

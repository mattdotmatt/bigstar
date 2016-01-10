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
	requests := make(chan DataRequest)

	// start watching request channel for work
	go processRequests(requests, fileLocation)

	return &JsonDB{Requests: requests}
}

func (db *JsonDB) GetAllCharacters() ([]models.Character, error) {

	arr := make([]models.Character, 0)

	characters, err := db.getHash()

	if err != nil {
		return arr, err
	}

	for _, value := range characters {
		arr = append(arr, value)
	}
	return arr, nil
}

func (db *JsonDB) SaveAllCharacters(characters []models.Character) error {

	job := NewSaveCharacters(characters)

	db.Requests <- job

	if err := <-job.ExitChan(); err != nil {
		return err
	}
	return nil
}

func (c *JsonDB) getHash() ([]models.Character, error) {

	request := NewReadCharacters()

	c.Requests <- request

	if err := <-request.ExitChan(); err != nil {
		return make([]models.Character, 0), err
	}

	return <-request.characters, nil
}

func checkDatabaseExists(fileLocation string) {
	if _, err := ioutil.ReadFile(fileLocation); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(fileLocation, []byte(str), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

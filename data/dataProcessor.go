package data

import (
	"encoding/json"
	"github.com/mattdotmatt/bigstar/models"
	"io/ioutil"
)

type DataRequest interface {
	ExitChan() chan error
	Run(characters []models.Character) ([]models.Character, error)
}

func processRequests(requests chan DataRequest, db string) {
	for {
		request := <-requests

		// Read the database
		characters := make([]models.Character, 0)
		content, err := ioutil.ReadFile(db)
		if err == nil {
			if err = json.Unmarshal(content, &characters); err == nil {

				c, err := request.Run(characters)

				if err == nil && c != nil {
					b, err := json.Marshal(c)
					if err == nil {
						err = ioutil.WriteFile(db, b, 0644)
					}
				}
			}
		}

		request.ExitChan() <- err
	}
}

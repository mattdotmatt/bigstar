package data

import (
	"encoding/json"
	"github.com/mattdotmatt/bigstar/models"
	"io/ioutil"
)

type DataRequest interface {
	ExitChan() chan error
	Run(characters []models.Character) []models.Character
}

func processRequests(incomingRequest chan DataRequest, db string) {
	for {
		request := <-incomingRequest

		content, err := ioutil.ReadFile(db)

		if err == nil {

			characters := []models.Character{}

			if err = json.Unmarshal(content, &characters); err == nil {

				c := request.Run(characters)

				if c != nil {
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

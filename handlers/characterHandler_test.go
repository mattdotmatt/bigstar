package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mattdotmatt/bigstar/models"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server        *httptest.Server
	characters    []*models.Character
	charactersUrl string
)

type mockDB struct{}

func (mdb *mockDB) AllCharacters() ([]*models.Character, error) {
	return characters, nil
}

func TestCharacterHandler(t *testing.T) {
	Convey("Given characters exist", t, func() {

		characters = make([]*models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		characters = append(characters, &character)

		r := mux.NewRouter()

		r.HandleFunc("/characters", GetCharacters(&mockDB{})).Methods("GET")

		server = httptest.NewServer(r)

		charactersUrl = fmt.Sprintf("%s/characters", server.URL)

		Convey("When I get characters", func() {

			request, err := http.NewRequest("GET", charactersUrl, nil)

			res, err := http.DefaultClient.Do(request)

			Convey("Then a 200 should be returned", func() {
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
			})

			Convey("And the correct character data should be returned", func() {
				content, _ := ioutil.ReadAll(res.Body)

				var character []models.Character

				json.Unmarshal(content, &character)

				So(len(character), ShouldEqual, 1)
				So(character[0].FirstName, ShouldEqual, "Matt")
				So(character[0].LastName, ShouldEqual, "Young")
			})
		})
	})
}

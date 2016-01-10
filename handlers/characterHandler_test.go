package handlers

import (
	"encoding/json"
	"errors"
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
	characters    []models.Character
	charactersUrl string
	dataError     error
)

type mockDB struct{}

func (mdb *mockDB) AllCharacters() ([]models.Character, error) {
	return characters, dataError
}

func (mdb *mockDB) SaveCharacters(characters []models.Character) error {
	return dataError
}

func TestCharacterHandler(t *testing.T) {
	Convey("Given characters exist", t, func() {

		characters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		characters = append(characters, character)

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

	Convey("Given no characters exist", t, func() {

		characters = make([]models.Character, 0)

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

			Convey("And empty character data should be returned", func() {
				content, _ := ioutil.ReadAll(res.Body)

				var character []models.Character

				json.Unmarshal(content, &character)

				So(len(character), ShouldEqual, 0)
			})
		})
	})

	Convey("Given characters exist", t, func() {

		characters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		characters = append(characters, character)

		r := mux.NewRouter()

		r.HandleFunc("/characters", GetCharacters(&mockDB{})).Methods("GET")

		server = httptest.NewServer(r)

		charactersUrl = fmt.Sprintf("%s/characters", server.URL)

		Convey("When there is an error getting characters", func() {

			dataError = errors.New("BOOM")

			request, err := http.NewRequest("GET", charactersUrl, nil)

			res, err := http.DefaultClient.Do(request)

			Convey("Then a 400 should be returned", func() {
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 400)
			})
		})
	})

}

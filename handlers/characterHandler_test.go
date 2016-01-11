package handlers

import (
	"bytes"
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
	server           *httptest.Server
	storedCharacters []models.Character
	charactersUrl    string
	dataError        error
)

type mockDB struct{}

func (mdb *mockDB) AllCharacters() ([]models.Character, error) {
	return storedCharacters, dataError
}

func (mdb *mockDB) SaveCharacters(characters []models.Character) error {

	storedCharacters = characters

	return dataError
}

func TestGetCharacters(t *testing.T) {

	Convey("Given characters exist", t, func() {

		storedCharacters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		storedCharacters = append(storedCharacters, character)

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

		storedCharacters = make([]models.Character, 0)

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

		storedCharacters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		storedCharacters = append(storedCharacters, character)

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

func TestSaveCharacters(t *testing.T) {

	Convey("Given characters exist", t, func() {

		dataError = nil
		storedCharacters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		storedCharacters = append(storedCharacters, character)

		r := mux.NewRouter()

		r.HandleFunc("/characters", SaveCharacters(&mockDB{})).Methods("POST")

		server = httptest.NewServer(r)

		charactersUrl = fmt.Sprintf("%s/characters", server.URL)

		Convey("When I post characters", func() {

			characters := append(storedCharacters, models.Character{"New", "User"})

			jsonBody, _ := json.Marshal(characters)

			body := bytes.NewReader(jsonBody)

			request, err := http.NewRequest("POST", charactersUrl, body)

			res, err := http.DefaultClient.Do(request)

			Convey("Then a 200 should be returned", func() {
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
			})

			Convey("And the new character data should be saved", func() {
				So(len(storedCharacters), ShouldEqual, 2)
				So(storedCharacters[0].FirstName, ShouldEqual, "Matt")
				So(storedCharacters[0].LastName, ShouldEqual, "Young")
				So(storedCharacters[1].FirstName, ShouldEqual, "New")
				So(storedCharacters[1].LastName, ShouldEqual, "User")
			})
		})
	})

	Convey("Given characters exist", t, func() {

		dataError = nil
		storedCharacters = make([]models.Character, 0)
		character := models.Character{FirstName: "Matt", LastName: "Young"}
		storedCharacters = append(storedCharacters, character)

		r := mux.NewRouter()

		r.HandleFunc("/characters", SaveCharacters(&mockDB{})).Methods("POST")

		server = httptest.NewServer(r)

		charactersUrl = fmt.Sprintf("%s/characters", server.URL)

		Convey("When I post incomplete characters", func() {

			characters := append(storedCharacters, models.Character{"", "User"})

			jsonBody, _ := json.Marshal(characters)

			body := bytes.NewReader(jsonBody)

			request, err := http.NewRequest("POST", charactersUrl, body)

			res, err := http.DefaultClient.Do(request)

			Convey("Then a 400 should be returned", func() {
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 400)
			})
		})
	})
}

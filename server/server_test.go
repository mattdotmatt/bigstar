package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/mattdotmatt/bigstar/data"
	"github.com/mattdotmatt/bigstar/models"
	"github.com/mattdotmatt/bigstar/repositories"
	"github.com/mattdotmatt/bigstar/routers"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	ts *httptest.Server
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func initiateTests() {
	var router routers.Router
	var graph inject.Graph

	setupTestData()

	db := data.NewJsonDB("../data/testing/testingDB.json")

	if err := graph.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: repositories.NewCharacterRepository()},
		&inject.Object{Value: &router}); err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Fatalf("Error populating dependencies: ", err.Error())
	}

	n := negroni.Classic()

	n.UseHandler(router.NewRouter())

	ts = httptest.NewServer(n)
}

func setupTestData() {
	characters := []models.Character{
		models.Character{"Matt", "Young"},
		models.Character{"Another", "person"},
		models.Character{"FirstOnly", ""},
		models.Character{"Kerry", "O'Sullivan"},
	}

	b, err := json.Marshal(characters)
	if err == nil {
		err = ioutil.WriteFile("../data/testing/testingDB.json", b, 0644)
	}
}

func TestIntegrationTests(t *testing.T) {
	Convey("Given a connected system", t, func() {

		initiateTests()

		Convey("When I get the characters", func() {

			request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/characters", ts.URL), nil)

			res, err := http.DefaultClient.Do(request)

			Convey("Then a 200 should be returned", func() {
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
			})

			Convey("And the character data should be returned", func() {
				content, _ := ioutil.ReadAll(res.Body)

				var characters []models.Character

				json.Unmarshal(content, &characters)

				So(len(characters), ShouldEqual, 4)
				So(characters[0].FirstName, ShouldEqual, "Matt")
				So(characters[0].LastName, ShouldEqual, "Young")
				So(characters[1].FirstName, ShouldEqual, "Another")
				So(characters[1].LastName, ShouldEqual, "person")
				So(characters[2].FirstName, ShouldEqual, "FirstOnly")
				So(characters[2].LastName, ShouldEqual, "")
				So(characters[3].FirstName, ShouldEqual, "Kerry")
				So(characters[3].LastName, ShouldEqual, "O'Sullivan")
			})
		})

		Convey("When I update the characters", func() {

			character := []models.Character{models.Character{"New", "User"}}

			jsonBody, _ := json.Marshal(character)

			body := bytes.NewReader(jsonBody)

			request, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/characters", ts.URL), body)

			http.DefaultClient.Do(request)

			Convey("And I get the characters", func() {

				request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/characters", ts.URL), nil)

				res, err := http.DefaultClient.Do(request)

				Convey("Then a 200 should be returned", func() {
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)
				})

				Convey("And the updated character data should be returned", func() {
					content, _ := ioutil.ReadAll(res.Body)

					var characters []models.Character

					json.Unmarshal(content, &characters)

					So(len(characters), ShouldEqual, 1)
					So(characters[0].FirstName, ShouldEqual, "New")
					So(characters[0].LastName, ShouldEqual, "User")
				})
			})
		})

		Reset(func() {
			ts.Close()
		})
	})
}

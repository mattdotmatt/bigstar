package server

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/mattdotmatt/bigstar/data"
	"github.com/mattdotmatt/bigstar/data/testing"
	"github.com/mattdotmatt/bigstar/routers"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	db *data.JsonDB
	ts *httptest.Server
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func initiateTests() {
	var router routers.Router
	var graph inject.Graph

	db, err := data_testing.NewTestingDb("./data/testing/testingDB.json")

	if err != nil {
		panic(err)
	}

	if err := graph.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: data.NewCharacterRepository()},
		&inject.Object{Value: &router}); err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Fatalf("Error populating dependencies: ", err.Error())
	}

	n := negroni.Classic()

	n.UseHandler(router.NewRouter())

	ts = httptest.NewServer(n)

	if err != nil {
		panic("Error: " + err.Error())
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

				var id string

				json.Unmarshal(content, &id)

				So(id, ShouldNotBeBlank)
			})
		})

		Reset(func() {
			ts.Close()
		})
	})
}

package server

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/mattdotmatt/bigstar/data"
	"github.com/mattdotmatt/bigstar/repositories"
	"github.com/mattdotmatt/bigstar/routers"
	"log"
	"net/http"
)

func Start(port int, fileLocation string) {

	var router routers.Router
	var graph inject.Graph

	// Create database
	db := data.NewJsonDB(fileLocation)

	// Setup DI
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

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), n)

	if err != nil {
		panic("Error: " + err.Error())
	}
}

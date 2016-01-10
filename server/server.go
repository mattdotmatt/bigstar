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

func Start(siteUrl string, port int, fileLocation string) {

	var router routers.Router
	var graph inject.Graph

	// create database
	db := data.NewJsonDB(fileLocation)

	if err := graph.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: repositories.NewCharacterRepository()},
		&inject.Object{Value: &router}); err != nil {
		log.Fatalf("Error pproviding dependencies: ", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Fatalf("Error populating dependencies: ", err.Error())
	}

	n := negroni.Classic()

	n.UseHandler(router.NewRouter())

	err := StartListening(siteUrl, port, n)

	if err != nil {
		panic("Error: " + err.Error())
	}
}

func StartListening(siteUrl string, port int, router *negroni.Negroni) error {

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		return err
	}

	return nil
}

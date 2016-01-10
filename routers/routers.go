package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mattdotmatt/bigstar/data"
	"github.com/mattdotmatt/bigstar/handlers"
	"net/http"
)

type Router struct {
	CharacterRepository data.CharacterRepository `inject:""`
}

func (router Router) NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(ApiHeaderMiddleware),
		negroni.Wrap(apiRouter(router.CharacterRepository)),
	))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r

}

func apiRouter(characters data.CharacterRepository) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/characters", handlers.GetCharacters(characters)).Methods("GET")

	return r
}

func ApiHeaderMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}

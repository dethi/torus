package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/pilu/xrequestid"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

const (
	VersionRoute = "api.version"
	SearchRoute  = "api.search"
)

var router *mux.Router
var decoder = schema.NewDecoder()

// NewWebRouter returns a router with all the Web routes.
func NewWebRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/api/version").Name(VersionRoute)
	r.Path("/api/search/{tracker}").Name(SearchRoute)
	return r
}

func init() {
	router = NewWebRouter()
	router.Get(VersionRoute).HandlerFunc(VersionHandler)
	router.Get(SearchRoute).HandlerFunc(SearchHandler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(xrequestid.New(3))
	n.Use(negroni.NewLogger())
	n.Use(cors.Default())
	n.UseHandler(router)

	http.Handle("/api/", n)
}

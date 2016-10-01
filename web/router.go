package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pilu/xrequestid"
	"github.com/urfave/negroni"
)

const (
	IndexRoute     = "web.index"
	DashboardRoute = "web.dashboard"
)

var router *mux.Router

// NewWebRouter returns a router with all the Web routes.
func NewWebRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").Name(IndexRoute)
	r.Path("/dashboard").Name(DashboardRoute)
	return r
}

func init() {
	router = NewWebRouter()
	router.Get(IndexRoute).HandlerFunc(IndexHandler)
	router.Get(DashboardRoute).HandlerFunc(DashboardHandler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(xrequestid.New(3))
	n.UseHandler(router)

	http.Handle("/", n)
}

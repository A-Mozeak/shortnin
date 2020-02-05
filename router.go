package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var storage = NewDB()

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{"Store", "GET", "/list", storage.Get},
	{"Stats", "GET", "/stats", storage.Stats},
	{"Redirect", "GET", "/{short}", storage.Goto},
	{"Generate", "POST", "/create", storage.PutURL},
}

// I considered using a "/stats/{shortlink}" interface but that seems better suited to CMS

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Initializing the storage here, potential to expose an interface instead for polymorphism.
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
	{"Store", "GET", "/list", storage.List},
	{"Stats", "GET", "/stats", storage.Stats},
	{"Redirect", "GET", "/{short}", storage.Redir},
	{"Generate", "POST", "/create", storage.Generate},
}

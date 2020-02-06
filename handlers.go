package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Generate handles requests to generate a shortlink. Responds with the requested shortlink.
func (d *MockDB) Generate(w http.ResponseWriter, r *http.Request) {
	query := GetQueryParams(r)
	shorty := Shorty{"", "", time.Now(), time.Now(), 0, 0, []int{}}

	// Gather the original URL from the request. Error check in case URL wasn't supplied.
	var orig string
	val, ok := query["url"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url parameter not supplied"))
	}
	orig = val[0]

	// Check if a custom shortlink has been provided. If so, assign it to the new Shorty.
	// If not, generate a random link and assign that.
	if custom, ok := query["custom"]; ok {
		shorty.Name = custom[0]
		shorty.Link = orig
		d.shortMap[shorty.Name] = &shorty
	} else if _, there := d.longMap[orig]; !there {
		shorty.Name = GenerateShort()
		shorty.Link = orig
		d.longMap[orig] = shorty.Name
		d.shortMap[shorty.Name] = &shorty
	}

	w.Write([]byte(shorty.Name))
}

// List marshals the short and long URLs that have been added to the server
// and writes them in a response.
func (d MockDB) List(w http.ResponseWriter, r *http.Request) {
	var items []struct {
		ShortURL string `json:"shortURL"`
		LongURL  string `json:"longURL"`
	}

	for _, item := range d.shortMap {
		items = append(items, struct {
			ShortURL string `json:"shortURL"`
			LongURL  string `json:"longURL"`
		}{
			item.Name,
			item.Link,
		})
	}

	json.NewEncoder(w).Encode(items)
}

// Redir handles redirection from shortlinks to external pages.
func (d *MockDB) Redir(w http.ResponseWriter, r *http.Request) {
	// Using mux to get a query-less path that is not specified in the handlers.
	shortStr := mux.Vars(r)["short"]
	var shortObj *Shorty

	// If a regular URL is passed, do a lookup to get its shortlink and the associated Shorty.
	// If a valid shortlink is passed, just get the associated Shorty.
	_, ok := d.shortMap[shortStr]
	if !ok {
		shortStr = d.longMap[shortStr]
		shortObj = d.shortMap[shortStr]
	} else {
		shortObj = d.shortMap[shortStr]
	}

	// If an associated Shorty is found, update its stats and redirect to the given site.
	// Otherwise, respond with an error.
	var original string
	if shortObj != nil {
		UpdateStats(shortObj)
		original = "https://" + shortObj.Link
		http.Redirect(w, r, original, http.StatusFound)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("unable to resolve shortlink"))
	}

}

// Stats marshals the stats of a given link to JSON and sends them in a response.
func (d MockDB) Stats(w http.ResponseWriter, r *http.Request) {
	query := GetQueryParams(r)

	val, ok := query["link"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("shortlink not provided"))
	}

	shortStr := val[0]
	shortObj := d.shortMap[shortStr]
	json.NewEncoder(w).Encode(shortObj)
}

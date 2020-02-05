package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (d *MockDB) Generate(w http.ResponseWriter, r *http.Request) {

	query := GetQueryParams(r)

	var orig string
	val, ok := query["url"]
	if !ok {
		panic(errors.New("url not supplied"))
	}
	orig = val[0]

	shorty := Shorty{"", "", time.Now(), time.Now(), 0, 0, []int{}}
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
	log.Println(items)
	json.NewEncoder(w).Encode(items)
}

func (d *MockDB) Redir(w http.ResponseWriter, r *http.Request) {
	shortStr := mux.Vars(r)["short"]
	var shortObj *Shorty
	_, ok := d.shortMap[shortStr]
	if !ok {
		shortStr = d.longMap[shortStr]
		shortObj = d.shortMap[shortStr]
	} else {
		shortObj = d.shortMap[shortStr]
	}

	var original string
	if shortObj != nil {
		UpdateStats(shortObj)
		original = "https://" + shortObj.Link
	} else {
		original = "https://www.google.com"
	}

	log.Println(original)
	http.Redirect(w, r, original, http.StatusFound)
}

func (d MockDB) Stats(w http.ResponseWriter, r *http.Request) {
	query := GetQueryParams(r)
	log.Println(query)

	val, ok := query["link"]
	if !ok {
		panic(errors.New("shortlink not supplied"))
	}

	shortStr := val[0]
	shortObj := d.shortMap[shortStr]
	log.Println("pre-encoding")
	json.NewEncoder(w).Encode(shortObj)
	log.Println("post-encoding")
}

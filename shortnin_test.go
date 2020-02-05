package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestMainAPI(t *testing.T) {
	t.Run("generating a short url", func(t *testing.T) {
		fmt.Println(GenerateShort())
	})
	t.Run("long URL creates new Shorty", func(t *testing.T) {
		storage := NewDB()
		req := httptest.NewRequest("POST", "/create?url=www.example.com", nil)
		res := httptest.NewRecorder()
		storage.Generate(res, req)

		_, got := storage.longMap["www.example.com"]
		want := true

		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
	t.Run("storage update cross check", func(t *testing.T) {
		storage := NewDB()
		req := httptest.NewRequest("POST", "/create?url=www.example.com", nil)
		res := httptest.NewRecorder()
		storage.Generate(res, req)

		shortlink, _ := storage.longMap["www.example.com"]
		model, _ := storage.shortMap[shortlink]

		shortMatch := model.Name == shortlink
		longMatch := model.Link == "www.example.com"

		if !shortMatch || !longMatch {
			t.Errorf("posted data did not match, short: %q, long: %q", model.Name, model.Link)
		}
	})
	t.Run("stats update", func(t *testing.T) {
		storage := NewDB()
		req := httptest.NewRequest("POST", "/create?url=www.example.com&custom=mylink", nil)
		res := httptest.NewRecorder()
		storage.Generate(res, req)

		model := storage.shortMap["mylink"]
		UpdateStats(model)
		UpdateStats(model)

		got := model.TotalVis
		want := 2

		if got != want {
			t.Errorf("link stats did not update, got %d, wanted %d", got, want)
		}
	})
	t.Run("redirect", func(t *testing.T) {
		storage := NewDB()
		req := httptest.NewRequest("POST", "/create?url=www.wikipedia.org&custom=mylink", nil)
		res := httptest.NewRecorder()
		storage.Generate(res, req)

		req2 := httptest.NewRequest("GET", "/mylink", nil)
		res2 := httptest.NewRecorder()
		storage.Redir(res2, req2)
	})
}

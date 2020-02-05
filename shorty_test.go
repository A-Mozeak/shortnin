package main

import (
	"fmt"
	"testing"
)

func TestBundling(t *testing.T) {
	t.Run("generating a short url", func(t *testing.T) {
		fmt.Println(GenerateShort())
	})
	t.Run("long URL creates new Shorty", func(t *testing.T) {
		storage := NewDB()

		got := storage.PutLong("www.example.com")
		want := 200

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
	t.Run("return 400 long already exists", func(t *testing.T) {
		storage := NewDB()

		storage.PutLong("www.example.com")
		got := storage.PutLong("www.example.com")
		want := 400

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
}

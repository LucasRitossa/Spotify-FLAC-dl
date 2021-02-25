package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
  "Spotify-FLAC-dl/Handler"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", Handler.GetAudio)
	http.ListenAndServe(":8080", r)
}


package main

import (
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Index)

	mux.Route("/dsa", func(r chi.Router) {
		r.Get("/bubble-sort", handlers.Repo.BubbleSort)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

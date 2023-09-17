package main

import (
	"log"
	"net/http"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/handlers"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
)

func main() {

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template Cache")
	}

	app.TemplateCache = templateCache

	http.HandleFunc("/", handlers.Index)

	http.ListenAndServe(":8080", nil)
}

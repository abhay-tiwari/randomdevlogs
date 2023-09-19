package main

import (
	"log"
	"net/http"
	"time"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/handlers"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template Cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}

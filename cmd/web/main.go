package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/driver"
	"github.com/abhay-tiwari/randomdevlogs/pkg/handlers"
	"github.com/abhay-tiwari/randomdevlogs/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "use template cache")
	dbHost := flag.String("dbhost", "localhost", "Database Host")
	dbName := flag.String("dbname", "", "Database Name")
	dbUser := flag.String("dbuser", "", "Database User")
	dbPassword := flag.String("dbpassword", "", "Database Password")
	dbPort := flag.String("dbport", "5432", "Database Port")
	dbSSL := flag.String("dbssl", "disable", "Database SSL settings (disable, prefer, require)")

	flag.Parse()

	app.InProduction = *inProduction

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	log.Println("Connecting to database...")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPassword, *dbSSL)

	db, err := driver.ConnectSQL(connectionString)

	defer db.SQL.Close()

	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	log.Println("Connected to database.")

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template Cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = *useCache

	repo := handlers.NewRepo(&app, db)

	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}

package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/dbDriver"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/handlers"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8082"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.Prod = true

	session = scs.New()
	session.Lifetime = 24 * time.Second
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Prod

	app.Session = session

	// Connect to db
	log.Println("Connecting to DB...")
	db, err := dbDriver.ConnectSQL("season2")
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	log.Println("Connected to DB")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

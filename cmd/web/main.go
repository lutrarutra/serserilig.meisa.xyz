package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
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

	app.Prod = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Prod

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

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

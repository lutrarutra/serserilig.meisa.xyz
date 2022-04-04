package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/handlers"
	"net/http"
)

func routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/about", handlers.Repo.About)
	mux.Get("/edit", handlers.Repo.Edit)
	mux.Get("/standings", handlers.Repo.Standings)
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/db-all-drivers", handlers.Repo.GetAllDrivers)

	// HTML static files location
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

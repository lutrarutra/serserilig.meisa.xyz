package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iMeisa/serserilig.meisa.xyz/pkg/config"
	"github.com/iMeisa/serserilig.meisa.xyz/pkg/handlers"
	"net/http"
)

func routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/about", handlers.Repo.About)
	mux.Get("/standings", handlers.Repo.Standings)
	mux.Get("/", handlers.Repo.Home)

	return mux
}

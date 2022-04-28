package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/handlers"
)

func routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Public routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/grid", handlers.Repo.Grid)
	mux.Get("/rules", handlers.Repo.Rules)
	mux.Get("/staff", handlers.Repo.Staff)
	mux.Get("/standings", handlers.Repo.Standings)
	mux.Get("/calendar", handlers.Repo.Calendar)

	// Admin routes
	mux.Get("/edit/drivers", handlers.Repo.EditDrivers)
	mux.Get("/edit/teams", handlers.Repo.EditTeams)
	mux.Get("/edit/calendar", handlers.Repo.EditCalendar)
	mux.Get("/edit/staff", handlers.Repo.EditStaff)
	mux.Get("/edit", handlers.Repo.Edit)

	// API routes
	mux.Get("/api/driver", handlers.Repo.GetDriver)
	mux.Get("/api/drivers", handlers.Repo.GetAllDrivers)
	mux.Get("/api/teams", handlers.Repo.GetAllTeams)
	mux.Get("/api/calendar", handlers.Repo.GetAllRaces)
	mux.Get("/api/staff", handlers.Repo.GetAllStaff)

	mux.Get("/api/drivers/add", handlers.Repo.AddDriver)
	mux.Get("/api/drivers/delete", handlers.Repo.DeleteDriver)
	mux.Get("/api/drivers/update", handlers.Repo.UpdateDriver)

	mux.Get("/api/teams/update", handlers.Repo.UpdateTeam)

	mux.Get("/api/races/add", handlers.Repo.AddRace)
	mux.Get("/api/races/delete", handlers.Repo.DeleteRace)
	mux.Get("/api/races/update", handlers.Repo.UpdateRace)

	mux.Post("/api/staff/add", handlers.Repo.AddStaff)
	mux.Get("/api/staff/delete", handlers.Repo.DeleteStaff)
	mux.Get("/api/staff/update", handlers.Repo.UpdateStaff)

	// HTML static files location
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

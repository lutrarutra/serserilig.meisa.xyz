package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/dbDriver"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository/dbrepo"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *dbDriver.DB) *Repository {
	return &Repository{
		App: a,
		DB: dbrepo.NewSqliteRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := m.DB.QueryAllDrivers()
	if err != nil {
		log.Fatal(err)
	}

	driversJSON, err := json.Marshal(drivers)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(driversJSON)
}

func (m *Repository) Edit(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "edit.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Standings(w http.ResponseWriter, r *http.Request) {
	newDriver := models.Driver{
		Name: "New Guy",
		TeamID: -1,
	}

	err := m.DB.InsertDriver(newDriver)
	if err != nil {
		fmt.Println(err)
	}

	render.Template(w, "standings.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

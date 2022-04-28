package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/dbDriver"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository/dbrepo"
)

// Main handlers file

// Repo the repository used by the handlers
var Repo *Repository
var templateData = &models.TemplateData{}

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *dbDriver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewSqliteRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) BuildTables() {
	err := m.DB.CreateDriverTable()
	if err != nil {
		log.Fatal(err)
	}
	err = m.DB.CreateTeamTable()
	if err != nil {
		log.Fatal(err)
	}
	err = m.DB.CreateCalendarTable()
	if err != nil {
		log.Fatal(err)
	}
	err = m.DB.CreateStaffTable()
	if err != nil {
		log.Fatal(err)
	}

	m.updateDriverJSON()
	m.updateTeamJSON()
	m.updateCalendarJSON()
	m.updateStaffJSON()
}

func (m *Repository) updateCalendarJSON() {
	races, err := m.DB.QueryAllRaces()
	if err != nil {
		log.Fatal(err)
	}

	calendarJSON, err := json.MarshalIndent(races, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	err = ioutil.WriteFile("./static/json/calendar.json", calendarJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

func (m *Repository) updateStaffJSON() {
	staff, err := m.DB.QueryAllStaff()
	if err != nil {
		log.Fatal(err)
	}

	staffJSON, err := json.MarshalIndent(staff, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	err = ioutil.WriteFile("./static/json/staff.json", staffJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

func (m *Repository) updateDriverJSON() {
	drivers, err := m.DB.QueryAllDrivers()
	if err != nil {
		log.Fatal(err)
	}

	driversJSON, err := json.MarshalIndent(drivers, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	err = ioutil.WriteFile("./static/json/drivers.json", driversJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

func (m *Repository) updateTeamJSON() {
	teams, err := m.DB.QueryAllTeams()
	if err != nil {
		log.Fatal(err)
	}

	teamsJSON, err := json.MarshalIndent(teams, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	err = ioutil.WriteFile("./static/json/teams.json", teamsJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

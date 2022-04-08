package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/dbDriver"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository/dbrepo"
	"log"
	"os"
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

func(m *Repository) UpdateDriverJSON() {
	drivers, err := m.DB.QueryAllDrivers()
	if err != nil {
		log.Fatal(err)
	}

	driversJSON, err := json.MarshalIndent(drivers, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}
	fmt.Println("Drivers JSON: ", string(driversJSON))

	err = os.WriteFile("./static/json/drivers.json", driversJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

func(m *Repository) UpdateTeamJSON() {
	teams, err := m.DB.QueryAllTeams()
	if err != nil {
		log.Fatal(err)
	}

	teamsJSON, err := json.MarshalIndent(teams, "", "	")
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}
	fmt.Println("Teams JSON: ", string(teamsJSON))

	err = os.WriteFile("./static/json/teams.json", teamsJSON, 0644)
	if err != nil {
		log.Println("Could not write to file")
	}
}

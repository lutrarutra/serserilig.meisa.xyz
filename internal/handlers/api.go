package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (m *Repository) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	m.DB.CreateDriverTable()

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

func (m *Repository) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	m.DB.CreateTeamTable()

	teams, err := m.DB.QueryAllTeams()
	if err != nil {
		log.Fatal(err)
	}

	teamsJSON, err := json.Marshal(teams)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(teamsJSON)
}

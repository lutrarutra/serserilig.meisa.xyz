package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (m *Repository) AddDriver(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println("params: ", params)

	if len(params) < 1 {
		log.Println("No key received")
		return
	}

	for key, value := range params {
		fmt.Println(key, value)
	}
}

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

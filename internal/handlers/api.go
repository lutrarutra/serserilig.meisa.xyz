package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/network"
	"log"
	"net/http"
)

func (m *Repository) AddDriver(w http.ResponseWriter, r *http.Request) {
	remoteIP, ok := r.URL.Query()["ip"]
	if !ok || len(remoteIP) < 1 {
		w.Write([]byte("ip not ok"))
		return
	}

	if remoteIP[0] != network.GetRealIP(r) {
		w.Write([]byte("Invalid request"))
		return
	}

	name, ok := r.URL.Query()["name"]
	if !ok || len(name) < 1 {
		w.Write([]byte("name not ok"))
		return
	}

	newDriver := models.Driver{
		Name:   name[0],
		TeamID: -1,
	}

	err := m.DB.InsertDriver(newDriver)
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	w.Write([]byte("Added Driver"))
}

func (m *Repository) GetDriver(w http.ResponseWriter, r *http.Request) {
	m.DB.CreateDriverTable()

	var driverJSON []byte

	query := r.URL.Query()
	if len(query) < 1 {
		w.Write([]byte("Invalid request"))
		return
	} else if val, ok := query["id"]; ok {
		driver, err := m.DB.QueryDriver("id", val[0])
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		driverJSON, err = json.Marshal(driver)
		if err != nil {
			log.Fatal("Could not convert to JSON:", err)
		}
	} else if val, ok = query["name"]; ok {
		driver, err := m.DB.QueryDriver("name", val[0])
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		driverJSON, err = json.Marshal(driver)
		if err != nil {
			log.Fatal("Could not convert to JSON:", err)
		}
	} else {
		w.Write([]byte("Invalid request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(driverJSON)
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

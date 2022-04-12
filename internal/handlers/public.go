package handlers

import (
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
	"log"
	"net/http"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Grid(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "grid.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Rules(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rules.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Staff(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "staff.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Standings(w http.ResponseWriter, r *http.Request) {
	drivers, err := m.DB.QueryAllDrivers()
	if err != nil {
		log.Println(err)
	}

	teams, err := m.DB.QueryAllTeams()
	if err != nil {
		log.Println(err)
	}

	dataMap := make(map[string]interface{})
	teamColors := make(map[int]string)
	for _, team := range teams {
		teamColors[team.ID] = team.Color
	}
	dataMap["team_colors"] = teamColors

	render.Template(w, r, "standings.page.tmpl", &models.TemplateData{
		Drivers: drivers,
		Teams: teams,
		Data: dataMap,
	})
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

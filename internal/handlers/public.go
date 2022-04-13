package handlers

import (
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
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
	templateData.GetDrivers()
	templateData.GetTeams()

	dataMap := make(map[string]interface{})
	driversByID := make(map[int]models.Driver)
	var reserveDrivers []models.Driver
	for _, driver := range templateData.Drivers {
		if driver.TeamID == -1 {
			reserveDrivers = append(reserveDrivers, driver)
		}
		driversByID[driver.ID] = driver
	}

	dataMap["drivers_by_id"] = driversByID
	dataMap["reserve_drivers"] = reserveDrivers

	intMap := make(map[string]int)
	intMap["reserve_count"] = len(reserveDrivers)

	templateData.Data = dataMap

	render.Template(w, r, "grid.page.tmpl", templateData)
}

func (m *Repository) Rules(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rules.page.tmpl", templateData)
}

func (m *Repository) Staff(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "staff.page.tmpl", templateData)
}

func (m *Repository) Standings(w http.ResponseWriter, r *http.Request) {
	templateData.GetDrivers()
	templateData.GetTeams()

	dataMap := make(map[string]interface{})
	teamColors := make(map[int]string)
	for _, team := range templateData.Teams {
		teamColors[team.ID] = team.Color
	}
	dataMap["team_colors"] = teamColors

	templateData.Data = dataMap

	render.Template(w, r, "standings.page.tmpl", templateData)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, r, "home.page.tmpl", templateData)
}

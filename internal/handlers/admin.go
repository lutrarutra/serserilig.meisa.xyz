package handlers

import (
	"net/http"

	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
)

func (m *Repository) EditDrivers(w http.ResponseWriter, r *http.Request) {
	templateData.GetDrivers()
	templateData.GetTeams()
	render.Template(w, r, "editdrivers.page.tmpl", templateData)
}

func (m *Repository) EditTeams(w http.ResponseWriter, r *http.Request) {
	templateData.GetDrivers()
	templateData.GetTeams()

	render.Template(w, r, "editteams.page.tmpl", templateData)
}

func (m *Repository) EditCalendar(w http.ResponseWriter, r *http.Request) {
	templateData.GetRaces()
	templateData.GetGrandPrixes()

	render.Template(w, r, "editcalendar.page.tmpl", templateData)
}

func (m *Repository) Edit(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "edit.page.tmpl", templateData)
}

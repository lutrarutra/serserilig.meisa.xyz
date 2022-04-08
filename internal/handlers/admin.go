package handlers

import (
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
	"net/http"
)

func (m *Repository) EditDrivers(w http.ResponseWriter, r *http.Request) {
	m.UpdateDriverJSON()
	render.Template(w, r, "editdrivers.page.tmpl", &models.TemplateData{})
}

func (m *Repository) EditTeams(w http.ResponseWriter, r *http.Request) {
	m.UpdateDriverJSON()
	m.UpdateTeamJSON()
	render.Template(w, r, "editteams.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Edit(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "edit.page.tmpl", &models.TemplateData{})
}

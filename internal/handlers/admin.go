package handlers

import (
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/render"
	"net/http"
)

func (m *Repository) EditDrivers(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "editdrivers.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Edit(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "edit.page.tmpl", &models.TemplateData{})
}

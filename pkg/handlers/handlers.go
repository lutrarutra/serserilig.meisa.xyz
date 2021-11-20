package handlers

import (
	"meisa_xyz/pkg/config"
	"meisa_xyz/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.Template(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {
	render.Template(w, "about.page.tmpl")
}

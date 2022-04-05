package handlers

import (
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/dbDriver"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository/dbrepo"
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

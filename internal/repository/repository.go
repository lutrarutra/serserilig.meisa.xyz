package repository

import "github.com/iMeisa/serserilig.meisa.xyz/internal/models"

type DatabaseRepo interface {
	CreateTeamTable() error
	InsertDriver(driver models.Driver) error
	QueryAllDrivers() ([]models.Driver, error)
	QueryAllTeams() ([]models.Team, error)
}

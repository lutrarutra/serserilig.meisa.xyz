package repository

import "github.com/iMeisa/serserilig.meisa.xyz/internal/models"

type DatabaseRepo interface {
	CreateDriverTable() error
	CreateTeamTable() error
	InsertDriver(driver models.Driver) error
	QueryDriver(colName, value string) (models.Driver, error)
	QueryAllDrivers() ([]models.Driver, error)
	QueryAllTeams() ([]models.Team, error)
}

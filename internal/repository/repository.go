package repository

import "github.com/iMeisa/serserilig.meisa.xyz/internal/models"

type DatabaseRepo interface {
	CreateDriverTable() error
	CreateTeamTable() error
	InsertDriver(driver models.Driver) error
	DeleteDriver(driverId string) error
	UpdateDriver(idCol, idVal, updateCol, updateVal string) error
	QueryDriver(colName, value string) (models.Driver, error)
	QueryAllDrivers() ([]models.Driver, error)
	UpdateTeam(idCol, idVal, updateCol, updateVal string) error
	QueryAllTeams() ([]models.Team, error)
}

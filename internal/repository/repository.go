package repository

import "github.com/iMeisa/serserilig.meisa.xyz/internal/models"

type DatabaseRepo interface {
	CreateDriverTable() error
	CreateTeamTable() error
	CreateCalendarTable() error
	InsertDriver(driver models.Driver) error
	DeleteDriver(driverId string) error
	UpdateDriver(idCol, idVal, updateCol, updateVal string) error
	InsertRace(race models.Race) error
	DeleteRace(raceId string) error
	UpdateRace(id, date, time string) error
	QueryDriver(colName, value string) (models.Driver, error)
	QueryAllDrivers() ([]models.Driver, error)
	QueryAllRaces() ([]models.Race, error)
	UpdateTeam(idCol, idVal, updateCol, updateVal string) error
	QueryAllTeams() ([]models.Team, error)
}

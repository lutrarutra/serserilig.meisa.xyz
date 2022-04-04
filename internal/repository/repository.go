package repository

import "github.com/iMeisa/serserilig.meisa.xyz/internal/models"

type DatabaseRepo interface {
	InsertDriver(driver models.Driver) error
	QueryAllDrivers() ([]models.Driver, error)
}

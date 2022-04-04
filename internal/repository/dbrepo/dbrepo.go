package dbrepo

import (
	"database/sql"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/repository"
)

type sqliteDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewSqliteRepo(conn *sql.DB, app *config.AppConfig) repository.DatabaseRepo {
	return &sqliteDBRepo{
		App: app,
		DB: conn,
	}
}

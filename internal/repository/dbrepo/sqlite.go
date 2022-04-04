package dbrepo

import (
	"context"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"time"
)

func (m *sqliteDBRepo) createDriverTable() error {
	statement := `create table if not exists drivers (id integer primary key not null, name TEXT, team_id integer, points integer)`

	_, err := m.DB.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) InsertDriver(driver models.Driver) error {
	err := m.createDriverTable()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into drivers (name, points) values ($1, $2)`

	_, err = m.DB.ExecContext(ctx, statement, driver.Name, 0)
	if err != nil {
		return err
	}


	return nil
}

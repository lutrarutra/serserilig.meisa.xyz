package dbrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"io/ioutil"
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

func (m *sqliteDBRepo) createTeamTable() error {
	statement := `select name from sqlite_master where type='table' and name='teams'`

	rows, err := m.DB.Query(statement)
	if err != nil {
		return err
	}

	if rows.Next() {
		return nil
	}
	defer rows.Close()

	statement = `create table teams (id integer primary key not null, name text, abbr text, points integer, driver1 text, driver2 text)`

	_, err = m.DB.Exec(statement)
	if err != nil {
		return err
	}

	teams := map[string]models.Team{}

	teamsFile, err := ioutil.ReadFile("./static/json/teams.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(teamsFile, &teams)
	if err != nil {
		return err
	}

	statement = `insert into teams (name, abbr, points) values ($1, $2, $3)`
	for _, team := range teams {
		_, err = m.DB.Exec(statement, team.Name, team.Abbreviation, team.Points)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *sqliteDBRepo) InsertDriver(driver models.Driver) error {
	fmt.Println(m.DB.Stats())

	//err := m.createDriverTable()
	//if err != nil {
	//	return err
	//}
	//err = m.createTeamTable()
	//if err != nil {
	//	return err
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`select name from drivers where name='%s'`, driver.Name)

	rows, err := m.DB.Query(statement)
	if err != nil {
		return err
	}

	for rows.Next() {
		return nil
	}
	rows.Close()

	statement = `insert into drivers (name, points) values ($1, $2)`

	_, err = m.DB.ExecContext(ctx, statement, driver.Name, 0)
	if err != nil {
		return err
	}

	return nil
}

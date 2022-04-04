package dbrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"io/ioutil"
	"log"
	"time"
)

func (m *sqliteDBRepo) CreateDriverTable() error {
	statement := `create table if not exists drivers (id integer primary key not null, name TEXT, team_id integer, points integer)`

	_, err := m.DB.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) CreateTeamTable() error {
	statement := `select name from sqlite_master where type='table' and name='teams'`

	rows, err := m.DB.Query(statement)
	if err != nil {
		return err
	}

	if rows.Next() {
		return nil
	}

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

	statement = `insert into teams (name, abbr, points, driver1, driver2) values ($1, $2, $3, $4, $5)`
	for _, team := range teams {
		_, err = m.DB.Exec(statement, team.Name, team.Abbreviation, team.Points, -1, -1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *sqliteDBRepo) InsertDriver(driver models.Driver) error {
	fmt.Println(m.DB.Stats())

	err := m.CreateDriverTable()
	if err != nil {
		return err
	}

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

	statement = `insert into drivers (name, team_id, points) values ($1, $2, $3)`

	_, err = m.DB.ExecContext(ctx, statement, driver.Name, driver.TeamID, 0)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) QueryAllDrivers() ([]models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from drivers`

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var drivers []models.Driver

	for rows.Next() {
		var newDriver models.Driver
		err = rows.Scan(&newDriver.ID, &newDriver.Name, &newDriver.TeamID, &newDriver.Points)
		if err != nil {
			log.Println(err)
			continue
		}

		drivers = append(drivers, newDriver)
	}

	return drivers, nil
}

func (m *sqliteDBRepo) QueryAllTeams() ([]models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from teams order by points desc`

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var teams []models.Team

	for rows.Next() {
		var newTeam models.Team
		err = rows.Scan(&newTeam.ID, &newTeam.Name, &newTeam.Abbreviation, &newTeam.Points, &newTeam.Driver1, &newTeam.Driver2)
		if err != nil {
			log.Println(err)
			continue
		}

		teams = append(teams, newTeam)
	}

	return teams, nil
}

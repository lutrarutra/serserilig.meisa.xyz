package dbrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
)

func (m *sqliteDBRepo) CreateDriverTable() error {
	statement := `create table if not exists drivers (id integer primary key not null, name TEXT, team_id integer, points integer, penalty_points integer)`

	_, err := m.DB.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) CreateCalendarTable() error {
	statement := `create table if not exists races (id integer primary key not null, gp_id integer, race_date TEXT, race_time TEXT)`

	_, err := m.DB.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) CreateStaffTable() error {
	statement := `create table if not exists staff (id integer primary key not null, name TEXT, role_id integer, title TEXT, image_path TEXT)`

	_, err := m.DB.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) InsertRace(race models.Race) error {
	err := m.CreateCalendarTable()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// statement := fmt.Sprintf(`select name from races where ID='%d'`, race.ID)

	// rows, err := m.DB.QueryContext(ctx, statement)
	// if err != nil {
	// 	return err
	// }

	// for rows.Next() {
	// 	return nil
	// }

	statement := `insert into races (gp_id, race_date, race_time) values ($1, $2, $3)`
	_, err = m.DB.ExecContext(ctx, statement, race.Gp.ID, race.Date, race.Time)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) DeleteRace(raceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`delete from races where id=%v`, raceId)

	_, err := m.DB.ExecContext(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) UpdateRace(id, date, race_time string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// statement := fmt.Sprintf(`update races set race_date=%v where ID=%v`, date, id)

	statement := `update races set race_date=$1, race_time=$2 where ID=$3`
	_, err := m.DB.ExecContext(ctx, statement, date, race_time, id)

	if err != nil {
		log.Println("Updating race failed")
		return err
	}

	// statement = fmt.Sprintf(`update races set race_time=%v where ID=%v`, race_time, id)

	// _, err = m.DB.ExecContext(ctx, statement)
	// if err != nil {
	// 	log.Println("Error updating race time")
	// 	return err
	// }

	return nil
}

func (m *sqliteDBRepo) QueryStaff(id string) (models.Staff, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var staff models.Staff

	statement := fmt.Sprintf(`select * from staff where ID='%v'`, id)

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return staff, err
	}
	//id integer primary key not null, name TEXT, role_id integer, title TEXT, image_path TEXT

	var roles []models.Role

	rawFile, err := ioutil.ReadFile("./static/json/role_reference.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &roles)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var ID int
		var name string
		var role_id int
		var title string
		var image_path string

		err = rows.Scan(&ID, &name, &role_id, &title, &image_path)

		staff = models.Staff{
			ID:        ID,
			Name:      name,
			Role:      roles[role_id-1],
			Title:     title,
			ImagePath: image_path,
		}

		if err != nil {
			log.Println(err)
			continue
		} else {
			return staff, nil
		}
	}

	return staff, nil
}

func (m *sqliteDBRepo) QueryAllStaff() ([]models.Staff, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from staff`

	rows, err := m.DB.QueryContext(ctx, statement)

	if err != nil {
		return nil, err
	}

	var staff []models.Staff
	var roles []models.Role

	rawFile, err := ioutil.ReadFile("./static/json/role_reference.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &roles)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var ID int
		var name string
		var role_id int
		var title string
		var image_path string

		err = rows.Scan(&ID, &name, &role_id, &title, &image_path)

		newStaff := models.Staff{
			ID:        ID,
			Name:      name,
			Role:      roles[role_id-1],
			Title:     title,
			ImagePath: image_path,
		}

		if err != nil {
			log.Println(err)
			continue
		}

		staff = append(staff, newStaff)
	}

	return staff, nil
}

func (m *sqliteDBRepo) InsertStaff(staff models.Staff) error {
	err := m.CreateStaffTable()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// (id integer primary key not null, name TEXT, role_id integer, title TEXT, image_path TEXT)`

	statement := `insert into staff (name, role_id, title, image_path) values ($1, $2, $3, $4)`

	_, err = m.DB.ExecContext(ctx, statement, staff.Name, staff.Role.ID, staff.Title, staff.ImagePath)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) UpdateStaff(id, name, title, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//  id integer primary key not null, name TEXT, role_id integer, title TEXT, image_path TEXT
	// `update races set race_date=$1, race_time=$2 where ID=$3`
	role_id, _ := strconv.Atoi(role)
	statement := `update staff set name=$1, role_id=$2, title=$3 where ID=$4`
	_, err := m.DB.ExecContext(ctx, statement, name, role_id, title, id)

	if err != nil {
		log.Println("Updating staff failed")
		return err
	}

	return nil
}

func (m *sqliteDBRepo) DeleteStaff(staffId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`delete from staff where id=%v`, staffId)

	_, err := m.DB.ExecContext(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) QueryAllRaces() ([]models.Race, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from races order by race_date asc`

	rows, err := m.DB.QueryContext(ctx, statement)

	if err != nil {
		return nil, err
	}

	var races []models.Race
	var gps []models.GrandPrix

	rawFile, err := ioutil.ReadFile("./static/json/gp_reference.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &gps)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var gp_id int
		var date string
		var time string
		var ID int

		err = rows.Scan(&ID, &gp_id, &date, &time)

		newRace := models.Race{
			ID:   ID,
			Date: date,
			Time: time,
			Gp:   gps[gp_id-1],
		}

		if err != nil {
			log.Println(err)
			continue
		}

		races = append(races, newRace)
	}

	return races, nil
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

	statement = `create table teams (id integer primary key not null, name text, abbr text, points integer, 
										driver1 text, driver2 text, color text)`

	_, err = m.DB.Exec(statement)
	if err != nil {
		return err
	}

	teams := map[string]models.Team{}

	teamsFile, err := ioutil.ReadFile("./static/json/teams_reference.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(teamsFile, &teams)
	if err != nil {
		return err
	}

	statement = `insert into teams (name, abbr, points, driver1, driver2, color) values ($1, $2, $3, $4, $5, $6)`
	for _, team := range teams {
		_, err = m.DB.Exec(statement, team.Name, team.Abbreviation, team.Points, -1, -1, team.Color)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *sqliteDBRepo) InsertDriver(driver models.Driver) error {
	err := m.CreateDriverTable()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`select name from drivers where name='%s'`, driver.Name)

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return err
	}

	for rows.Next() {
		return nil
	}

	statement = `insert into drivers (name, team_id, points, penalty_points) values ($1, $2, $3, $4)`

	_, err = m.DB.ExecContext(ctx, statement, driver.Name, -1, 0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) DeleteDriver(driverId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`delete from drivers where id=%v`, driverId)

	_, err := m.DB.ExecContext(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) UpdateDriver(idCol, idVal, updateCol, updateVal string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`update drivers set %v=%v where %v=%v`, updateCol, updateVal, idCol, idVal)

	_, err := m.DB.ExecContext(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) QueryDriver(colName, value string) (models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var driver models.Driver

	statement := fmt.Sprintf(`select * from drivers where %v='%v'`, colName, value)

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return driver, err
	}

	for rows.Next() {
		err = rows.Scan(&driver.ID, &driver.Name, &driver.TeamID, &driver.Points, &driver.PenaltyPoints)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return driver, nil
}

func (m *sqliteDBRepo) QueryAllDrivers() ([]models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from drivers order by points desc, name asc`

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var drivers []models.Driver

	for rows.Next() {
		var newDriver models.Driver
		err = rows.Scan(&newDriver.ID, &newDriver.Name, &newDriver.TeamID, &newDriver.Points, &newDriver.PenaltyPoints)
		if err != nil {
			log.Println(err)
			continue
		}

		drivers = append(drivers, newDriver)
	}

	return drivers, nil
}

func (m *sqliteDBRepo) UpdateTeam(idCol, idVal, updateCol, updateVal string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := fmt.Sprintf(`update teams set %v=%v where %v=%v`, updateCol, updateVal, idCol, idVal)

	_, err := m.DB.ExecContext(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) QueryAllTeams() ([]models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select * from teams order by points desc, name asc`

	rows, err := m.DB.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var teams []models.Team

	for rows.Next() {
		var newTeam models.Team
		err = rows.Scan(&newTeam.ID, &newTeam.Name, &newTeam.Abbreviation, &newTeam.Points, &newTeam.Driver1, &newTeam.Driver2, &newTeam.Color)
		if err != nil {
			log.Println(err)
			continue
		}

		teams = append(teams, newTeam)
	}

	return teams, nil
}

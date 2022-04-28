package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Drivers   []Driver
	Teams     []Team
	Races     []Race
	GPs       []GrandPrix
	Roles     []Role
	Staff     []Staff
	CSRFToken string
	Flash     string // Some message
	Warning   string
	Error     string
}

func (data *TemplateData) GetRaces() {
	var races []Race
	rawFile, err := ioutil.ReadFile("./static/json/calendar.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &races)
	if err != nil {
		log.Println(err)
	}

	sort.Slice(races, func(i, j int) bool {
		return races[i].Date < races[j].Date
	})

	data.Races = races
}

func (data *TemplateData) GetRoles() {
	var roles []Role
	rawFile, err := ioutil.ReadFile("./static/json/role_reference.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &roles)
	if err != nil {
		log.Println(err)
	}

	data.Roles = roles
}

func (data *TemplateData) GetStaff() {
	var staff []Staff
	rawFile, err := ioutil.ReadFile("./static/json/staff.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &staff)
	if err != nil {
		log.Println(err)
	}

	data.Staff = staff
}

func (data *TemplateData) GetGrandPrixes() {
	var gps []GrandPrix
	rawFile, err := ioutil.ReadFile("./static/json/gp_reference.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(rawFile, &gps)
	if err != nil {
		log.Println(err)
	}

	sort.Slice(gps, func(i, j int) bool {
		return gps[i].Country < gps[j].Country
	})

	data.GPs = gps
}

func (data *TemplateData) GetDrivers() {
	var drivers []Driver
	rawFile, err := ioutil.ReadFile("./static/json/drivers.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(rawFile, &drivers)
	if err != nil {
		log.Println(err)
	}

	noDriver := Driver{
		ID:            -1,
		Name:          "No Driver",
		PenaltyPoints: 0,
	}
	drivers = append(drivers, noDriver)

	data.Drivers = drivers
}

func (data *TemplateData) GetTeams() {
	var teams []Team
	rawFile, err := ioutil.ReadFile("./static/json/teams.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(rawFile, &teams)
	if err != nil {
		log.Println(err)
	}

	data.Teams = teams
}

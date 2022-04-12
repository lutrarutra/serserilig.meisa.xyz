package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	Drivers    []Driver
	Teams      []Team
	CSRFToken  string
	Flash      string // Some message
	Warning    string
	Error      string
}

func(data *TemplateData) GetDrivers() {
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
		ID: -1,
		Name: "No Driver",
		PenaltyPoints: 0,
	}
	drivers = append(drivers, noDriver)
	
	data.Drivers = drivers
}

func(data *TemplateData) GetTeams() {
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

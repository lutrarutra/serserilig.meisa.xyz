package models

type Driver struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team-id"`
	Points int    `json:"points"`
}

type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abr"`
	Driver1      int    `json:"driver1"`
	Driver2      int    `json:"driver2"`
	Points       int    `json:"points"`
}

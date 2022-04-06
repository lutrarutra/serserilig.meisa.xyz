package models

type Driver struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team-id"`
	Points int    `json:"points"`
}

type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbr"`
	Driver1      int    `json:"driver1"`
	Driver2      int    `json:"driver2"`
	Points       int    `json:"points"`
	Color        string `json:"color"`
}

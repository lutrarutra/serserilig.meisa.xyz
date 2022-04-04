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
	Drivers      []int  `json:"driver-ids"`
	Points       int    `json:"points"`
}

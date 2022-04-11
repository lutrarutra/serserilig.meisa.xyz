package models

type Driver struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	TeamID        int    `json:"team-id"`
	Points        int    `json:"points"`
	PenaltyPoints int    `json:"penalty-points"`
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

type RaceResult struct {
	ID       int    `json:"id"`
	RaceID   int    `json:"race-id"`
	Pos      int    `json:"pos"`
	DriverID int    `json:"driver-id"`
	TeamID   string `json:"team-id"`
	DNS      bool   `json:"dns"`
	DNF      bool   `json:"dnf"`
}

type Race struct {
	ID       int          `json:"id"`
	Country  string       `json:"country"`
	FlagName string       `json:"flag-name"`
	Circuit  string       `json:"circuit"`
	Results  []RaceResult `json:"results"`
}

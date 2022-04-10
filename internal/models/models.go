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

type DriverResult struct {
	Pos    int    `json:"pos"`
	Driver Driver `json:"driver"`
	TeamID string `json:"team-id"`
	DNS    bool   `json:"dns"`
	DNF    bool   `json:"dnf"`
}

type Race struct {
	ID       int            `json:"id"`
	Country  string         `json:"country"`
	FlagName string         `json:"flag-name"`
	Circuit  string         `json:"circuit"`
	Results  []DriverResult `json:"results"`
}

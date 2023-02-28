package userplants

import "time"

type UserPlants struct {
	ID            int
	TagName       string
	Date          string
	WateringState bool
	DryState      bool
	HumidState    bool
	PlantID       string
	UserID        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

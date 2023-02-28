package plant

import (
	"time"
)

type Plant struct {
	ID               int
	Name             string
	Image            string
	Description      string
	Temperature      string
	WateringDuration string
	Soil             string
	Light            string
	Humidity         string
	Rainfall         string
	Tutorial         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}


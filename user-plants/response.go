package userplants

type UserPlantsResponse struct {
	ID            int       `json:"id"`
	TagName       string    `json:"tag_name"`
	WateringDate  string    `json:"watering_date"`
	MoveDate      string    `json:"move_date"`
	WateringState bool      `json:"watering_state"`
	DryState      bool      `json:"dry_state"`
	HumidState    bool      `json:"humid_state"`
	Plant         PlantItem `json:"plant"`
	User          UserItem  `json:"user"`
}

type PlantItem struct {
	PlantID          int    `json:"plant_id"`
	MinTemp          string `json:"min_temperature"`
	MaxTemp          string `json:"max_temperature"`
	Image            string `json:"image"`
	PlantName        string `json:"plant_name"`
	WateringDuration string `json:"watering_duration"`
}

type UserItem struct {
	UserID int    `json:"user_id"`
	Name   string `json:"user_name"`
}

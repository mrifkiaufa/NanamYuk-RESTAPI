package userplants

type UserPlantsRequestCreate struct {
	TagName       string `json:"tag_name" binding:"required"`
	WateringDate  string `json:"watering_date" binding:"required"`
	MoveDate      string `json:"move_date" binding:"required"`
	WateringState bool   `json:"watering_state"`
	DryState      bool   `json:"dry_state"`
	HumidState    bool   `json:"humid_state"`
	PlantID       string `json:"plant_id" binding:"required"`
	UserID        string `json:"user_id" binding:"required"`
}

type UserPlantsRequestUpdate struct {
	TagName       string `json:"tag_name"`
	WateringDate  string `json:"watering_date"`
	MoveDate      string `json:"move_date"`
	WateringState bool   `json:"watering_state"`
	DryState      bool   `json:"dry_state"`
	HumidState    bool   `json:"humid_state"`
	PlantID       string `json:"plant_id"`
	UserID        string `json:"user_id"`
}
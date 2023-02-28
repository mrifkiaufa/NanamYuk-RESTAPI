package userplants

type UserPlantsRequestCreate struct {
	TagName       string `json:"tag_name" binding:"required"`
	Date          string `json:"date" binding:"required"`
	WateringState bool   `json:"watering_state"`
	DryState      bool   `json:"dry_state"`
	HumidState    bool   `json:"humid_state"`
	PlantID       string `json:"plant_id" binding:"required"`
	UserID        string `json:"user_id" binding:"required"`
}

type UserPlantsRequestUpdate struct {
	TagName       string `json:"tag_name"`
	Date          string `json:"date"`
	WateringState bool   `json:"watering_state"`
	DryState      bool   `json:"dry_state"`
	HumidState    bool   `json:"humid_state"`
	PlantID       string `json:"plant_id"`
	UserID        string `json:"user_id"`
}
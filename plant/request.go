package plant

type PlantRequestCreate struct {
	Name             string `json:"name" binding:"required"`
	Image            string `json:"image" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Temperature      string `json:"temperature" binding:"required"`
	WateringDuration string `json:"watering_duration" binding:"required"`
	Soil             string `json:"soil" binding:"required"`
	Light            string `json:"light" binding:"required"`
	Humidity         string `json:"humidity" binding:"required"`
	Rainfall         string `json:"rainfall" binding:"required"`
	Tutorial         string `json:"tutorial" binding:"required"`
}

type PlantRequestUpdate struct {
	Name             string `json:"name"`
	Image            string `json:"image"`
	Description      string `json:"description"`
	Temperature      string `json:"temperature"`
	WateringDuration string `json:"watering_duration"`
	Soil             string `json:"soil"`
	Light            string `json:"light"`
	Humidity         string `json:"humidity"`
	Rainfall         string `json:"rainfall"`
	Tutorial         string `json:"tutorial"`
}
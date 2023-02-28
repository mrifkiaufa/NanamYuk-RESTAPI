package plant

type PlantResponse struct {
	Name             string `json:"name"`
	Image            string `json:"image"`
	ID               int    `json:"id"`
	Description      string `json:"description"`
	Temperature      string `json:"temperature"`
	WateringDuration string `json:"watering_duration"`
	Soil             string `json:"soil"`
	Light            string `json:"light"`
	Humidity         string `json:"humidity"`
	Rainfall         string `json:"rainfall"`
	Tutorial         string `json:"tutorial"`
}
package session

type SessionResponse struct {
	ID           int    `json:"id"`
	Date         string `json:"date"`
	UserPlantsID string `json:"user_plants_id"`
}
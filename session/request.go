package session

type SessionRequestCreate struct {
	Date         string `json:"date" binding:"required"`
	UserPlantsID string `json:"user_plants_id" binding:"required"`
}

type SessionRequestUpdate struct {
	Date         string `json:"date"`
	UserPlantsID string `json:"user_plants_id"`
}
package session

import (
	"time"
)

type Session struct {
	ID               int
	Date             string
	UserPlantsID     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}


package initializers

import (
	"nanam-yuk/auth"
	"nanam-yuk/plant"
	"nanam-yuk/session"
	userplants "nanam-yuk/user-plants"
)

func SyncDatabase() {
	DB.AutoMigrate(&plant.Plant{})
	DB.AutoMigrate(&userplants.UserPlants{})
	DB.AutoMigrate(&auth.Auth{})
	DB.AutoMigrate(&session.Session{})
}
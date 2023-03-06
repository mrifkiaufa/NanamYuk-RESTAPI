package main

import (
	"nanam-yuk/handler"
	"nanam-yuk/initializers"
	"nanam-yuk/plant"
	userplants "nanam-yuk/user-plants"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	plantRepository := plant.NewRepository(initializers.DB)
	plantService := plant.NewService(plantRepository)
	plantHandler := handler.NewPlantHandler(plantService)

	userPlantsRepository := userplants.NewRepository(initializers.DB)
	userPlantsService := userplants.NewService(userPlantsRepository)
	userPlantsHandler := handler.NewUserPlantsHandler(userPlantsService)

	//Membuat router untuk Endpoint
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if gin.Mode() == gin.ReleaseMode {
		router.Use(gin.Recovery())
	} else {
		router.Use(gin.Logger(), gin.Recovery())
	}

	plant := router.Group("/Plant")
	userPlants := router.Group("/UserPlants")
	auth := router.Group("/auth")

	plant.GET("/", plantHandler.GetPlants)
	plant.GET("/:id", plantHandler.GetPlant)
	plant.POST("/", plantHandler.CreatePlant)
	plant.PATCH("/:id", plantHandler.UpdatePlant)
	plant.DELETE("/:id", plantHandler.DeletePlant)

	userPlants.GET("/", userPlantsHandler.GetUserPlants)
	userPlants.POST("/", userPlantsHandler.CreatePlant)
	userPlants.PATCH("/:id", userPlantsHandler.UpdatePlant)
	userPlants.DELETE("/:id", userPlantsHandler.DeletePlant)

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
	auth.POST("/logout", handler.Logout)

	router.Run(":5000")
}

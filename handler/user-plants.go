package handler

import (
	"fmt"
	"nanam-yuk/auth"
	"nanam-yuk/initializers"
	"nanam-yuk/plant"
	userplants "nanam-yuk/user-plants"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userPlantsHandler struct {
	userPlantsService userplants.Service
}

func NewUserPlantsHandler(userPlantsService userplants.Service) *userPlantsHandler {
	return &userPlantsHandler{userPlantsService}
}

func (h *userPlantsHandler) GetUserPlants(c *gin.Context) {
	userPlants, err := h.userPlantsService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})
		return
	}

	var userPlantsResponse []userplants.UserPlantsResponse
	
	for _, p := range userPlants {
		
		userPlantResponse := convertToUserPlantResponse(c, p, p.UserID, p.PlantID)

		userPlantsResponse = append(userPlantsResponse, userPlantResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": userPlantsResponse,
	})
}

func (h *userPlantsHandler) GetUserPlant(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	p, err := h.userPlantsService.FindByID(int(id))

	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})
		return
	}

	userPlantResponse := convertToUserPlantResponse(c, p, p.UserID, p.PlantID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": userPlantResponse,
	})
}

func (h *userPlantsHandler) CreatePlant(c *gin.Context) {
	var userPlantRequest userplants.UserPlantsRequestCreate

	err := c.ShouldBindJSON(&userPlantRequest)
	if err != nil {

		errorMessages := []string{} 
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON((http.StatusBadRequest), gin.H{
			"status": http.StatusBadRequest,
			"message": errorMessages,
		})
		return
	}

	userPlant, err := h.userPlantsService.Create(userPlantRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})

		return
	}

	userPlantResponse := convertToUserPlantResponse(c, userPlant, userPlant.UserID, userPlant.PlantID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"response": userPlantResponse,
	})
}

func (h *userPlantsHandler) UpdatePlant(c *gin.Context) {
	var userPlantRequest userplants.UserPlantsRequestUpdate

	err := c.ShouldBindJSON(&userPlantRequest)
	if err != nil {

		errorMessages := []string{} 
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON((http.StatusBadRequest), gin.H{
			"status": http.StatusBadRequest,
			"message": errorMessages,
		})
		return
	}

	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	userPlant, err := h.userPlantsService.Update(id, userPlantRequest)

	if userPlant.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})

		return
	}

	userPlantResponse := convertToUserPlantResponse(c, userPlant, userPlant.UserID, userPlant.PlantID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": userPlantResponse,
	})
}

func (h *userPlantsHandler) DeletePlant(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	userPlant, err := h.userPlantsService.Delete(int(id))

	if userPlant.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Data deleted successfully",
	})
}

func convertToUserPlantResponse(c *gin.Context, p userplants.UserPlants, userID string, plantID string) userplants.UserPlantsResponse {
	var user auth.Auth
	initializers.DB.First(&user, "ID = ?", p.UserID)

	var plant plant.Plant
	initializers.DB.First(&plant, "ID = ?", p.PlantID)
	
	return userplants.UserPlantsResponse{
		ID: p.ID,
		TagName: p.TagName,
		Date: p.Date,
		WateringState: p.WateringState,
		DryState: p.DryState,
		HumidState: p.HumidState,
		Plant: userplants.PlantItem{
			PlantID: plant.ID,
			Image: plant.Image,
			PlantName: plant.Name,
			WateringDuration: plant.WateringDuration,
		},
		User: userplants.UserItem{
			UserID: user.ID,
			Name: user.Name,
		},
	}
}
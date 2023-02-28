package handler

import (
	"fmt"
	"nanam-yuk/plant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type plantHandler struct {
	plantService plant.Service
}

func NewPlantHandler(plantService plant.Service) *plantHandler {
	return &plantHandler{plantService}
}

func (h *plantHandler) GetPlants(c *gin.Context) {
	plants, err := h.plantService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})
		return
	}

	var plantsResponse []plant.PlantResponse
	
	for _, p := range plants {
		plantResponse := convertToPlantResponse(p)
	
		plantsResponse = append(plantsResponse, plantResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": plantsResponse,
	})
}

func (h *plantHandler) GetPlant(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	p, err := h.plantService.FindByID(int(id))

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

	plantResponse := convertToPlantResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": plantResponse,
	})
}

func (h *plantHandler) CreatePlant(c *gin.Context) {
	var plantRequest plant.PlantRequestCreate

	err := c.ShouldBindJSON(&plantRequest)
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

	plant, err := h.plantService.Create(plantRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"response": plant,
	})
}

func (h *plantHandler) UpdatePlant(c *gin.Context) {
	var plantRequest plant.PlantRequestUpdate

	err := c.ShouldBindJSON(&plantRequest)
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

	plant, err := h.plantService.Update(id, plantRequest)

	if plant.ID == 0 {
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
		"response": plant,
	})
}

func (h *plantHandler) DeletePlant(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	plant, err := h.plantService.Delete(int(id))

	if plant.ID == 0 {
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

	plantResponse := convertToPlantResponse(plant)
	
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": plantResponse,
	})
}

func convertToPlantResponse(p plant.Plant) plant.PlantResponse {
	return plant.PlantResponse{
		ID: p.ID,
		Name: p.Name,
		Image: p.Image,
		Description: p.Description,
		Temperature: p.Temperature,
		WateringDuration: p.WateringDuration,
		Soil: p.Soil,
		Light: p.Light,
		Humidity: p.Humidity,
		Rainfall: p.Rainfall,
		Tutorial: p.Tutorial,
	}
}
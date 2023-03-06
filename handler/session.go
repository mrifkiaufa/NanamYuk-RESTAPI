package handler

import (
	"fmt"
	"nanam-yuk/session"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type sessionHandler struct {
	sessionService session.Service
}

func NewsessionHandler(sessionService session.Service) *sessionHandler {
	return &sessionHandler{sessionService}
}

func (h *sessionHandler) GetSessions(c *gin.Context) {
	sessions, err := h.sessionService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})
		return
	}

	var sessionsResponse []session.SessionResponse
	
	for _, p := range sessions {
		sessionResponse := convertTosessionResponse(p)
	
		sessionsResponse = append(sessionsResponse, sessionResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": sessionsResponse,
	})
}

func (h *sessionHandler) GetSession(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	p, err := h.sessionService.FindByID(int(id))

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

	sessionResponse := convertTosessionResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": sessionResponse,
	})
}

func (h *sessionHandler) CreateSession(c *gin.Context) {
	var sessionRequest session.SessionRequestCreate

	err := c.ShouldBindJSON(&sessionRequest)
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

	session, err := h.sessionService.Create(sessionRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"response": session,
	})
}

func (h *sessionHandler) UpdateSession(c *gin.Context) {
	var sessionRequest session.SessionRequestUpdate

	err := c.ShouldBindJSON(&sessionRequest)
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

	session, err := h.sessionService.Update(id, sessionRequest)

	if session.ID == 0 {
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
		"response": session,
	})
}

func (h *sessionHandler) DeleteSession(c *gin.Context) {
	idString := c.Param(("id"))
	id, _ := strconv.Atoi(idString)

	session, err := h.sessionService.Delete(int(id))

	if session.ID == 0 {
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

	sessionResponse := convertTosessionResponse(session)
	
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": sessionResponse,
	})
}

func convertTosessionResponse(p session.Session) session.SessionResponse {
	return session.SessionResponse{
		ID: p.ID,
		Date: p.Date,
		UserPlantsID: p.UserPlantsID,
	}
}
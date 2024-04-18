package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(ctx *gin.Context) {
	// get user id from middleware
	userId := ctx.GetInt64("UserId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't Fetch Event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't Register Event for Event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Registered"})

}
func cancelForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("UserId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Event ID"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Event ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cancelled"})

}

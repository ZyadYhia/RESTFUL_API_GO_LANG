package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	//context.Request.Body() // to get request body
	//context.HTML() // for response html
	// code parameter is http code

	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch events."})
		return
	}
	//context.JSON(http.StatusOK, gin.H{"message": "Hello", "status": 200}) // for response json
	ctx.JSON(http.StatusOK, events) // for response json

}
func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse event ID."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't find event ID."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
func createEvents(ctx *gin.Context) {
	userId := ctx.GetInt64("UserId")

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Request Data."})
		return
	}
	event.UserID = userId
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create event."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "success", "event": event})

}
func updateEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("UserId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse event ID."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't find event ID."})
		return
	}
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update Event."})
		return
	}
	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully"})
}
func deleteEvent(ctx *gin.Context) {

	userId := ctx.GetInt64("UserId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse event ID."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't find event ID."})
		return
	}
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to Delete Event."})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Can't delete the event")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully"})
}

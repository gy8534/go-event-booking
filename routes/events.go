package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gy8534/go-event-booking/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events, try again later",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}

	userID := ctx.GetInt64("userID")
	event.UserID = userID

	err := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create events, try again later",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event create",
		"event":   event,
	})
}

func getEventByID(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse eventID",
		})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func updateEventByID(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse eventID",
		})
		return
	}

	userID := ctx.GetInt64("userID")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}

	if event.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorised the update event",
		})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update the event",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "event updated successfully",
	})
}

func deleteEventByID(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse eventID",
		})
		return
	}

	userID := ctx.GetInt64("userID")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
		})
		return
	}

	if event.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorised the delete event",
		})
		return
	}

	err = event.DeleteEventByID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not delete the event",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "event deleted successfully",
	})
}

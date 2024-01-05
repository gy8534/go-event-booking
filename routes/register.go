package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gy8534/go-event-booking/models"
)

func registerForEvent(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event ID",
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

	err = event.Register(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not register user for event",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "registered successfully",
	})

}

func cancelRegistration(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event ID",
		})
		return
	}

	var event models.Event
	event.ID = eventID
	err = event.CancelRegistration(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not cancel registration",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "registration cancelled successfully",
	})
}

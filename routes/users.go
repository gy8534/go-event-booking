package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gy8534/go-event-booking/models"
	"github.com/gy8534/go-event-booking/utils"
)

func signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}

	err := user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not save the user data",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func login(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not authenticated user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})

}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gy8534/go-event-booking/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorised",
		})
		return
	}

	userID, err := utils.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorised",
		})
		return
	}

	ctx.Set("userID", userID)
	ctx.Next()
}

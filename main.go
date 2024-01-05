package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy8534/go-event-booking/db"
	"github.com/gy8534/go-event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		return
	}
}

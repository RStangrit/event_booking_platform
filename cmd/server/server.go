package server

import (
	events "main/internal/event"
	users "main/internal/user"

	"github.com/gin-gonic/gin"
)

func LaunchServer() {
	server := gin.Default()

	users.RegisterRoutes(server)
	events.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}

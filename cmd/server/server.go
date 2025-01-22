package server

import (
	users "main/internal/user"

	"github.com/gin-gonic/gin"
)

func LaunchServer() {
	server := gin.Default()

	users.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}

package server

import "github.com/gin-gonic/gin"

func LaunchServer() {
	server := gin.Default()
	server.Run(":8080") //localhost:8080
}

package event

import "github.com/gin-gonic/gin"

//for registering routes

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events")
	server.GET("/events")
	server.GET("/events/:id")
	server.PUT("/events/:id")
	server.DELETE("/events/:id")
}

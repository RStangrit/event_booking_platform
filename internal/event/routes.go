package event

import "github.com/gin-gonic/gin"

//for registering routes

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", createEventHandler)
	server.GET("/events", getEventsHandler)
	server.GET("/events/:id", getEventHandler)
	server.PUT("/events/:id", updateEventHandler)
	server.DELETE("/events/:id", deleteEventHandler)
}

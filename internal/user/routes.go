package users

//for registering routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/users", createUserHandler)
	server.GET("/users", getUsersHandler)
	server.GET("/users/:id", getUserHandler)
	server.PUT("/users/:id", updateUserHandler)
	server.DELETE("/users/:id", deleteUserHandler)
}

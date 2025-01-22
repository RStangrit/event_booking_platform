package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/users")
	server.POST("/users")
	server.GET("/users/:id")
	server.PUT("/users/:id")
	server.DELETE("/users/:id")
	// server.GET("/users", GetUsersHandler)
	// server.POST("/users", CreateUserHandler)
	// server.GET("/users/:id", GetUserHandler)
	// server.PUT("/users/:id", UpdateUserHandler)
	// server.DELETE("/users/:id", DeleteUserHandler)
}

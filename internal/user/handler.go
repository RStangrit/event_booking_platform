package users

import (
	"main/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to handle HTTP requests related to users.
func createUserHandler(context *gin.Context) {
	var user UserRequest

	err := context.ShouldBindJSON(&user)
	if err != nil {
		util.ProvideResponse(context, 400, "Invalid JSON format")
	}

	err = user.Save()
	if err != nil {
		util.ProvideResponse(context, 500, "Failed to save user")
		return
	}

	util.ProvideResponse(context, 201, "User created successfully")
}

func getUsersHandler(context *gin.Context) {
	users, err := getAllUsers()
	if err != nil {
		util.ProvideResponse(context, 500, "Could not fetch users")
		return
	}

	if len(users) == 0 {
		util.ProvideResponse(context, 404, "No users found")
		return
	}

	context.JSON(http.StatusOK, users)
}

func getUserHandler(context *gin.Context) {
	userId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "Invalid user ID format")
		return
	}

	user, err := getOneUser(userId)
	if err != nil {
		util.ProvideResponse(context, 404, "User not found")
		return
	}

	context.JSON(http.StatusOK, user)
}

func updateUserHandler(context *gin.Context) {
	userId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "Invalid user ID format")
		return
	}

	user, err := getOneUser(userId)
	if err != nil {
		util.ProvideResponse(context, 404, "User not found")
		return
	}

	var updatedUser UserResponse
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		util.ProvideResponse(context, 400, "Invalid request data format")
		return
	}

	updatedUser.Id = user.Id

	err = updatedUser.Update()
	if err != nil {
		util.ProvideResponse(context, 500, "Failed to update user")
		return
	}

	util.ProvideResponse(context, 200, "User updated successfully")
}

func deleteUserHandler(context *gin.Context) {
	userId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse user id")
		return
	}

	user, err := getOneUser(userId)
	if err != nil {
		util.ProvideResponse(context, 400, "Invalid user ID format")
		return
	}

	if user.Role == "admin" {
		util.ProvideResponse(context, 401, "Not authorized to delete an admin user")
		return
	}

	err = user.Delete()
	if err != nil {
		util.ProvideResponse(context, 400, "Failed to delete user")
		return
	}
	util.ProvideResponse(context, 200, "User deleted successfully")
}

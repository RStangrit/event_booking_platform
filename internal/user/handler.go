package users

import (
	"fmt"
	"main/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// to handle HTTP requests related to users.
func createUserHandler(context *gin.Context) {
	var user UserRequest

	err := context.ShouldBindJSON(&user)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse request data")
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		util.ProvideResponse(context, 500, "could not save user")
		return
	}

	util.ProvideResponse(context, 201, "user created successfully")
}

func getUsersHandler(context *gin.Context) {
	users, err := getAll()
	if err != nil {
		util.ProvideResponse(context, 500, "could not fetch users")
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse user id")
		return
	}

	user, err := getOne(userId)
	if err != nil {
		util.ProvideResponse(context, 400, "could not fetch user")
		return
	}
	context.JSON(http.StatusOK, user)
}

func updateUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse user id")
		return
	}

	user, err := getOne(userId)
	if err != nil {
		util.ProvideResponse(context, 500, "could not fetch user")
		return
	}

	var updatedUser UserResponse
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse request data")
		return
	}

	updatedUser.Id = user.Id
	err = updatedUser.Update()
	if err != nil {
		util.ProvideResponse(context, 500, "could not update user")
		return
	}
	util.ProvideResponse(context, 200, "user updated successfully")
}

func deleteUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse user id")
		return
	}

	user, err := getOne(userId)
	if err != nil {
		util.ProvideResponse(context, 500, "could not fetch user")
		return
	}

	if user.Role == "admin" {
		util.ProvideResponse(context, 401, "not authorized to delete user")
		return
	}

	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not delete user"})
		util.ProvideResponse(context, 400, "could not delete user")
		return
	}
	util.ProvideResponse(context, 200, "user deleted successfully")
}

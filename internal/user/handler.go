package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// to handle HTTP requests related to users.
func createUserHandler(context *gin.Context) {
	var user UserRequest

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func getUsersHandler(context *gin.Context) {
	users, err := getAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch users"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse user id"})
		return
	}

	user, err := getOne(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func updateUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse user id"})
		return
	}

	user, err := getOne(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}

	var updatedUser UserResponse
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedUser.Id = user.Id
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func deleteUserHandler(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse user id"})
		return
	}

	user, err := getOne(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}

	if user.Role == "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete user"})
		return
	}

	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not delete user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

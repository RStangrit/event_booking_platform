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

	unique, err := IsEmailUnique(user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sometthing wrong with database"})
	}

	if !unique {
		context.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func getUsersHandler(context *gin.Context) {
	users, err := getAll()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}
	context.JSON(http.StatusOK, user)
}

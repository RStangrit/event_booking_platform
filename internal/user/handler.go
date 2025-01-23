package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// to handle HTTP requests related to users.
func createUserHandler(context *gin.Context) {
	var user User

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

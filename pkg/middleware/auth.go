package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not find authentication token"})
		return
	}

	// userId, err := utils.VerifyToken(token)

	// if err != nil {
	// 	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not verify authentication token"})
	// }

	// context.Set("userId", userId)
	context.Next()
}

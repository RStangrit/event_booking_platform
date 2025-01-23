package util

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ProvideResponse(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
	})
	context.Abort()
}

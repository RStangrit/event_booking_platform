package util

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetIntParam(context *gin.Context, intParamName string) (int64, error) {
	intValue, err := strconv.ParseInt(context.Param(intParamName), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(intValue), nil

}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetEnvVariable(param string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	paramEnv := os.Getenv(param)
	if paramEnv == "" {
		panic("Error reading parameter")
	}
	return paramEnv
}

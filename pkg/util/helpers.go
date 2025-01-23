package util

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

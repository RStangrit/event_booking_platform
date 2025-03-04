package event

import (
	"fmt"
	"main/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to handle HTTP requests related to users.
func createEventHandler(context *gin.Context) {
	var event EventRequest

	err := context.ShouldBindJSON(&event)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse request data")
	}

	err = event.Save()
	if err != nil {
		util.ProvideResponse(context, 500, "could not save event")
		return
	}

	util.ProvideResponse(context, 201, "event created successfully")
}

func getEventsHandler(context *gin.Context) {
	events, err := getAll()
	if err != nil {
		fmt.Println(err)
		util.ProvideResponse(context, 500, "could not fetch events")
	}
	context.JSON(http.StatusOK, events)
}

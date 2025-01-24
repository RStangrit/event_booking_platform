package event

import (
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
	events, err := getAllEvents()
	if err != nil {
		util.ProvideResponse(context, 500, "could not fetch events")
	}
	context.JSON(http.StatusOK, events)
}

func getEventHandler(context *gin.Context) {
	eventId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse event id")
		return
	}

	event, err := getOneEvent(eventId)
	if err != nil {
		util.ProvideResponse(context, 400, "could not fetch event")
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEventHandler(context *gin.Context) {
	eventId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse event id")
		return
	}

	event, err := getOneEvent(eventId)
	if err != nil {
		util.ProvideResponse(context, 400, "could not fetch event")
		return
	}

	var updatedEvent EventResponse
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse request data")
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		util.ProvideResponse(context, 500, "could not update event")
		return
	}
	util.ProvideResponse(context, 200, "event updated successfully")
}

func deleteEventHandler(context *gin.Context) {
	eventId, err := util.GetIntParam(context, "id")
	if err != nil {
		util.ProvideResponse(context, 400, "could not parse event id")
		return
	}

	event, err := getOneEvent(eventId)
	if err != nil {
		util.ProvideResponse(context, 400, "could not fetch event")
		return
	}

	err = event.Delete()
	if err != nil {
		util.ProvideResponse(context, 400, "could not delete event")
	}
	util.ProvideResponse(context, 200, "event deleted successfully")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-api/models"
)

func main() {
	// setup the "Engine" (HTTP SERVER) with Logger and Recovery middleware
	// Recovery - recovers from crashes if they are not entire server crashes
	var server = gin.Default()

	// handlers registration
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

// context will be sent by Gin automatically if this function is registered as handler in server.GET()
func getEvents(context *gin.Context) {
	var events = models.GetAllEvents()

	// instead of returning anything from this function we have to use JSON method of the context
	context.JSON(
		http.StatusOK,
		events)
}

func createEvent(ctx *gin.Context) {
	var eventModel models.Event

	// map the json body to Event type and store it in eventModel variable
	// gin by default does not complain if any fields missing, it will mark them as nil
	// but we use `binding:required` tags on our properties to mark which of them are mandatory
	var err = ctx.ShouldBindJSON(&eventModel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Could not parse Event during Event Creation")
		// we must use return here otherwise the code below will be executed anyway despite the error we send back
		return
	}

	eventModel.Id = 1
	eventModel.UserId = 1

	eventModel.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created and stored",
		"event":   eventModel})
}

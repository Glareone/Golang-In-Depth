package routes

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-api/models"
)

// context will be sent by Gin automatically if this function is registered as handler in server.GET()
func getEvents(ctx *gin.Context) {
	var events, err = models.GetAllEvents()

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Events cannot be read, error",
				"error":   err.Error(),
			})
		return
	}

	// instead of returning anything from this function we have to use JSON method of the context
	ctx.JSON(
		http.StatusOK,
		events)
}

func getEventById(ctx *gin.Context) {
	// Get the "eventId" query parameter
	eventIdStr := ctx.Param("id")

	// Convert the ID to an integer
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Id could not be parsed",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "event with this Id is not found",
					"error":   err.Error(),
				})
			return
		} else {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "something went wrong",
					"error":   err.Error(),
				})
			return
		}
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var eventModel models.Event

	// map the json body to Event type and store it in eventModel variable
	// gin by default does not complain if any fields missing, it will mark them as nil
	// but we use `binding:required` tags on our properties to mark which of them are mandatory
	err := ctx.ShouldBindJSON(&eventModel)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse Event during Event Creation",
				"error":   err.Error(),
			})
		// we must use return here otherwise the code below will be executed anyway despite the error we send back
		return
	}

	eventModel.Id = 1
	eventModel.UserId = 1

	err = eventModel.Save()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "event cannot be saved",
				"error":   err.Error(),
			})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created and stored",
		"event":   eventModel})
}

func updateEvent(ctx *gin.Context) {
	// Get the "eventId" query parameter
	eventIdStr := ctx.Param("id")

	// Convert the ID to an integer
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Id could not be parsed",
				"error":   err.Error(),
			})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "event with this Id is not found",
					"error":   err.Error(),
				})
			return
		} else {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "something went wrong",
					"error":   err.Error(),
				})
			return
		}
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid data submitted in the body",
				"error":   err.Error(),
			})
		return
	}

	// set the ID of the event we get from the database
	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEvent()

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "sql update process failed",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(ctx *gin.Context) {
	// Get the "eventId" query parameter
	eventIdStr := ctx.Param("id")

	// Convert the ID to an integer
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Id could not be parsed",
				"error":   err.Error(),
			})
		return
	}

	err = models.DeleteEventTransactional(eventId)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Event deletion raised an error",
				"error":   err.Error(),
			})
		return
	}

	ctx.Status(http.StatusNoContent)
}

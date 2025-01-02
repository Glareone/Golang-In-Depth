package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// setup the "Engine" (HTTP SERVER) with Logger and Recovery middleware
	// Recovery - recovers from crashes if they are not entire server crashes
	var server = gin.Default()

	// handlers registration
	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

// context will be sent by Gin automatically if this function is registered as handler in server.GET()
func getEvents(context *gin.Context) {
	// instead of returning anything from this function we have to use JSON method of the context
	context.JSON(
		http.StatusOK,
		// gin H - shorthand for 'map' key-value structure.
		gin.H{"event": "hello world event!"})
}

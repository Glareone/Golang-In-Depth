package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"web-api/database"
	"web-api/routes"
)

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// setup the "Engine" (HTTP SERVER) with Logger and Recovery middleware
	// Recovery - recovers from crashes if they are not entire server crashes
	var server = gin.Default()

	// initialize Database Connection
	database.InitDatabase()

	// Close the connection when the application shuts down
	// we dont need to close connection pool each time
	// we also can manage connections manually calling defer database.DB.conn.Close(), but it's less common practice
	defer database.DB.Close()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

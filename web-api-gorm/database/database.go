package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"web-api/models"
)

var DB *gorm.DB

func InitDatabase() {
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_NAME")

	connectionStringPostgres := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionStringPostgres), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection with the database cannot be established")
	}

	// AutoMigrate will create the tables if they don't exist
	err = DB.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatalf("[Migration] Error during auto-migration: %w", err)
	}
}

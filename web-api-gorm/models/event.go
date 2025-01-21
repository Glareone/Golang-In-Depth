package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"web-api/database"
)

type Event struct {
	gorm.Model            // GORM's base model (includes ID, CreatedAt, UpdatedAt, DeletedAt)
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Location    string    `gorm:"not null"`
	DateTime    time.Time `gorm:"not null"`
	UserID      int
}

func (e *Event) Save() error {
	result := database.DB.Create(e) // Use GORM's Create method
	if result.Error != nil {
		return fmt.Errorf("error saving event: %w", result.Error)
	}
	return nil
}

func (e *Event) UpdateEvent() error {
	result := database.DB.Save(e)
	if result.Error != nil {
		return fmt.Errorf("event cannot be updated, error: %w", result.Error)
	}

	return nil
}

func DeleteEvent(eventId int64) error {
	// Tell GORM to delete an Event with the given ID
	result := database.DB.Delete(&Event{}, eventId)

	if result.Error != nil {
		return fmt.Errorf("error deleting event: %w", result.Error)
	}

	// Check if any rows were affected (optional)
	if result.RowsAffected == 0 {
		return fmt.Errorf("no event found with ID %d", eventId)
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	// Create a slice to hold the events
	var events []Event

	// Use GORM's Find method to retrieve all events
	result := database.DB.Find(&events)

	// Check for errors
	if result.Error != nil {
		return nil, fmt.Errorf("error getting events: %w", result.Error)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	var event Event

	result := database.DB.First(&event, eventId)

	// Check for errors, including "record not found"
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("event is not found: %w", result.Error)
		}
		return nil, fmt.Errorf("error getting event by ID: %w", result.Error)
	}

	return &event, nil
}

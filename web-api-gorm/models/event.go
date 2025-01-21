package models

import (
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

//func GetEventById(eventId int64) (Event, error) {
//	query := `SELECT *
//			FROM events e
//			WHERE e.Id = $1
//			LIMIT 1` // Use LIMIT 1 instead of TOP(1) in PostgreSQL
//
//	statement, err := database.DB.Prepare(query)
//
//	defer statement.Close()
//
//	if err != nil {
//		return Event{}, fmt.Errorf("query preparation failed raising the following error: %w", err)
//	}
//
//	var event Event
//	err = statement.QueryRowContext(context.Background(), eventId).Scan(
//		&event.Id, &event.Name,
//		&event.Description, &event.Location,
//		&event.DateTime, &event.UserId)
//
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return Event{}, fmt.Errorf("event not found: %w", err) // Specific error for not found
//		}
//		return Event{}, fmt.Errorf("query execution failed raising the error: %w", err)
//	}
//
//	return event, nil
//}

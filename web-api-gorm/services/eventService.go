package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"web-api/database"
	"web-api/domain"
	"web-api/models"
)

type EventService interface {
	GetEventById(eventId int64) (*domain.Event, error)
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(domainEvent *domain.Event) (*domain.Event, error)
	DeleteEventById(eventId int64) error
	UpdateEvent(domainEvent *domain.Event) error
}

type eventService struct{}

func NewEventService() EventService {
	return &eventService{}
}

func (s *eventService) CreateEvent(domainEvent *domain.Event) (*domain.Event, error) {
	modelsEvent := models.Event{
		Name:        domainEvent.Name,
		Description: domainEvent.Description,
		Location:    domainEvent.Location,
		DateTime:    domainEvent.DateTime,
		UserID:      domainEvent.UserID,
	}

	result := database.DB.Create(&modelsEvent) // Use GORM's Create method
	if result.Error != nil {
		return nil, fmt.Errorf("error saving event: %w", result.Error)
	}

	domainEvent.ID = int64(modelsEvent.ID)

	return domainEvent, nil
}

func (s *eventService) GetAllEvents() ([]domain.Event, error) {
	var modelEvents []models.Event
	var domainEvents []domain.Event

	// Use GORM's Find method to retrieve all events
	result := database.DB.Find(&modelEvents)

	// Check for errors
	if result.Error != nil {
		return nil, fmt.Errorf("error getting events: %w", result.Error)
	}

	// Convert database models to domain models
	for _, modelEvent := range modelEvents {
		domainEvents = append(domainEvents, domain.Event{
			ID:          int64(modelEvent.ID), // Convert uint to int if needed
			Name:        modelEvent.Name,
			Description: modelEvent.Description,
			Location:    modelEvent.Location,
			DateTime:    modelEvent.DateTime,
			UserID:      modelEvent.UserID,
		})
	}

	return domainEvents, nil
}

func (s *eventService) DeleteEventById(eventId int64) error {
	// Tell GORM to delete an Event with the given ID
	result := database.DB.Delete(&models.Event{}, eventId)

	if result.Error != nil {
		return fmt.Errorf("error deleting event: %w", result.Error)
	}

	// Check if any rows were affected (optional)
	if result.RowsAffected == 0 {
		return fmt.Errorf("no event found with ID %d", eventId)
	}

	return nil
}

func (s *eventService) UpdateEvent(domainEvent *domain.Event) error {
	modelEvent := models.Event{
		Model:       gorm.Model{ID: uint(domainEvent.ID)}, // Set the ID from domainEvent
		Name:        domainEvent.Name,
		Description: domainEvent.Description,
		Location:    domainEvent.Location,
		DateTime:    domainEvent.DateTime,
		UserID:      domainEvent.UserID,
	}

	result := database.DB.Save(&modelEvent)
	if result.Error != nil {
		return fmt.Errorf("event cannot be updated, error: %w", result.Error)
	}

	return nil
}

func (s *eventService) GetEventById(eventId int64) (*domain.Event, error) {
	var modelEvent models.Event

	result := database.DB.First(&modelEvent, eventId)

	// Check for errors, including "record not found"
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("event is not found: %w", result.Error)
		}
		return nil, fmt.Errorf("error getting event by ID: %w", result.Error)
	}

	domainEvent := domain.Event{
		ID:          int64(modelEvent.ID), // Convert uint to int if needed
		Name:        modelEvent.Name,
		Description: modelEvent.Description,
		Location:    modelEvent.Location,
		DateTime:    modelEvent.DateTime,
		UserID:      modelEvent.UserID,
	}

	return &domainEvent, nil
}

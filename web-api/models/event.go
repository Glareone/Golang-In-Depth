package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"web-api/database"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"` // we mark fields as required for gin
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int       // Id and UserId are not required obviously
}

var events = []Event{}

func (e *Event) Save() error {
	var query = `
        INSERT INTO events (name, description, location, dateTime, user_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING Id
    `

	// prepare the query statement
	// used to inject parameters
	// PERFORMANCE TIP:
	// Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed
	// multiple times (potentially with different data for its placeholders).
	// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions.
	// In that case, there wouldn't be any advantages.
	// PS recommended to use PrepareContext to handle cancellations
	var statement, err = database.DB.PrepareContext(context.Background(), query)

	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	// Defer closing the statement after execution is finished
	// The defer stmt.Close() statement is now immediately after the stmt is created.
	// This is the recommended practice because it ensures that the statement is closed as soon as the function exits,
	// regardless of how it exits (normally or due to an error).
	// defer Keyword: The defer keyword schedules a function call to be executed when the surrounding function returns.
	// LIFO Order: If you have multiple defer statements, they are executed in Last-In, First-Out (LIFO) order.
	// Guaranteed Execution: The defer statement ensures that the deferred function (stmt.Close() in this case) is called even if the function exits due to a panic or an error.
	defer statement.Close()

	// Execute the query and retrieve the ID
	// context.Background() to provide a context to the QueryRowContext() method.
	// This allows pgx to potentially handle cancellation or timeouts during the query execution,
	// even if you're not explicitly setting any deadlines or cancellation signals in this specific example.
	//
	// .Scan(&e.Id) - to setup Id to Event struct and to avoid using LastInsertId method (which is not supported by pgx driver)
	// you have to use QueryRow together with Scan in order to retrieve Id of inserted element back
	err = statement.QueryRowContext(context.Background(),
		e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
	if err != nil {
		return fmt.Errorf("error saving event: %w", err)
	}

	return nil
}

func (e *Event) UpdateEvent() error {
	query := `
		UPDATE events e
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5`

	statement, err := database.DB.PrepareContext(context.Background(), query)

	if err != nil {
		return fmt.Errorf("error prepare statement for updating the event failed: %w", err)
	}

	defer statement.Close()

	_, err = statement.ExecContext(context.Background(), e.Name, e.Description, e.Location, e.DateTime, e.Id)

	if err != nil {
		return fmt.Errorf("error updating the event: %w", err)
	}

	return nil
}

func DeleteEventTransactional(eventId int64) error {
	// start transaction with default level of isolation
	transaction, err := database.DB.Begin()
	if err != nil {
		return fmt.Errorf("transaction cannot be started: %w", err)
	}

	defer transaction.Rollback()

	query := `
		DELETE
		FROM events
		WHERE id = $1
	`

	result, err := transaction.ExecContext(context.Background(), query, eventId)

	if err != nil {
		return fmt.Errorf("error during deleting the event: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no event found with ID %d", eventId)
	}

	err = transaction.Commit()

	if err != nil {
		return fmt.Errorf("issue with committing the transaction %w", err)
	}

	return nil
}

// normal function, not a method of *Event
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// we still can use prepare here, but SELECT * is pretty simple sql call, so we do not prepare it
	// and we do not store it in pgx memory
	// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions.
	// In that case, there wouldn't be any advantages.

	eventRows, err := database.DB.QueryContext(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("error saving event: %w", err)
	}

	// prevent further enumeration
	defer eventRows.Close()

	for eventRows.Next() {
		var event Event
		err := eventRows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, fmt.Errorf("error scanning and populating event collection: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (Event, error) {
	query := `SELECT *
			FROM events e
			WHERE e.Id = $1
			LIMIT 1` // Use LIMIT 1 instead of TOP(1) in PostgreSQL

	statement, err := database.DB.Prepare(query)

	defer statement.Close()

	if err != nil {
		return Event{}, fmt.Errorf("query preparation failed raising the following error: %w", err)
	}

	var event Event
	err = statement.QueryRowContext(context.Background(), eventId).Scan(
		&event.Id, &event.Name,
		&event.Description, &event.Location,
		&event.DateTime, &event.UserId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, fmt.Errorf("event not found: %w", err) // Specific error for not found
		}
		return Event{}, fmt.Errorf("query execution failed raising the error: %w", err)
	}

	return event, nil
}

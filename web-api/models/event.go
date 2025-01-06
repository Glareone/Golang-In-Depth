package models

import (
	"context"
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
	var statement, err = database.DB.Prepare(query)

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
	err = statement.QueryRowContext(context.Background(),
		e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
	if err != nil {
		return fmt.Errorf("error saving event: %w", err)
	}

	return nil
}

// normal function, not a method of *Event
func GetAllEvents() []Event {
	return events
}

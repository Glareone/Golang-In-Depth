package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"` // we mark fields as required for gin
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int       // Id and UserId are not required obviously
}

var events = []Event{}

func (e *Event) Save() {
	// later: save to the database
	// for now we just save new Event e in our slice events
	// also use dereference to get the copy of the "e" by the address
	events = append(events, *e)
}

// normal function, not a method of *Event
func GetAllEvents() []Event {
	return events
}

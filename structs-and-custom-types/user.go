package main

import "time"

// making it upper case you make it available outside of this file
// making local types you can use lowerCase custom type names
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model            // GORM's base model (includes ID, CreatedAt, UpdatedAt, DeletedAt)
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Location    string    `gorm:"not null"`
	DateTime    time.Time `gorm:"not null"`
	UserID      int
}

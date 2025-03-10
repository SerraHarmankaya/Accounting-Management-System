package models

import (
	"time"

	"gorm.io/gorm"
)

// Ticket Model

type Ticket struct {
	gorm.Model
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status" gorm:"default:open"` // open, in-progress, closed
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
}

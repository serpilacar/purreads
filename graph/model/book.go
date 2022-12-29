package model

import (
	"time"
)

type Book struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

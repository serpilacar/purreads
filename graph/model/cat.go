package model

import "time"

type AddBookToCatInput struct {
	CatID     int  `json:"catId"`
	BookID    int  `json:"bookId"`
	Completed bool `json:"completed"`
}

type Cat struct {
	ID              int    `gorm:"primaryKey" json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `gorm:"unique" json:"email"`
	ReadBooks       []Book `gorm:"many2many:cat_read_books" json:"readBooks"`
	InProgressBooks []Book `gorm:"many2many:cat_in_progress_books" json:"inProgressBooks"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

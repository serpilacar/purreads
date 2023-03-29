package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"gorm.io/gorm"
)

type Resolver struct {
	// books []*model.Book
	// cats []*model.Cat
	DB *gorm.DB
}

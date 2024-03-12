package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"purreads/graph/model"

	"gorm.io/gorm/clause"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	book := &model.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	err := r.DB.Create(&book).Error
	return book, err
}

// CreateCat is the resolver for the createCat field.
func (r *mutationResolver) CreateCat(ctx context.Context, input model.NewCat) (*model.Cat, error) {
	cat := &model.Cat{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	err := r.DB.FirstOrCreate(&cat, &model.Cat{Email: input.Email}).Error
	return cat, err
}

// AddBookToCat is the resolver for the addBookToCat field.
func (r *mutationResolver) AddBookToCat(ctx context.Context, input model.AddBookToCatInput) (*model.Cat, error) {
	cat := &model.Cat{ID: input.CatID}
	book := &model.Book{ID: input.BookID}

	if err := r.DB.First(cat).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(book).Error; err != nil {
		return nil, err
	}

	if input.Completed {
		r.DB.Model(cat).Association("ReadBooks").Append(book)
		return cat, nil
	} else {
		r.DB.Model(cat).Association("InProgressBooks").Append(book)
		return cat, nil
	}
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	books := []*model.Book{}
	err := r.DB.Find(&books).Error
	return books, err
}

// Cats is the resolver for the cats field.
func (r *queryResolver) Cats(ctx context.Context, id *int) ([]*model.Cat, error) {
	if id != nil {
		cat := &model.Cat{ID: *id}
		err := r.DB.First(cat).Error
		return []*model.Cat{cat}, err
	}
	cats := []*model.Cat{}
	err := r.DB.Model(&model.Cat{}).Preload(clause.Associations).Find(&cats).Error
	return cats, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

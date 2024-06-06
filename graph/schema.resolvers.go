package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"Jibambe/engine/books"
	"Jibambe/graph/model"
	"context"
)

// CreateBook is the resolver for the CreateBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	return books.CreateBook(input)
}

// UpdateBook is the resolver for the UpdateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.BookInput) (*model.Book, error) {
	return books.UpdateBook(id, input)
}

// DeleteBook is the resolver for the DeleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (bool, error) {
	return books.DeleteBook(id)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return books.Books()
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	return books.Book(id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package domain

import "context"

type BookRepository interface {
	Create(context.Context, Book) error
	GetAll(context.Context) ([]Book, error)
	GetById(context.Context, int) (Book, error)
	Update(context.Context, Book) error
	Delete(context.Context, int) (Book, error)
}

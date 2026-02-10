package repository

import (
	"context"
	"errors"

	"github.com/abhirup7477/go-books-api/internal/domain"
)

type InMemoryRepo struct {
	books map[int]domain.Book
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{books: make(map[int]domain.Book)}
}

func (r *InMemoryRepo) Create(ctx context.Context, book domain.Book) error {
	_, ok := r.books[book.Id]
	if ok {
		return errors.New("Id already exists")
	}
	r.books[book.Id] = book
	return nil
}

func (r *InMemoryRepo) GetAll(ctx context.Context) ([]domain.Book, error) {
	books := make([]domain.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	if len(books) == 0 {
		return books, errors.New("No book found")
	}
	return books, nil
}

func (r *InMemoryRepo) GetById(ctx context.Context, id int) (domain.Book, error) {
	book, ok := r.books[id]
	if !ok {
		return book, errors.New("Book not Found")
	}
	return book, nil
}

func (r *InMemoryRepo) Update(ctx context.Context, book domain.Book) error {
	_, ok := r.books[book.Id]
	if !ok {
		return errors.New("Book not Found")
	}
	r.books[book.Id] = book
	return nil
}

func (r *InMemoryRepo) Delete(ctx context.Context, id int) (domain.Book, error) {
	book, ok := r.books[id]
	if !ok {
		return book, errors.New("Book not Found")
	}
	delete(r.books, id)
	return book, nil
}

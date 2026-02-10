package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/abhirup7477/go-books-api/internal/domain"
)

type BookUsecase struct {
	repo   domain.BookRepository
	mailer domain.Mailer
}

func NewBookUsecase(r domain.BookRepository, m domain.Mailer) *BookUsecase {
	return &BookUsecase{repo: r, mailer: m}
}

func (uc *BookUsecase) CreateBook(
	ctx context.Context, book domain.Book, email string,
) error {
	err := uc.repo.Create(ctx, book)
	if err != nil {
		return err
	}
	sub := "Book created successfully"
	body := "Book Added"
	go uc.SendEmailBG(ctx, email, sub, body)
	return nil
}

func (uc *BookUsecase) GetBooks(
	ctx context.Context, email string,
) ([]domain.Book, error) {
	books, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	sub := "Books Retreived"
	body := "All books retreived"
	go uc.SendEmailBG(ctx, email, sub, body)
	return books, nil
}

func (uc *BookUsecase) GetBook(
	ctx context.Context, id int, email string,
) (domain.Book, error) {
	book, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return book, err
	}
	sub := "Book Retreived"
	body := fmt.Sprintf("Book-%+v retreived successfully\r\n", book.Id)
	body += fmt.Sprintf("Book Info: %+v\r\n", book)
	go uc.SendEmailBG(ctx, email, sub, body)
	return book, nil
}

func (uc *BookUsecase) UpdateBook(
	ctx context.Context, book domain.Book, email string,
) error {
	err := uc.repo.Update(ctx, book)
	if err != nil {
		return err
	}
	sub := "Book Updated"
	body := fmt.Sprintf("Book-%+v Updated successfully\r\n", book.Id)
	body += fmt.Sprintf("Book Info: %+v\r\n", book)
	go uc.SendEmailBG(ctx, email, sub, body)
	return nil
}

func (uc *BookUsecase) DeleteBook(
	ctx context.Context, id int, email string,
) (domain.Book, error) {
	book, err := uc.repo.Delete(ctx, id)
	if err != nil {
		return book, err
	}
	sub := "Book Deleted"
	body := fmt.Sprintf("Book-%+v Deleted successfully\r\n", book.Id)
	body += fmt.Sprintf("Book Info: %+v\r\n", book)
	go uc.SendEmailBG(ctx, email, sub, body)
	return book, nil
}

func (uc *BookUsecase) SendEmailBG(
	ctx context.Context, email string, sub string, body string,
) {
	bg := context.Background()

	if email == "" {
		log.Println("No email id found!")
	} else {
		err := uc.mailer.SendEmail(bg, email, sub, body)
		if err != nil {
			log.Printf("Email failed: %v\n", err)
		}
	}
}

package services

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
)

func (s Service) InsertBook(ctx context.Context, book models.Book) (models.Book, error) {
	if book.Title == "" {
		return models.Book{}, ErrBookTitleEmpty
	}
	if book.Description == "" {
		return models.Book{}, ErrBookDescriptionEmpty
	}
	if book.Genre == nil {
		return models.Book{}, ErrBookGenreEmpty
	}
	if book.Author == nil {
		return models.Book{}, ErrBookAuthorEmpty
	}
	if book.CategoryId == 0 {
		return models.Book{}, ErrBookCategoryIdEmpty
	}

	return s.Repository.InsertBook(ctx, book)
}

func (s Service) ListBooks (ctx context.Context) ([]models.Book, error) {
	books, err := s.Repository.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s Service) GetBookById(ctx context.Context, id int64) (models.Book, error) {
	book, err := s.Repository.GetBookById(ctx, id)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s Service) UpdatateBook(ctx context.Context, id int64) error {
	err := s.Repository.UpdatateBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteBook(ctx context.Context, id int64) error {
	err := s.Repository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) SearchBook(ctx context.Context, searchQ string) ([]models.Book, error) {
	books, err := s.Repository.SearchBook(ctx, searchQ)
	if err != nil {
		return nil, err
	}
	return books, nil
}

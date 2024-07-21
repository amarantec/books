package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/picpay/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryPostgres) InsertBook (ctx context.Context, book models.Book) (models.Book, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO books (title, description, genre, author, category_id, user_id) VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, title, description, genre, author, category_id, user_id`, book.Title, book.Description, book.Genre, book.Author, book.CategoryId, book.UserId).Scan(
			&book.Id, &book.Title, &book.Description, &book.Genre, &book.Author, &book.CategoryId, &book.UserId)

			if err != nil {
				return models.Book{}, err
			}

			return book, nil
}

func (r *RepositoryPostgres) ListBooks (ctx context.Context) ([]models.Book, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, title, description, genre, author, category_id, user_id FROM categories`,)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var books []models.Book
		for rows.Next() {
			var book models.Book
			if err := rows.Scan(
				&book.Id,
				&book.Title,
				&book.Description,
				&book.Genre,
				&book.Author,
				&book.CategoryId,
				&book.UserId); err != nil {
					return nil, err
				}
				books = append(books, book)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
		return books, nil
}

func (r *RepositoryPostgres) GetBookById(ctx context.Context, id int64) (models.Book, error) {
	var book = models.Book{Id: id}
	
	err := r.Conn.QueryRow(
		ctx,
		`SELECT title, description, genre, author, category_id, user_id WHERE id=$1`, id).Scan(
			&book.Title, &book.Description, &book.Genre, &book.Author, &book.CategoryId, &book.UserId)

		if err == pgx.ErrNoRows {
			return models.Book{}, errors.New("book not found")
		}

		if err != nil {
			return models.Book{}, err
		}

		return book, nil

}

func (r *RepositoryPostgres) DeleteBook(ctx context.Context, id int64) error {
	tag, err := r.Conn.Exec(
		ctx,
		`DELETE FROM books WHERE id=$1`, id)
		if tag.RowsAffected() == 0 {
			return errors.New("category not found")
		}

		return err
}

func (r *RepositoryPostgres) UpdatateBook(ctx context.Context, id int64) error {
	var book = models.Book{Id: id}
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE books SET title = $2, description = $3, genre = $4, author = $5, category_id = $6 WHERE id = $1`, id, book.Title, book.Description, book.Genre, book.Author, &book.CategoryId)
		if err != nil {
			return err
		}
		return nil
}

func (r *RepositoryPostgres) SearchBook(ctx context.Context, searchQ string) ([]models.Book, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, title, description, genre, author, category_id, user_id WHERE title ILIKE '%' $1 || ´%´ OR description ILIKE '%' || '%';`,searchQ)

		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var books []models.Book
		for rows.Next() {
			var book models.Book
			if err := rows.Scan(
				&book.Id,
				&book.Title,
				&book.Description,
				&book.Genre,
				&book.Author,
				&book.CategoryId,
				&book.UserId); err != nil {
					return nil, err
				}

				books = append(books, book)
		}

		return books, nil
}

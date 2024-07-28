package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/books/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryPostgres) InsertBook (ctx context.Context, book models.Book) (models.Book, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO books (title, description, genre, author, image_url, category_id, user_id) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id, title, description, genre, author, image_url,  category_id, user_id;`,
		 book.Title, book.Description, book.Genre, book.Author, book.ImageURL, book.CategoryId, book.UserId)						.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.Genre,
		  &book.Author,
		  &book.ImageURL,
			&book.CategoryId,
			&book.UserId)

			if err != nil {
				return models.Book{}, err
			}

			return book, nil
}

func (r *RepositoryPostgres) ListBooks (ctx context.Context) ([]models.Book, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT b.id, 
						b.title, 
						b.description, 
						b.genre,
						b.author,
						b.image_url,
						b.category_id,
						b.user_id
						c.id,
						c.name,
						c.url
						FROM books AS b
						JOIN categories AS c ON b.category_id = c.id;`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var books []models.Book
		for rows.Next() {
			var book models.Book
			var category models.Category
			if err := rows.Scan(
				&book.Id,
				&book.Title,
				&book.Description,
				&book.Genre,
				&book.Author,
       	&book.ImageURL,
				&book.CategoryId,
				&book.UserId,
				&category.Id,
				&category.Name,
				&category.Url); err != nil {
					return nil, err
				}
				book.Category = category
				books = append(books, book)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
		return books, nil
}

func (r *RepositoryPostgres) GetBookById(ctx context.Context, id int64) (models.Book, error) {
	var book = models.Book{Id: id}
	var category models.Category	
	err := r.Conn.QueryRow(
		ctx,
		``SELECT b.id, 
						 b.title, 
						 b.description, 
					 	 b.genre,
						 b.author,
						 b.image_url,
						 b.category_id,
						 b.user_id
						 c.id,
						 c.name,
						 c.url
						 FROM books AS b
						 JOIN categories AS c ON b.category_id = c.id
						 WHERE b.id = $1;`, id).Scan(
			&book.Title,
		  &book.Description,
		  &book.Genre,
		  &book.Author,
		  &book.ImageURL,
			&book.CategoryId,
		  &book.UserId,
		  &category.Id,
		  &category.Name,
		  &category.Url)

		book.Category = category

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
		`DELETE FROM books WHERE id=$1;`, id)
		if tag.RowsAffected() == 0 {
			return errors.New("category not found")
		}

		return err
}

func (r *RepositoryPostgres) UpdatateBook(ctx context.Context, id int64) error {
	var book = models.Book{Id: id}
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE books SET title = $2, description = $3, genre = $4, author = $5, image_url = $6, category_id = $7 WHERE id = $1;`, id, book.Title, book.Description, book.Genre, book.Author, book.ImageURL, book.CategoryId)
		if err != nil {
			return err
		}
		return nil
}

func (r *RepositoryPostgres) SearchBook(ctx context.Context, searchQ string) ([]models.Book, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT b.id,
					  b.title,
					  b.description,
					  b.genre,
					  b.author,
					  b.image_url,
						b.category_id,
					  b.user_id,
						c.id,
						c.name,
						c.url,
					  FROM books AS b 
						JOIN categories AS c ON b.category_id = c.id
						WHERE title ILIKE '%' $1 || ´%´ OR description ILIKE '%' || '%';`, searchQ)


		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var books []models.Book
		for rows.Next() {
			var book models.Book
			var category models.Category
			if err := rows.Scan(
				&book.Id,
				&book.Title,
				&book.Description,
				&book.Genre,
				&book.Author,
        &book.ImageURL,
				&book.CategoryId,
				&book.UserId,
				&category.Id,
				&category.Name,	
				&category.Url); err != nil {
					return nil, err
				}
				book.Category = category
				books = append(books, book)
		}

		return books, nil
}

func (r *RepositoryPostgres) FindBookByCategory(ctx context.Context, categoryUrl string) ([]models.Book, error) {
	rows, err := r.Conn.QueryRow(
		ctx,
		`SELECT b.id,
						b.title,
						b.description,
						b.genre,
						b.author,
						b.image_url,
						b.category_id,
						b.user_id,
						c.id,
						c.name,
						c.url
						FROM books AS b
						JOIN categories AS c ON b.category_id = c.id
						WHERE c.url = $1;`, categoryUrl)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		var category models.Category
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.Genre,
			&book.Author,
			&book.ImageURL,
			&book.CategoryId
			&book.UserId,
			&category.Id,
			&category.Name,
			&category.Url); err != nil {
				return nil, err
			}
			book.Category = category
			books = append(books, book)	
	}	
	return books, nil
}

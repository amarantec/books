package repositories

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

type Repository interface {
	SaveUser(ctx context.Context, user models.User) (models.User, error)
	ValidateUserCredentials(ctx context.Context, user models.User) (int64, error)
	InsertCategory(ctx context.Context, category models.Category) (models.Category, error)
	ListCategories(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	InsertBook(ctx context.Context, book models.Book) (models.Book, error)
	ListBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(ctx context.Context, id int64) (models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
	UpdatateBook(ctx context.Context, id int64) error
	SearchBook(ctx context.Context, searchQ string) ([]models.Book, error)
}

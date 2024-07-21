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
	ValidateUserCredentials(ctx context.Context, user models.User) error
	InsertCategory(ctx context.Context, category models.Category) (models.Category, error)
	ListCategories(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id int64) error
}

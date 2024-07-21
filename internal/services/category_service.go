package services

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
)

func (s Service) InsertCategory(ctx context.Context, category models.Category) (models.Category, error) {
	if category.Name == "" {
		return models.Category{}, ErrCategoryNameEmpty
	}
	if category.Url == "" {
		return models.Category{}, ErrCategoryUrlEmpty
	}

	return s.Repository.InsertCategory(ctx, category)
}

func (s Service) ListCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.Repository.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s Service) DeleteCategory(ctx context.Context, id int64) error {
	return s.Repository.DeleteCategory(ctx, id)
}

package services

import (
	"context"

	"github.com/amarantec/books/internal/models"
)

func (s Service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "" {
		return models.User{}, ErrUserNameEmpty
	}
	if user.Email == "" {
		return models.User{}, ErrUserEmailEmpty
	}
	if user.Password == "" {
		return models.User{}, ErrUserPasswordEmpty
	}

	return s.Repository.SaveUser(ctx, user)
}

func (s Service) ValidateUserCredentials(ctx context.Context, user models.User) (int64, error) {
	if user.Email == "" {
		return 0, ErrUserEmailEmpty
	}
	if user.Password == "" {
		return 0, ErrUserPasswordEmpty
	}
	return s.Repository.ValidateUserCredentials(ctx, user)
}

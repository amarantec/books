package services

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
)

func (s Service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "" {
		return models.User{}, ErrUserFirstNameEmpty
	}
	if user.Email == "" {
		return models.User{}, ErrUserEmailEmpty
	}
	if user.Password == "" {
		return models.User{}, ErrUserPasswordEmpty
	}

	return s.Repository.SaveUser(ctx, user)
}

func (s Service) ValidateUserCredentials(ctx context.Context, user models.User) error {
	if user.Email == "" {
		return ErrUserEmailEmpty
	}
	if user.Password == "" {
		return ErrUserPasswordEmpty
	}
	return s.Repository.ValidateUserCredentials(ctx, user)
}

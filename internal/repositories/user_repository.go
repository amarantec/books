package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/picpay/internal/models"
	"github.com/amarantec/picpay/internal/utils"
)

func (r *RepositoryPostgres) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	err = r.Conn.QueryRow(
		ctx,
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email`,
		user.Name, user.Email, hashedPassword).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *RepositoryPostgres) ValidateUserCredentials(ctx context.Context, user models.User) error {
	var retriviedPassword string

	err := r.Conn.QueryRow(
		ctx,
		`SELECT id, first_name, last_name, document, password FROM users WHERE email=$1`, user.Email).Scan(&user.Id, &user.Name, &retriviedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(user.Password, retriviedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}

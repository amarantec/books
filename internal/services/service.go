package services

import (
	"errors"

	"github.com/amarantec/picpay/internal/repositories"
)

type Service struct {
	Repository repositories.Repository
}

var ErrUserNameEmpty = errors.New("user first name is empty")
var ErrUserEmailEmpty = errors.New("user email is empty")
var ErrUserPasswordEmpty = errors.New("user password is empty")
var ErrCategoryNameEmpty = errors.New("category name is empty")
var ErrCategoryUrlEmpty = errors.New("category url is empry")

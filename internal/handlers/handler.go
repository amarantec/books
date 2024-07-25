package handlers

import (
	"github.com/amarantec/books/internal/database"
	"github.com/amarantec/books/internal/repositories"
	"github.com/amarantec/books/internal/services"
)

var service services.Service

func Configure() {
	service = services.Service{
		Repository: &repositories.RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}

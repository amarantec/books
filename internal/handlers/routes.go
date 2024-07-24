package handlers

import (
	"net/http"

	"github.com/amarantec/picpay/internal/middleware"
)

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user-signup", signupUser)
	mux.HandleFunc("/user-login", loginUser)
	mux.HandleFunc("/insert-category", inserCategory)
	mux.HandleFunc("/list-categories", listCategories)
	mux.HandleFunc("/delete-category", deleteCategory)
	

	mux.HandleFunc("/insert-book", middleware.Authenticate(insertBook))
	mux.HandleFunc("/list-books", listBooks)
	mux.HandleFunc("/get-book/{id}", getBookById)
	mux.HandleFunc("/search-books", searchBook)
	mux.HandleFunc("/update-book/{id}", middleware.Authenticate(updateBook))
	mux.HandleFunc("/delete-book/{id}", middleware.Authenticate(deleteBook))
	return middleware.CorsMiddleware(mux)
}

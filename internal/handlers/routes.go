package handlers

import (
	"net/http"
)

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user-signup", signupUser)
	mux.HandleFunc("/user-login", loginUser)
	mux.HandleFunc("/insert-category", inserCategory)
	mux.HandleFunc("/list-categories", listCategories)
	mux.HandleFunc("/delete-category", deleteCategory)
	return mux
}

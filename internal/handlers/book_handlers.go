package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/picpay/internal/middleware"
	"github.com/amarantec/picpay/internal/models"
	"github.com/amarantec/picpay/internal/services"
)

func insertBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	userId := r.Context().Value(middleware.UserIdKey).(int64)
	newBook.UserId = userId

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "could not decode this book", http.StatusBadRequest)
		return
	}
	
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	nB, err := service.InsertBook(ctxTimeout, newBook)
	if err != nil {
		http.Error(w, "could not insert this book", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(nB, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	books, err := service.ListBooks(ctxTimeout)
	if err != nil {
		http.Error(w, "could not fetch books", http.StatusBadRequest)
		return
	}

	jsonResp, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/get-book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	book, err := service.GetBookById(ctxTimeout, int64(id))
	if err != nil {
		http.Error(w, "could not get this book", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func searchBook(w http.ResponseWriter, r *http.Request) {
	searchQ := r.URL.Path[len("/search-books/"):]
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	books, err := service.SearchBook(ctxTimeout, searchQ)
	if err != nil {
		http.Error(w, "could not search this book", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}


func updateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/update-book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	var uBook models.Book
	
	userId := r.Context().Value(middleware.UserIdKey).(int64)
	uBook.UserId = userId


	err = json.NewDecoder(r.Body).Decode(&uBook)
	if err != nil {
		http.Error(w, "could not parse this book", http.StatusBadRequest)
		return
	}

	if err := service.UpdatateBook(ctxTimeout, int64(id)); err != nil {
		http.Error(w, "could not update this book", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Book %d updated", id)))
	
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/delete-book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dBook models.Book
	userId := r.Context().Value(middleware.UserIdKey).(int64)
	dBook.UserId = userId
		
	if err = service.DeleteBook(ctxTimeout, int64(id)); err != nil {
		if err == services.ErrBookNotFound {
		http.Error(w, "book not found", http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}


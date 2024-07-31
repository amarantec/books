package handlers

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
  "fmt"
	"net/http"

	"github.com/amarantec/books/internal/models"
)

func inserCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category

	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "could not decode this category", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nC, err := service.InsertCategory(ctxTimeout, newCategory)
	if err != nil {
		http.Error(w, "could not insert this category", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(nC, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}

func listCategories(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categories, err := service.ListCategories(ctxTimeout)
	if err != nil {
		http.Error(w, "could not list categories", http.StatusBadRequest)
		return
	}

	jsonResp, err := json.MarshalIndent(categories, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/delete-category/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	
	err = service.DeleteCategory(ctxTimeout, int64(id))
	if err != nil {
		http.Error(w, "could not delete this category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
  idStr := r.URL.Path[len("/update-category/"):]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }
  
  var uCategory models.Category
  
  err = json.NewDecoder(r.Body).Decode(&uCategory)
  if err != nil {
    http.Error(w, "could not parse this category", http.StatusBadRequest)
    return
  }
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  if err := service.UpdateCategory(ctxTimeout, int64(id)); err != nil {
      http.Error(w, "could not update this category", http.StatusInternalServerError)
      return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}

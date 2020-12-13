package controllers

import (
	"alvinr.ca/learn-go/api/data"
	"alvinr.ca/learn-go/api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// GetBooks gets all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.BooksDB)
}

// GetBook single
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range data.BooksDB {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Book{})
}

// CreateBooks and stores into fake DB 'BooksDB'
func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(rand.Intn(10000000))
	data.BooksDB = append(data.BooksDB, newBook)

	json.NewEncoder(w).Encode(newBook)
}

// UpdateBooks from fake DB 'BooksDB'
func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range data.BooksDB {
		if item.ID == params["id"] {

			// Delete existing book at that index
			data.BooksDB = append(data.BooksDB[:index], data.BooksDB[index+1:]...)

			// Create new book at that index
			var newBook models.Book
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = item.ID
			data.BooksDB = append(data.BooksDB, newBook)

			json.NewEncoder(w).Encode(newBook)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Book{})
}

// DeleteBooks from fake DB 'BooksDB'
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range data.BooksDB {
		if item.ID == params["id"] {
			data.BooksDB = append(data.BooksDB[:index], data.BooksDB[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(data.BooksDB)
}

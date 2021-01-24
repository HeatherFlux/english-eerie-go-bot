package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "go.uber.org/zap"
)

// Book Model
type Book struct {
	ID     string  `json:"ID"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author  model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as slice of Book struct. slice variable length array
var books []Book

var logger, _ = log.NewProduction()
var sugar = logger.Sugar()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((books))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	http.NotFound(w, r)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func editBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id not real
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "meh", Title: "B1", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "meh", Title: "B2", Author: &Author{Firstname: "JJ", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "meh", Title: "B3", Author: &Author{Firstname: "Jane", Lastname: "Hurlet"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", editBook).Methods("PUT")

	sugar.Fatal(http.ListenAndServe(":8080", r))
}

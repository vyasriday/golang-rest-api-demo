package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Data Scehma for  a Book
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Schema
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books variables as a slice of structs

var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// set params variable
	params := mux.Vars(r)
	// loop thorugh books to find correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Int())
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// we return updated books from here
	json.NewEncoder(w).Encode(books)
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init mux router
	Router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "1211210912", Title: "Book Title", Author: &Author{Firstname: "Hridayesh", Lastname: "Sharma"}})
	books = append(books, Book{ID: "2", Isbn: "12112109sd", Title: "Book Title 2", Author: &Author{Firstname: "Hriday", Lastname: "Sharma"}})
	// Create route handlers for a different Request that takes a handler method
	Router.HandleFunc("/api/books", getBooks).Methods("GET")
	Router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// Router.HandleFunc("/api/books", createBook).Methods('POST')
	// Router.HandleFunc("/api/books/{id}", updateBook).Methods('PUT')
	// Router.HandleFunc("/api/books/{id}", deleteBook).Methods('DELETE')

	// create server and run it

	log.Fatal(http.ListenAndServe(":8000", Router))

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	BOOKID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init Books variable as a slice Book struct
var books []Book

// Get all Books
func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get params

	// loop through books and find with id
	for _, item := range books {
		if item.BOOKID == params["bookid"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func addBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.BOOKID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for idx, item := range books {
		if item.BOOKID == params["bookid"] {
			books = append(books[:idx], books[idx+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.BOOKID = params["bookid"]
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
}

func deleteBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for idx, item := range books {
		if item.BOOKID == params["bookid"] {
			books = append(books[:idx], books[idx+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Console output to let user know server is running
	fmt.Println("Server started on port 8000")

	// Mock Data - @todo - implement DB
	books = append(books, Book{BOOKID: "1", Isbn: "0123456789999", Title: "Test Book 1",
		Author: &Author{Firstname: "Chris", Lastname: "Girvin"}})
	books = append(books, Book{BOOKID: "2", Isbn: "0123456790004", Title: "Test Book 2",
		Author: &Author{Firstname: "Thomas", Lastname: "Girvin"}})
	books = append(books, Book{BOOKID: "3", Isbn: "0123456790006", Title: "Test Book 3",
		Author: &Author{Firstname: "Janette", Lastname: "Girvin"}})
	books = append(books, Book{BOOKID: "4", Isbn: "0123456790018", Title: "Test Book 4",
		Author: &Author{Firstname: "Jim", Lastname: "Thom"}})
	books = append(books, Book{BOOKID: "5", Isbn: "0123456790022", Title: "Test Book 5",
		Author: &Author{Firstname: "Rob", Lastname: "Thomas"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books/books", getAllBooks).Methods("GET")
	router.HandleFunc("/api/books/book/{bookid}", getBookById).Methods("GET")
	router.HandleFunc("/api/books/books", addBook).Methods("POST")
	router.HandleFunc("/api/books/book/{bookid}", updateBookById).Methods("PUT")
	router.HandleFunc("/api/books/book/{bookid}", deleteBookById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
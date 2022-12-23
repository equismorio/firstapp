package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	// "github.com/gorilla/mux"
)

// Book struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func helloHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	// io.WriteString(w, "Hello, Go!\n")
	fmt.Printf("Query: %v", method)
	switch method {
	case "GET":
		// w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	// io.WriteString(w, "Hello, Go!\n")
	fmt.Printf("Query: %v", method)
	switch method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(books)
	}
}
func main() {
	// Init Router
	r := mux.NewRouter()

	// Hello world, the web server

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "438228", Title: "Book Two", Author: &Author{Firstname: "Jane", Lastname: "Doe"}})

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/api/addBook", addHandler)
	// http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

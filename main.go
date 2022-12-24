package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/books" {
		w.Header().Set("Content-Type", "application/json")

		m, _ := url.ParseQuery(r.URL.RawQuery)
		if len(m) != 0 {
			v := m["id"][0]
			_, err := strconv.Atoi(v)
			if err == nil {
				for _, item := range books {
					if item.ID == v {
						json.NewEncoder(w).Encode(item)
						break
					}
				}

			} else {
				json.NewEncoder(w).Encode(books)
			}
			// fmt.Println(m)
			fmt.Println(m["id"][0])
		} else {
			json.NewEncoder(w).Encode(books)
		}
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/book" {
		fmt.Println(" Delete request being processed...")
		w.Header().Set("Content-Type", "application/json")

		m, _ := url.ParseQuery(r.URL.RawQuery)
		if len(m) != 0 {
			v := m["id"][0]
			_, err := strconv.Atoi(v)
			if err == nil {
				var book Book
				for index, item := range books {
					if item.ID == v {
						books = append(books[:index], books[index+1:]...)
						book = item
						json.NewEncoder(w).Encode(book)
						break
					}
				}
				if (Book{} == book) {
					http.Error(w, "Bad Request. (ID was not found.) try again", http.StatusBadRequest)
				}
			} else {
				http.Error(w, "Bad Request. (ID is not numeric.) Pls try again", http.StatusBadRequest)
			}
			// fmt.Println(m)
			fmt.Println(m["id"][0])
		} else {
			http.Error(w, "Bad Request. (Missing book's ID number.) try again", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Bad Request. Pls try again", http.StatusBadRequest)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/book" {
		fmt.Println(" Update request being processed...")
		w.Header().Set("Content-Type", "application/json")
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		if (Book{} != book) {
			v := book.ID
			for index, item := range books {
				if item.ID == v {
					books = append(books[:index], books[index+1:]...)
					books = append(books, book)
					json.NewEncoder(w).Encode(book)
					break
				}
			}
		} else {
			http.Error(w, "Bad Request. (Book object missing in request.) Pls try again", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Bad Request. Pls try again", http.StatusBadRequest)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/book" {
		w.Header().Set("Content-Type", "application/json")
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		fmt.Println("Book:", book)
		if (Book{} != book) {
			book.ID = strconv.Itoa(rand.Intn(10000000)) // not safe for production
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
		} else {
			http.Error(w, "Bad Request, try again", http.StatusBadRequest)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	fmt.Println(" URL: " + r.URL.Path)
	fmt.Println(" RawQuery: " + r.URL.RawQuery)
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		panic(err)
	}
	if len(m) != 0 {
		fmt.Println(m)
		fmt.Println(m["id"][0])
	}
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		// Handle the GET request
		getHandler(w, r)
		break
	case http.MethodPost:
		// Handle the POST request
		postHandler(w, r)
		break
	case http.MethodDelete:
		// Handle the DELETE request
		deleteHandler(w, r)
		break
	case http.MethodPut:
		// Handle the Update request
		updateHandler(w, r)
		break

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)
		http.Error(w, "method not allowed. Method:"+r.Method, http.StatusMethodNotAllowed)
	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed. Method:"+r.Method, http.StatusMethodNotAllowed)
	}
}
func main() {
	// Init Router
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	// Hello world, the web server

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "438228", Title: "Book Two", Author: &Author{Firstname: "Jane", Lastname: "Doe"}})

	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/api/addBook", addHandler)
	// http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

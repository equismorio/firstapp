package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		// fmt.Fprint(w, "<h1>Welcom to the world!</h1>")
		w.Write([]byte("<h1 style='color: steelblue'>Welcome to the world!</h1>"))
	} else if r.URL.Path == "/contact" {
		// fmt.Fprint(w, "To get in touch, please send an email to <a href=\"equismorio@gmail.com\">equismorio@gmail.com</a>.")
		w.Write([]byte("<h1 style='color: red'>To get in touch, please send an email to <a href=\"equismorio@gmail.com\">equismorio@gmail.com</a>."))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :( </h1><p>Please try again.</a>")
	}
}
func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

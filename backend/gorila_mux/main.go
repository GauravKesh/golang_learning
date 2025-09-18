package main

/*
1. The biggest strength of the gorilla/mux Router is the ability to extract segments from the request URL.
2. Once you have a new router you can register request handlers like usual. Instead of http.HandleFunc(...),
   you call HandleFunc on your router like this: r.HandleFunc(...).
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func serverStartLog(p string) {
	fmt.Printf("SERVER RUNNING AT PORT %s\n", p)
}

func getNoOfUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(res, "10")
}

func getBookData(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]
	page := vars["page"]

	data := map[string]string{
		"chapter_name": title,
		"page_no":      page,
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func BookHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	data := map[string]string{
		"chapter_name": title,
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

// Get all books
func AllBooks(res http.ResponseWriter, req *http.Request) {
	data := []map[string]string{
		{"title": "Book One", "author": "Author A"},
		{"title": "Book Two", "author": "Author B"},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

// Get specific book by title
func GetBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	data := map[string]string{
		"title":  title,
		"author": "Author Unknown",
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func main() {
	// main router for application
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	const PORT string = "8080"

	api.HandleFunc("/user/count", getNoOfUser).Methods("GET")

	api.HandleFunc("/book/{title}/page/{page}", getBookData).Methods("GET")

	// Hostnames & Subdomains
	r.HandleFunc("/books/{title}", BookHandler).Host("localhost")

	// Path Prefixes & Subrouters
	bookrouter := api.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks).Methods("GET")
	bookrouter.HandleFunc("/{title}", GetBook).Methods("GET")

	serverStartLog(PORT) // log before starting server

	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

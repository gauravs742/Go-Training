package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func sequence(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}

// Book represents a single book
type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}

var nextID = sequence(0)

var books = map[int]Book{
	1: Book{ID: nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
	2: Book{ID: nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
}

// GET /books/{id}
func bookShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])
	log.Info("Getting bookID: ", bookID)
	book := books[bookID]

	json.NewEncoder(w).Encode(book)
}

// GET /books
func booksIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	port := "9000"
	r := mux.NewRouter()

	r.HandleFunc("/books", booksIndexHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookShowHandler).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	log.Info("Server is running on port: ", port)

	http.ListenAndServe(":"+port, n)
}

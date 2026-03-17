package main

import (
	"net/http"
	"fmt"
)

func PrintBook(w http.ResponseWriter, r *http.Request) {
	book := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	fmt.Fprintf(w, "%v by %v - %v copies\n", book.Title, book.Author, book.Copies)
}

func NewRouter() http.Handler {
	r := http.DefaultServeMux
	r.HandleFunc("GET /books", PrintBook)
	return r
}

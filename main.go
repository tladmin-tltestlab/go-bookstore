package main

import "fmt"
import "net/http"
import "os"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	router := NewRouter()
	err := http.ListenAndServe("0.0.0.0:8081", router)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
}

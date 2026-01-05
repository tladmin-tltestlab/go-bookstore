package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	book := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	fmt.Println("Books in stock:")
	printBook(book)
}

func printBook(book Book) {
	fmt.Println(book.Title, "by", book.Author, "-", book.Copies, "copies")
}

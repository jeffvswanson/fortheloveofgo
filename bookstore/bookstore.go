package main

import "fmt"

// Book represents information about a book.
type Book struct {
	Title   string
	Author  string
	Copies  int
	Edition int
}

func main() {
	book := Book{
		Title:   "For the Love of Go",
		Author:  "Jeff Swanson",
		Copies:  1,
		Edition: 2,
	}
	fmt.Printf("%+v\n", book)
}

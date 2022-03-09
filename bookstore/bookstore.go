package main

import "fmt"

// Book represents information about a book.
type Book struct {
	Title       string
	Author      string
	Copies      int
	Edition     int
	Price       int // Price in cents
	SeriesOrder int
}

func main() {
	b := Book{
		Title:       "For the Love of Go",
		Author:      "Jeff Swanson",
		Copies:      1,
		Edition:     2,
		Price:       100,
		SeriesOrder: 1,
	}
	fmt.Printf("%+v\n", b)
}

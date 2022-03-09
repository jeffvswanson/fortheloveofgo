package main

import "fmt"

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	Edition         int
	Price           int // Price in cents
	SeriesOrder     int
	DiscountPercent uint // 10 represents a 10% discount
}

func main() {
	b := Book{
		Title:           "For the Love of Go",
		Author:          "Jeff Swanson",
		Copies:          1,
		Edition:         2,
		Price:           100,
		SeriesOrder:     1,
		DiscountPercent: 0,
	}
	fmt.Printf("%+v\n", b)
}

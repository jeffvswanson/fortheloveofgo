package main

import (
	"book"
	"fmt"
)

func main() {
	b := book.Book{
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

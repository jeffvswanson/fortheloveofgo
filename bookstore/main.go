package main

import (
	"bookstore/book"
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
		DiscountPercent: 5,
	}
	fmt.Printf("%+v\n", b)
	fmt.Printf("Price after discount: %d\n", b.NetPrice())
}

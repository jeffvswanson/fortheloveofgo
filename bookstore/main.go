package main

import (
	"bookstore/book"
	"bookstore/catalog"
	"fmt"
)

func main() {
	b := book.Book{
		Title:           "For the Love of Go",
		Author:          []string{"Jeff", "Theodore"},
		Copies:          1,
		Edition:         2,
		Price:           100,
		SeriesOrder:     1,
		DiscountPercent: 5,
	}
	c := make(catalog.Catalog, 0)
	c = c.AddToCatalog(b)
	b = book.Book{
		Title:           "Yummy Cooking",
		Author:          []string{"A Cook"},
		Copies:          10,
		Edition:         3,
		Price:           500,
		SeriesOrder:     1,
		DiscountPercent: 0,
		Featured:        true,
	}
	c = c.AddToCatalog(b)
	f := c.FeaturedBooks()
	for _, book := range f {
		fmt.Printf("Title: %s, Author: %v\n", book.Title, book.Author)
		fmt.Printf("Price after discount: %d\n", book.NetPrice())
	}
}

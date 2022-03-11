package catalog

import "bookstore/book"

type Catalog []book.Book

// Appends a book to the bookstore catalog
func (c Catalog) AddToCatalog(b book.Book) Catalog {
	c = append(c, b)
	return c
}

// Returns a Catalog of only books part of the Read of the Month program
func (c Catalog) FeaturedBooks() Catalog {
	f := make(Catalog, 0)
	for _, b := range c {
		if b.Featured {
			f = append(f, b)
		}
	}
	return f
}

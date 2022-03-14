package catalog

import (
	"bookstore/book"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	book0 := book.Book{Title: "Hello There"}
	book1 := book.Book{Title: "Goodbye"}
	want := Catalog{book0, book1}
	got := want.GetAllBooks()
	cmp.Equal(got, want)
}

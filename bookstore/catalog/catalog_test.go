package catalog

import (
	"bookstore/book"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddToCatalog(t *testing.T) {
	book1 := book.Book{Title: "First!"}
	book2 := book.Book{Title: "Segundo"}
	tests := []struct {
		c    Catalog
		b    book.Book
		want Catalog
	}{
		{nil, book1, Catalog{book1}},                   // nil Catalog
		{make(Catalog, 0), book1, Catalog{book1}},      // Empty catalog
		{Catalog{book1}, book2, Catalog{book1, book2}}, // Add a Book
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("AddToCatalog%d(%v, %+v", i, tc.c, tc.b), func(t *testing.T) {
			got := tc.c.AddToCatalog(tc.b)
			if !cmp.Equal(got, tc.want) {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestGetAllBooks(t *testing.T) {
	book0 := book.Book{Title: "Hello There"}
	book1 := book.Book{Title: "Goodbye"}
	want := Catalog{book0, book1}
	got := want.GetAllBooks()
	cmp.Equal(got, want)
}
